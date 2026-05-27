package common

import (
	"context"
	"fmt"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type GetWorkTicketIdConfig struct {
	OperationType string `mapstructure:"operation_type"`
	LayerName     string `mapstructure:"layer_name"`
}

type StepGetWorkTicketId struct {
	Config *GetWorkTicketIdConfig
}

// Create unidesk soap client and connect to ELM server
func (s *StepGetWorkTicketId) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	UiSayf(ui, "StepGetWorkTicketId: OperationType: %s, LayerName: %s", s.Config.OperationType, s.Config.LayerName)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)

	operationType := elmsoap.ApplayeringOperationType(s.Config.OperationType)
	switch operationType {
	case elmsoap.CONNECT_REVISION_OS_VM_ONLY,
		elmsoap.CONNECT_REVISION_PLATFORM_VM_ONLY,
		elmsoap.CONNECT_CREATE_PLATFORM_VM_ONLY,
		elmsoap.CONNECT_CREATE_APP_VM_ONLY,
		elmsoap.CONNECT_REVISION_APP_VM_ONLY:
		// LayerName is the unified field for all ops
	default:
		err := fmt.Errorf("invalid operation type: %s", s.Config.OperationType)
		ui.Errorf("%v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	layerName := s.Config.LayerName

	workTicketId, err := helper.GetWorkTicketId(ui, operationType, layerName)
	if err != nil {
		ui.Errorf("Error getting work ticket ID for operation type %s and layer name %s: %v", s.Config.OperationType, layerName, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	UiSayf(ui, "Found work ticket ID: %d for operation type: %s and layer name: %s", workTicketId, s.Config.OperationType, layerName)
	state.Put("WORK_TICKET_ID", workTicketId)
	return multistep.ActionContinue
}

func (s *StepGetWorkTicketId) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
