package common

import (
	"context"
	"fmt"
	"log"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type CreatePlatformLayerConfig struct {
	// Platform type IDs
	HypervisorPlatform   string `mapstructure:"hypervisor_platform"`   // required
	ProvisioningPlatform string `mapstructure:"provisioning_platform"` // required, "None" to omit
	BrokerPlatform       string `mapstructure:"broker_platform"`       // required, "None" to omit
	// Command
	OsLayerName                 string `mapstructure:"os_layer_name"`
	OsLayerVersionName          string `mapstructure:"os_layer_version_name"`
	PackagingDiskFileName       string `mapstructure:"packaging_disk_file_name"` // optional, defaults to LayerName
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name"`

	// Command.LayerInfo
	LayerName string `mapstructure:"layer_name"`
	IconId    int64  `mapstructure:"icon_id"`
	// Command.RevisionInfo
	VersionName        string `mapstructure:"version_name"`
	VersionDescription string `mapstructure:"version_description"`
	VersionSizeGb      int32  `mapstructure:"version_size_gb"` // in GB, default 10
	// Command.Reason
	Comment          string `mapstructure:"comment"`
	SkipCleanupOnFailure bool   `mapstructure:"skip_cleanup_on_failure"`
}

type StepCreatePlatformLayering struct {
	Config *CreatePlatformLayerConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepCreatePlatformLayering) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
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
	// Convert GB to MiB, default 10 GB
	layerSizeMiB := s.Config.VersionSizeGb * 1024
	if layerSizeMiB == 0 {
		layerSizeMiB = 10 * 1024
	}
	// Translate "None" to empty string (omitempty will drop from SOAP)
	hypervisorPlatform := s.Config.HypervisorPlatform
	if hypervisorPlatform == "None" {
		hypervisorPlatform = ""
	}
	provisioningPlatform := s.Config.ProvisioningPlatform
	if provisioningPlatform == "None" {
		provisioningPlatform = ""
	}
	brokerPlatform := s.Config.BrokerPlatform
	if brokerPlatform == "None" {
		brokerPlatform = ""
	}
	platformLayerCreatequest := &elmsoap.CreatePlatformLayer{
		Command: &elmsoap.CreatePlatformLayerCommand{
			CreateLayerCommand: &elmsoap.CreateLayerCommand{
				OsLayerRevisionId: osLayerRevisionId,
				LayerInfo: &elmsoap.LayerInfo{
					Name:   s.Config.LayerName,
					IconId: s.Config.IconId,
				},
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
			},
			HypervisorPlatformTypeId:   hypervisorPlatform,
			ProvisioningPlatformTypeId: provisioningPlatform,
			BrokerPlatformTypeId:       brokerPlatform,
		},
	}
	createResp, err := helper.Client.CreatePlatformLayer(platformLayerCreatequest)
	if err != nil {
		errMsg := fmt.Errorf("error calling CreatePlatformLayer: %v", err)
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	if appErr := GetCreateLayerResultError(createResp.CreatePlatformLayerResult.CreateLayerResult); appErr != nil {
		errMsg := fmt.Errorf("CreatePlatformLayer failed: %s", FormatELMError(appErr))
		ui.Errorf("%v", errMsg)
		state.Put("error", errMsg)
		return multistep.ActionHalt
	}
	workTicketId := createResp.CreatePlatformLayerResult.CreateLayerResult.WorkTicketId
	UiSayf(ui, "Creating Platform Layer, WorkTicketId: %d", workTicketId)
	state.Put("WORK_TICKET_ID", workTicketId)
	return multistep.ActionContinue
}

func (s *StepCreatePlatformLayering) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	_, hasError := state.GetOk("error")

	if !cancelled && !halted && !hasError {
		return
	}
	if !cancelled && s.Config.SkipCleanupOnFailure {
		log.Printf("[INFO] StepCreatePlatformLayering.Cleanup: failure detected but cleanup_on_failure is false, skipping work ticket cancellation")
		return
	}

	workTicketId, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		log.Printf("[WARN] StepCreatePlatformLayering.Cleanup: no WORK_TICKET_ID in state, cannot cancel work ticket")
		return
	}
	helper, ok := state.GetOk("soap_helper")
	if !ok {
		log.Printf("[WARN] StepCreatePlatformLayering.Cleanup: soap_helper not in state, cannot cancel work ticket")
		return
	}
	if err := helper.(*elmsoap.SoapHelper).CancelWorkTicket(workTicketId.(int64)); err != nil {
		log.Printf("[WARN] StepCreatePlatformLayering.Cleanup: failed to cancel work ticket %d: %v", workTicketId.(int64), err)
	} else {
		log.Printf("[INFO] StepCreatePlatformLayering.Cleanup: cancelled work ticket %d", workTicketId.(int64))
	}
}
