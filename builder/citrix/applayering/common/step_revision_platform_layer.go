package common

import (
	"context"
	"fmt"
	"log"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type RevisionPlatformLayerConfig struct {
	// Platform type IDs (optional for revision, defaults to base version)
	HypervisorPlatform   string `mapstructure:"hypervisor_platform"`   // optional
	ProvisioningPlatform string `mapstructure:"provisioning_platform"` // optional, defaults to base version
	BrokerPlatform       string `mapstructure:"broker_platform"`       // optional, defaults to base version
	// Command
	OsLayerName                 string `mapstructure:"os_layer_name"`
	OsLayerVersionName          string `mapstructure:"os_layer_version_name"`
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

type StepRevisionPlatformLayering struct {
	Config *RevisionPlatformLayerConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepRevisionPlatformLayering) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)
	// Get OsLayerRevisionId
	osLayerRevisionId, err := helper.GetOsLayerRevisionId(s.Config.OsLayerName, s.Config.OsLayerVersionName)
	if err != nil {
		ui.Errorf("Error getting OsLayerRevisionId for OsLayerName=%s, OsLayerVersionName=%s: %v", s.Config.OsLayerName, s.Config.OsLayerVersionName, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "OsLayerName %s, OsLayerVersionName %s found OsLayerRevisionId: %d", s.Config.OsLayerName, s.Config.OsLayerVersionName, osLayerRevisionId)
	platformLayerRevisionId, err := helper.GetPlatformLayerRevisionId(s.Config.LayerName, s.Config.BaseVersionName)
	if err != nil {
		ui.Errorf("Error getting PlatformLayerRevisionId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "LayerName %s, BaseVersionName %s found PlatformLayerRevisionId: %d", s.Config.LayerName, s.Config.BaseVersionName, platformLayerRevisionId)
	platformLayerId, err := helper.GetPlatformLayerId(s.Config.LayerName)
	if err != nil {
		ui.Errorf("Error getting PlatformLayerId: %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "LayerName %s found PlatformLayerId: %d", s.Config.LayerName, platformLayerId)

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
		baseSize, sizeErr := helper.GetPlatformLayerRevisionSizeMiB(s.Config.LayerName, s.Config.BaseVersionName)
		if sizeErr != nil {
			ui.Errorf("Error getting base platform layer revision size: %v", sizeErr)
			state.Put("error", sizeErr)
			return multistep.ActionHalt
		}
		layerSizeMiB = baseSize
	}
	// Resolve platform type IDs (optional, default from base version when all empty)
	hypervisorPlatform := s.Config.HypervisorPlatform
	provisioningPlatform := s.Config.ProvisioningPlatform
	brokerPlatform := s.Config.BrokerPlatform
	if hypervisorPlatform == "" && provisioningPlatform == "" && brokerPlatform == "" {
		baseDetail, detailErr := helper.GetPlatformLayerRevisionDetailByName(s.Config.LayerName, s.Config.BaseVersionName)
		if detailErr != nil {
			ui.Errorf("Error getting base platform layer revision details: %v", detailErr)
			state.Put("error", detailErr)
			return multistep.ActionHalt
		}
		hypervisorPlatform = baseDetail.HypervisorPlatformTypeId
		provisioningPlatform = baseDetail.ProvisioningPlatformTypeId
		brokerPlatform = baseDetail.BrokerPlatformTypeId
	}
	// Translate "None" → ""
	if hypervisorPlatform == "None" {
		hypervisorPlatform = ""
	}
	if provisioningPlatform == "None" {
		provisioningPlatform = ""
	}
	if brokerPlatform == "None" {
		brokerPlatform = ""
	}

	platformLayerRevisionRequest := &elmsoap.CreatePlatformLayerRevision{
		Command: &elmsoap.CreatePlatformLayerRevisionCommand{
			CreateLayerRevisionCommand: &elmsoap.CreateLayerRevisionCommand{
				LayerId: platformLayerId,
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
				BaseLayerRevisionId:       &platformLayerRevisionId,
			},
			OsLayerRevisionId:          osLayerRevisionId,
			HypervisorPlatformTypeId:   hypervisorPlatform,
			ProvisioningPlatformTypeId: provisioningPlatform,
			BrokerPlatformTypeId:       brokerPlatform,
		},
	}
	revisionResp, err := helper.Client.CreatePlatformLayerRevision(platformLayerRevisionRequest)
	if err != nil {
		errMsg := fmt.Errorf("error calling RevisionPlatformLayer: %v", err)
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	if appErr := GetRevisionResultError(revisionResp.CreatePlatformLayerRevisionResult); appErr != nil {
		errMsg := fmt.Errorf("RevisionPlatformLayer failed: %s", FormatELMError(appErr))
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	workTicketId := revisionResp.CreatePlatformLayerRevisionResult.WorkTicketId
	UiSayf(ui, "Revision Platform Layer, WorkTicketId: %d", workTicketId)
	state.Put("WORK_TICKET_ID", workTicketId)
	return multistep.ActionContinue
}

func (s *StepRevisionPlatformLayering) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	_, hasError := state.GetOk("error")

	if !cancelled && !halted && !hasError {
		return
	}
	if !cancelled && s.Config.SkipCleanupOnFailure {
		log.Printf("[INFO] StepRevisionPlatformLayering.Cleanup: failure detected but cleanup_on_failure is false, skipping work ticket cancellation")
		return
	}

	workTicketId, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		log.Printf("[WARN] StepRevisionPlatformLayering.Cleanup: no WORK_TICKET_ID in state, cannot cancel work ticket")
		return
	}
	helper, ok := state.GetOk("soap_helper")
	if !ok {
		log.Printf("[WARN] StepRevisionPlatformLayering.Cleanup: soap_helper not in state, cannot cancel work ticket")
		return
	}
	if err := helper.(*elmsoap.SoapHelper).CancelWorkTicket(workTicketId.(int64)); err != nil {
		log.Printf("[WARN] StepRevisionPlatformLayering.Cleanup: failed to cancel work ticket %d: %v", workTicketId.(int64), err)
	} else {
		log.Printf("[INFO] StepRevisionPlatformLayering.Cleanup: cancelled work ticket %d", workTicketId.(int64))
	}
}
