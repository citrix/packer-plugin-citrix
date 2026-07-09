// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// queryWorkTicketsMock embeds *baseApiSoap and overrides only the SOAP method
// SoapHelper.GetWorkTicketId actually calls. Any unrelated SOAP call would
// panic via baseApiSoap, making accidental coupling loud.
type queryWorkTicketsMock struct {
	*baseApiSoap

	resp *elmsoap.QueryWorkTicketsAsPendingOpResponse
	err  error

	calls []*elmsoap.QueryWorkTicketsAsPendingOp
}

func (m *queryWorkTicketsMock) QueryWorkTicketsAsPendingOp(req *elmsoap.QueryWorkTicketsAsPendingOp) (*elmsoap.QueryWorkTicketsAsPendingOpResponse, error) {
	m.calls = append(m.calls, req)
	return m.resp, m.err
}

// strPtr is a one-liner for *string field population (gowsdl-generated types
// use *string everywhere via ArrayOfString.Astring).
func strPtr(s string) *string { return &s }

// workTicketResp builds a minimal valid response with one WorkTicketResult
// that matches the given layerName via ResourceArgs (the second match branch
// in SoapHelper.GetWorkTicketId — doesn't rely on operationToWorkTicket map).
func workTicketResp(id int64, layerName string) *elmsoap.QueryWorkTicketsAsPendingOpResponse {
	return &elmsoap.QueryWorkTicketsAsPendingOpResponse{
		QueryWorkTicketsAsPendingOpResult: &elmsoap.PendingOperationResult{
			OperationResult: &elmsoap.WorkTicketsResult{
				WorkTickets: &elmsoap.ArrayOfWorkTicketResult{
					WorkTicketResult: []*elmsoap.WorkTicketResult{
						{
							Id: id,
							ResourceArgs: &elmsoap.ArrayOfString{
								Astring: []*string{strPtr(layerName)},
							},
						},
					},
				},
			},
		},
	}
}

// stateForStepGetWorkTicketId builds a state bag pre-populated with ui and
// a SoapHelper whose Client is the given mock.
func stateForStepGetWorkTicketId(client elmsoap.ApiSoap) (multistep.StateBag, *recordingUi) {
	state := new(multistep.BasicStateBag)
	ui := &recordingUi{}
	state.Put("ui", packersdk.Ui(ui))
	state.Put("soap_helper", &elmsoap.SoapHelper{Client: client})
	return state, ui
}

func TestStepGetWorkTicketId_HappyPath(t *testing.T) {
	mock := &queryWorkTicketsMock{
		baseApiSoap: &baseApiSoap{},
		resp:        workTicketResp(12345, "MyAppLayer"),
	}
	state, _ := stateForStepGetWorkTicketId(mock)

	step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{
		OperationType: "CONNECT_CREATE_APP_VM_ONLY",
		LayerName:     "MyAppLayer",
	}}

	action := step.Run(context.Background(), state)

	if action != multistep.ActionContinue {
		t.Fatalf("action = %v, want ActionContinue", action)
	}
	if len(mock.calls) != 1 {
		t.Fatalf("QueryWorkTicketsAsPendingOp called %d times, want 1", len(mock.calls))
	}
	if got, _ := state.Get("WORK_TICKET_ID").(int64); got != 12345 {
		t.Errorf("WORK_TICKET_ID = %v, want 12345", got)
	}
	if _, ok := state.GetOk("error"); ok {
		t.Errorf("error should not be set on success")
	}
}

func TestStepGetWorkTicketId_AllSupportedOperationTypes(t *testing.T) {
	supported := []string{
		"CONNECT_REVISION_OS_VM_ONLY",
		"CONNECT_REVISION_PLATFORM_VM_ONLY",
		"CONNECT_CREATE_PLATFORM_VM_ONLY",
		"CONNECT_CREATE_APP_VM_ONLY",
		"CONNECT_REVISION_APP_VM_ONLY",
	}
	for _, op := range supported {
		t.Run(op, func(t *testing.T) {
			mock := &queryWorkTicketsMock{
				baseApiSoap: &baseApiSoap{},
				resp:        workTicketResp(1, "L"),
			}
			state, _ := stateForStepGetWorkTicketId(mock)
			step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{
				OperationType: op,
				LayerName:     "L",
			}}

			if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
				t.Errorf("action = %v, want ActionContinue (op should be accepted)", action)
			}
		})
	}
}

func TestStepGetWorkTicketId_InvalidOperationType(t *testing.T) {
	// Mock without override: any SOAP call would panic via baseApiSoap.
	// Validation should reject before reaching SOAP.
	mock := &queryWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
	state, _ := stateForStepGetWorkTicketId(mock)

	step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{
		OperationType: "CREATE_APP_LAYER", // valid layer op, but not a CONNECT_* op
		LayerName:     "MyAppLayer",
	}}

	action := step.Run(context.Background(), state)

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if len(mock.calls) != 0 {
		t.Errorf("SOAP should not have been called for invalid op type; got %d calls", len(mock.calls))
	}
	err, ok := state.GetOk("error")
	if !ok {
		t.Fatalf("error should be set on invalid op type")
	}
	if got := err.(error).Error(); got != "invalid operation type: CREATE_APP_LAYER" {
		t.Errorf("error message = %q, want %q", got, "invalid operation type: CREATE_APP_LAYER")
	}
}

func TestStepGetWorkTicketId_SoapError(t *testing.T) {
	wantErr := errors.New("ELM not reachable")
	mock := &queryWorkTicketsMock{
		baseApiSoap: &baseApiSoap{},
		err:         wantErr,
	}
	state, _ := stateForStepGetWorkTicketId(mock)

	step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{
		OperationType: "CONNECT_CREATE_APP_VM_ONLY",
		LayerName:     "MyAppLayer",
	}}

	action := step.Run(context.Background(), state)

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if len(mock.calls) != 1 {
		t.Errorf("SOAP should have been called once; got %d", len(mock.calls))
	}
	err, ok := state.GetOk("error")
	if !ok {
		t.Fatalf("error should be set on SOAP failure")
	}
	if !errors.Is(err.(error), wantErr) {
		t.Errorf("error = %v, want to wrap %v", err, wantErr)
	}
	if _, hasId := state.GetOk("WORK_TICKET_ID"); hasId {
		t.Errorf("WORK_TICKET_ID should NOT be set when SOAP fails")
	}
}

func TestStepGetWorkTicketId_NoMatchingWorkTicket(t *testing.T) {
	// Note: SoapHelper.GetWorkTicketId matches on TitleResourceId OR
	// ResourceArgs.Astring[0]. For CONNECT_* op types the operationToWorkTicket
	// map returns "" — so we must set TitleResourceId to a non-empty value to
	// exercise the no-match path (otherwise an empty TitleResourceId would
	// match the empty map lookup result and return the ticket).
	resp := &elmsoap.QueryWorkTicketsAsPendingOpResponse{
		QueryWorkTicketsAsPendingOpResult: &elmsoap.PendingOperationResult{
			OperationResult: &elmsoap.WorkTicketsResult{
				WorkTickets: &elmsoap.ArrayOfWorkTicketResult{
					WorkTicketResult: []*elmsoap.WorkTicketResult{
						{
							Id:              999,
							TitleResourceId: "SomeOtherDescription",
							ResourceArgs: &elmsoap.ArrayOfString{
								Astring: []*string{strPtr("OtherLayer")},
							},
						},
					},
				},
			},
		},
	}
	mock := &queryWorkTicketsMock{baseApiSoap: &baseApiSoap{}, resp: resp}
	state, _ := stateForStepGetWorkTicketId(mock)

	step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{
		OperationType: "CONNECT_CREATE_APP_VM_ONLY",
		LayerName:     "MyAppLayer",
	}}

	action := step.Run(context.Background(), state)

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt (no matching work ticket)", action)
	}
	if _, hasId := state.GetOk("WORK_TICKET_ID"); hasId {
		t.Errorf("WORK_TICKET_ID should NOT be set when no matching ticket found")
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("error should be set when no matching ticket found")
	}
}

func TestStepGetWorkTicketId_Cleanup_NoOp(t *testing.T) {
	step := &StepGetWorkTicketId{Config: &GetWorkTicketIdConfig{}}
	// Should not panic with an empty state.
	step.Cleanup(new(multistep.BasicStateBag))
}
