package common

import (
	"context"
	"fmt"
	"log"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type RevisionOsLayerConfig struct {
	// Command
	LayerName                   string `mapstructure:"layer_name"`
	BaseVersionName             string `mapstructure:"base_version_name"`
	PackagingDiskFileName       string `mapstructure:"packaging_disk_file_name"` // optional, defaults to LayerName
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name"`

	// Command.RevisionInfo
	VersionName        string `mapstructure:"version_name"`
	VersionDescription string `mapstructure:"version_description"`
	VersionSizeGb      int32  `mapstructure:"version_size_gb"` // in GB, default 10
	// Command.Reason
	Comment          string `mapstructure:"comment"`
	SkipCleanupOnFailure bool   `mapstructure:"skip_cleanup_on_failure"`
}

type StepRevisionOsLayering struct {
	Config *RevisionOsLayerConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepRevisionOsLayering) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)
	// Get OsLayerRevisionId (base version)
	osLayerRevisionId, err := helper.GetOsLayerRevisionId(s.Config.LayerName, s.Config.BaseVersionName)
	if err != nil {
		ui.Errorf("Error getting OsLayerRevisionId for LayerName=%s, BaseVersionName=%s: %v", s.Config.LayerName, s.Config.BaseVersionName, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "LayerName %s, BaseVersionName %s found OsLayerRevisionId: %d", s.Config.LayerName, s.Config.BaseVersionName, osLayerRevisionId)
	osLayerId, err := helper.GetOsLayerId(s.Config.LayerName)
	if err != nil {
		ui.Errorf("Error getting OsLayerId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "LayerName %s found OsLayerId: %d", s.Config.LayerName, osLayerId)
	// Get PlatformConnectorConfigId
	platformConnectorConfigId, err := helper.GetPlatformConnectorConfigId(s.Config.PlatformConnectorConfigName)
	if err != nil {
		ui.Errorf("Error getting PlatformConnectorConfigId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
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
		baseSize, sizeErr := helper.GetOsLayerRevisionSizeMiB(s.Config.LayerName, s.Config.BaseVersionName)
		if sizeErr != nil {
			ui.Errorf("Error getting base OS layer revision size: %v", sizeErr)
			state.Put("error", sizeErr)
			return multistep.ActionHalt
		}
		layerSizeMiB = baseSize
	}
	platformLayerCreatequest := &elmsoap.CreateOsLayerRevision{
		Command: &elmsoap.CreateOsLayerRevisionCommand{
			CreateLayerRevisionCommand: &elmsoap.CreateLayerRevisionCommand{
				LayerId: osLayerId,
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
				BaseLayerRevisionId:       &osLayerRevisionId,
			},
		},
	}
	revisionResp, err := helper.Client.CreateOsLayerRevision(platformLayerCreatequest)
	if err != nil {
		errMsg := fmt.Errorf("error calling RevisionOsLayer: %v", err)
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	if appErr := GetRevisionResultError(revisionResp.CreateOsLayerRevisionResult); appErr != nil {
		errMsg := fmt.Errorf("RevisionOsLayer failed: %s", FormatELMError(appErr))
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	workTicketId := revisionResp.CreateOsLayerRevisionResult.WorkTicketId
	UiSayf(ui, "Revision OS Layer, WorkTicketId: %d", workTicketId)
	state.Put("WORK_TICKET_ID", workTicketId)
	return multistep.ActionContinue
}

func (s *StepRevisionOsLayering) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	_, hasError := state.GetOk("error")

	if !cancelled && !halted && !hasError {
		return
	}
	if !cancelled && s.Config.SkipCleanupOnFailure {
		log.Printf("[INFO] StepRevisionOsLayering.Cleanup: failure detected but cleanup_on_failure is false, skipping work ticket cancellation")
		return
	}

	workTicketId, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		log.Printf("[WARN] StepRevisionOsLayering.Cleanup: no WORK_TICKET_ID in state, cannot cancel work ticket")
		return
	}
	helper, ok := state.GetOk("soap_helper")
	if !ok {
		log.Printf("[WARN] StepRevisionOsLayering.Cleanup: soap_helper not in state, cannot cancel work ticket")
		return
	}
	if err := helper.(*elmsoap.SoapHelper).CancelWorkTicket(workTicketId.(int64)); err != nil {
		log.Printf("[WARN] StepRevisionOsLayering.Cleanup: failed to cancel work ticket %d: %v", workTicketId.(int64), err)
	} else {
		log.Printf("[INFO] StepRevisionOsLayering.Cleanup: cancelled work ticket %d", workTicketId.(int64))
	}
}
