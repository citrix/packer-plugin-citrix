package common

import (
	"context"
	"fmt"
	"log"
	"strings"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type RevisionAppLayerConfig struct {
	// Command
	OsLayerName                 string   `mapstructure:"os_layer_name"`
	OsLayerVersionName          string   `mapstructure:"os_layer_version_name"`
	PlatformLayerName           string   `mapstructure:"platform_layer_name"`
	PlatformLayerVersionName    string   `mapstructure:"platform_layer_version_name"`
	LayerName                   string   `mapstructure:"layer_name"`
	BaseVersionName             string   `mapstructure:"base_version_name"`        // optional (empty = auto-latest)
	PrerequisiteLayers          []string `mapstructure:"prerequisite_layers"`      // optional, format "LayerName:VersionName"
	PackagingDiskFileName       string   `mapstructure:"packaging_disk_file_name"` // optional, defaults to LayerName
	PlatformConnectorConfigName string   `mapstructure:"platform_connector_config_name"`

	// Command.RevisionInfo
	VersionName        string `mapstructure:"version_name"`
	VersionDescription string `mapstructure:"version_description"`
	VersionSizeGb      int32  `mapstructure:"version_size_gb"` // in GB, default 10
	// Command.Reason
	Comment          string `mapstructure:"comment"`
	SkipCleanupOnFailure bool   `mapstructure:"skip_cleanup_on_failure"`
}

type StepRevisionAppLayering struct {
	Config *RevisionAppLayerConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepRevisionAppLayering) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)
	// Get OsLayerRevisionId
	osLayerRevisionId, err := helper.GetOsLayerRevisionId(s.Config.OsLayerName, s.Config.OsLayerVersionName)
	if err != nil {
		ui.Errorf("Error getting OsLayerRevisionId for OsLayerName=%s, OsLayerVersionName=%s: %v", s.Config.OsLayerName, s.Config.OsLayerVersionName, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	log.Printf("OsLayerName %s, OsLayerVersionName %s found OsLayerRevisionId: %d", s.Config.OsLayerName, s.Config.OsLayerVersionName, osLayerRevisionId)
	// PlatformLayerRevisionId is optional
	var platformLayerRevisionId *int64
	if s.Config.PlatformLayerName != "" && s.Config.PlatformLayerVersionName != "" {
		revId, platErr := helper.GetPlatformLayerRevisionId(s.Config.PlatformLayerName, s.Config.PlatformLayerVersionName)
		if platErr != nil {
			ui.Errorf("Error getting PlatformLayerRevisionId: %v", platErr)
			state.Put("error", platErr)
			return multistep.ActionHalt
		}
		platformLayerRevisionId = &revId
		log.Printf("PlatformLayerName %s, PlatformLayerVersionName %s found PlatformLayerRevisionId: %d", s.Config.PlatformLayerName, s.Config.PlatformLayerVersionName, revId)
	}
	appLayerId, err := helper.GetAppLayerId(s.Config.LayerName)
	if err != nil {
		ui.Errorf("Error getting AppLayerId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	log.Printf("LayerName %s found AppLayerId: %d", s.Config.LayerName, appLayerId)
	// Resolve base version (optional, empty = auto-detect latest)
	var appLayerRevisionId int64
	var baseRevisionSizeMegs int32
	if s.Config.BaseVersionName == "" {
		latestRev, latestErr := helper.GetLatestAppLayerRevision(s.Config.LayerName)
		if latestErr != nil {
			ui.Errorf("Error getting latest app layer revision for LayerName=%s: %v", s.Config.LayerName, latestErr)
			state.Put("error", latestErr)
			return multistep.ActionHalt
		}
		appLayerRevisionId = latestRev.Id
		baseRevisionSizeMegs = latestRev.SizeMegs
		log.Printf("LayerName %s: latest revision auto-detected, AppLayerRevisionId: %d", s.Config.LayerName, appLayerRevisionId)
	} else {
		var revErr error
		appLayerRevisionId, revErr = helper.GetAppLayerRevisionId(s.Config.LayerName, s.Config.BaseVersionName)
		if revErr != nil {
			ui.Errorf("Error getting AppLayerRevisionId for LayerName=%s, BaseVersionName=%s: %v", s.Config.LayerName, s.Config.BaseVersionName, revErr)
			state.Put("error", revErr)
			return multistep.ActionHalt
		}
		log.Printf("LayerName %s, BaseVersionName %s found AppLayerRevisionId: %d", s.Config.LayerName, s.Config.BaseVersionName, appLayerRevisionId)
	}
	// Get PlatformConnectorConfigId
	platformConnectorConfigId, err := helper.GetPlatformConnectorConfigId(s.Config.PlatformConnectorConfigName)
	if err != nil {
		ui.Errorf("Error getting PlatformConnectorConfigId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	log.Printf("PlatformConnectorConfigName %s found PlatformConnectorConfigId: %s", s.Config.PlatformConnectorConfigName, platformConnectorConfigId)
	// Auto-detect file share
	fileShareId, err := helper.GetDefaultFileShareId()
	if err != nil {
		ui.Errorf("Error getting default file share: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	// Default PackagingDiskFileName to LayerName
	if s.Config.PackagingDiskFileName == "" {
		s.Config.PackagingDiskFileName = s.Config.LayerName
	}
	// Hardcode disk format to VHD
	diskFormat := elmsoap.DiskFormatVhdDiskFormat
	// Convert GB to MiB, default to base version disk size
	layerSizeMiB := s.Config.VersionSizeGb * 1024
	if layerSizeMiB == 0 {
		// Use base version's actual disk size
		if s.Config.BaseVersionName != "" {
			baseSize, sizeErr := helper.GetAppLayerRevisionSizeMiB(s.Config.LayerName, s.Config.BaseVersionName)
			if sizeErr != nil {
				ui.Errorf("Error getting base app layer revision size: %v", sizeErr)
				state.Put("error", sizeErr)
				return multistep.ActionHalt
			}
			layerSizeMiB = baseSize
		} else {
			// BaseVersionName was empty → latest was auto-detected above, use its size
			layerSizeMiB = baseRevisionSizeMegs
		}
	}
	// Resolve prerequisite layers
	var prereqRevisionIds *elmsoap.ArrayOfLong
	if len(s.Config.PrerequisiteLayers) > 0 {
		seen := make(map[string]bool)
		ids := make([]int64, 0, len(s.Config.PrerequisiteLayers))
		for _, prereq := range s.Config.PrerequisiteLayers {
			parts := strings.SplitN(prereq, ":", 2)
			if len(parts) != 2 {
				err := fmt.Errorf("invalid prerequisite_layers entry %q: expected format 'LayerName:VersionName'", prereq)
				ui.Errorf("%v", err)
				state.Put("error", err)
				return multistep.ActionHalt
			}
			layerName, versionName := parts[0], parts[1]
			if seen[layerName] {
				err := fmt.Errorf("duplicate prerequisite layer name %q in prerequisite_layers", layerName)
				ui.Errorf("%v", err)
				state.Put("error", err)
				return multistep.ActionHalt
			}
			seen[layerName] = true
			revId, prereqErr := helper.GetAppLayerRevisionId(layerName, versionName)
			if prereqErr != nil {
				ui.Errorf("Error resolving prerequisite layer %q version %q: %v", layerName, versionName, prereqErr)
				state.Put("error", prereqErr)
				return multistep.ActionHalt
			}
			ids = append(ids, revId)
		}
		prereqRevisionIds = &elmsoap.ArrayOfLong{Long: ids}
	}

	appLayerCreatequest := &elmsoap.CreateAppLayerRevision{
		Command: &elmsoap.CreateAppLayerRevisionCommand{
			CreateLayerRevisionCommand: &elmsoap.CreateLayerRevisionCommand{
				LayerId: appLayerId,
				RevisionInfo: &elmsoap.LayerRevisionInfo{
					Name:         s.Config.VersionName,
					Description:  s.Config.VersionDescription,
					LayerSizeMiB: layerSizeMiB,
				},
				SelectedFileShare:     fileShareId,
				PackagingDiskFilename: s.Config.PackagingDiskFileName,
				PackagingDiskFormat:   &diskFormat,
				Reason: &elmsoap.ChangeSpecification{
					Description: s.Config.Comment,
				},
				PlatformConnectorConfigId: platformConnectorConfigId,
				BaseLayerRevisionId:       &appLayerRevisionId,
			},
			OsLayerRevisionId:               osLayerRevisionId,
			PlatformLayerRevisionId:         platformLayerRevisionId,
			CreationEnvironmentAppRevisions: prereqRevisionIds,
		},
	}
	revisionResp, err := helper.Client.CreateAppLayerRevision(appLayerCreatequest)
	if err != nil {
		errMsg := fmt.Errorf("error calling RevisionAppLayer: %v", err)
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	if appErr := GetRevisionResultError(revisionResp.CreateAppLayerRevisionResult); appErr != nil {
		errMsg := fmt.Errorf("RevisionAppLayer failed: %s", FormatELMError(appErr))
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	workTicketId := revisionResp.CreateAppLayerRevisionResult.WorkTicketId
	UiSayf(ui, "Revision APP Layer, WorkTicketId: %d", workTicketId)
	state.Put("WORK_TICKET_ID", workTicketId)
	return multistep.ActionContinue
}

func (s *StepRevisionAppLayering) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	_, hasError := state.GetOk("error")

	if !cancelled && !halted && !hasError {
		return
	}
	if !cancelled && s.Config.SkipCleanupOnFailure {
		log.Printf("[INFO] StepRevisionAppLayering.Cleanup: failure detected but cleanup_on_failure is false, skipping work ticket cancellation")
		return
	}

	workTicketId, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		log.Printf("[WARN] StepRevisionAppLayering.Cleanup: no WORK_TICKET_ID in state, cannot cancel work ticket")
		return
	}
	helper, ok := state.GetOk("soap_helper")
	if !ok {
		log.Printf("[WARN] StepRevisionAppLayering.Cleanup: soap_helper not in state, cannot cancel work ticket")
		return
	}
	if err := helper.(*elmsoap.SoapHelper).CancelWorkTicket(workTicketId.(int64)); err != nil {
		log.Printf("[WARN] StepRevisionAppLayering.Cleanup: failed to cancel work ticket %d: %v", workTicketId.(int64), err)
	} else {
		log.Printf("[INFO] StepRevisionAppLayering.Cleanup: cancelled work ticket %d", workTicketId.(int64))
	}
}
