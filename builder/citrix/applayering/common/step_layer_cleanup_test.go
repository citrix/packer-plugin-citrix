// Copyright (c) Citrix, Inc.

package common

import (
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// cancelWorkTicketsMock embeds *baseApiSoap and overrides only
// Client.CancelWorkTickets — the single SOAP call SoapHelper.CancelWorkTicket
// performs.
type cancelWorkTicketsMock struct {
	*baseApiSoap
	err   error
	calls []*elmsoap.CancelWorkTickets
}

func (m *cancelWorkTicketsMock) CancelWorkTickets(req *elmsoap.CancelWorkTickets) (*elmsoap.CancelWorkTicketsResponse, error) {
	m.calls = append(m.calls, req)
	if m.err != nil {
		return nil, m.err
	}
	return &elmsoap.CancelWorkTicketsResponse{}, nil
}

// stateWithSoapHelper returns a state pre-populated with ui + a SoapHelper
// backed by the given mock. Callers fill additional keys.
func stateWithSoapHelper(mock elmsoap.ApiSoap) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("ui", packersdk.Ui(&recordingUi{}))
	state.Put("soap_helper", &elmsoap.SoapHelper{Client: mock})
	return state
}

// TestStepCreateAppLayer_Cleanup_Branches exercises every branch of the
// shared Cleanup logic — the other 4 layer steps copy this body verbatim, so
// covering all branches here proves the structure; smoke tests below confirm
// each step's wiring.
func TestStepCreateAppLayer_Cleanup_Branches(t *testing.T) {
	t.Run("no failure indicators returns early without SOAP", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
		state := stateWithSoapHelper(mock)
		state.Put("WORK_TICKET_ID", int64(42))

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{}}
		step.Cleanup(state)

		if len(mock.calls) != 0 {
			t.Errorf("CancelWorkTickets called %d times on success path, want 0", len(mock.calls))
		}
	})

	t.Run("SkipCleanupOnFailure halt returns early without SOAP", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
		state := stateWithSoapHelper(mock)
		state.Put("WORK_TICKET_ID", int64(42))
		state.Put(multistep.StateHalted, true)

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{SkipCleanupOnFailure: true}}
		step.Cleanup(state)

		if len(mock.calls) != 0 {
			t.Errorf("CancelWorkTickets called %d times when SkipCleanupOnFailure, want 0", len(mock.calls))
		}
	})

	t.Run("cancelled overrides SkipCleanupOnFailure and triggers cancel", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
		state := stateWithSoapHelper(mock)
		state.Put("WORK_TICKET_ID", int64(42))
		state.Put(multistep.StateCancelled, true)

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{SkipCleanupOnFailure: true}}
		step.Cleanup(state)

		if len(mock.calls) != 1 {
			t.Errorf("CancelWorkTickets called %d times when cancelled, want 1", len(mock.calls))
		}
	})

	t.Run("missing WORK_TICKET_ID returns without SOAP", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
		state := stateWithSoapHelper(mock)
		state.Put("error", errors.New("step failed"))
		// No WORK_TICKET_ID set.

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{}}
		step.Cleanup(state)

		if len(mock.calls) != 0 {
			t.Errorf("CancelWorkTickets called %d times without WORK_TICKET_ID, want 0", len(mock.calls))
		}
	})

	t.Run("missing soap_helper returns without SOAP", func(t *testing.T) {
		state := new(multistep.BasicStateBag)
		state.Put("ui", packersdk.Ui(&recordingUi{}))
		state.Put("error", errors.New("step failed"))
		state.Put("WORK_TICKET_ID", int64(42))
		// No soap_helper set.

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{}}
		step.Cleanup(state) // must not panic
	})

	t.Run("error triggers CancelWorkTickets and logs INFO on success", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
		state := stateWithSoapHelper(mock)
		state.Put("error", errors.New("step failed"))
		state.Put("WORK_TICKET_ID", int64(7))

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{}}
		step.Cleanup(state)

		if len(mock.calls) != 1 {
			t.Fatalf("CancelWorkTickets called %d times, want 1", len(mock.calls))
		}
		ids := mock.calls[0].Command.WorkTicketIds.Long
		if len(ids) != 1 || ids[0] != 7 {
			t.Errorf("cancelled ids = %v, want [7]", ids)
		}
	})

	t.Run("SOAP cancel failure does not panic (logged as WARN)", func(t *testing.T) {
		mock := &cancelWorkTicketsMock{
			baseApiSoap: &baseApiSoap{},
			err:         errors.New("SOAP cancel failed"),
		}
		state := stateWithSoapHelper(mock)
		state.Put("error", errors.New("step failed"))
		state.Put("WORK_TICKET_ID", int64(7))

		step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{}}
		step.Cleanup(state) // must not panic even when CancelWorkTickets errors

		if len(mock.calls) != 1 {
			t.Errorf("CancelWorkTickets called %d times, want 1 (attempt happens regardless)", len(mock.calls))
		}
	})
}

// Below: smoke tests for the other 4 layer steps. Each just verifies the
// step's Cleanup wires through to CancelWorkTickets on error, since the
// branching logic is identical (covered by the table above).

func TestStepCreatePlatformLayering_Cleanup_OnError_CancelsTicket(t *testing.T) {
	mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
	state := stateWithSoapHelper(mock)
	state.Put("error", errors.New("step failed"))
	state.Put("WORK_TICKET_ID", int64(42))

	step := &StepCreatePlatformLayering{Config: &CreatePlatformLayerConfig{}}
	step.Cleanup(state)

	if len(mock.calls) != 1 {
		t.Errorf("CancelWorkTickets calls = %d, want 1", len(mock.calls))
	}
}

func TestStepRevisionOsLayering_Cleanup_OnError_CancelsTicket(t *testing.T) {
	mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
	state := stateWithSoapHelper(mock)
	state.Put("error", errors.New("step failed"))
	state.Put("WORK_TICKET_ID", int64(42))

	step := &StepRevisionOsLayering{Config: &RevisionOsLayerConfig{}}
	step.Cleanup(state)

	if len(mock.calls) != 1 {
		t.Errorf("CancelWorkTickets calls = %d, want 1", len(mock.calls))
	}
}

func TestStepRevisionPlatformLayering_Cleanup_OnError_CancelsTicket(t *testing.T) {
	mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
	state := stateWithSoapHelper(mock)
	state.Put("error", errors.New("step failed"))
	state.Put("WORK_TICKET_ID", int64(42))

	step := &StepRevisionPlatformLayering{Config: &RevisionPlatformLayerConfig{}}
	step.Cleanup(state)

	if len(mock.calls) != 1 {
		t.Errorf("CancelWorkTickets calls = %d, want 1", len(mock.calls))
	}
}

func TestStepRevisionAppLayering_Cleanup_OnError_CancelsTicket(t *testing.T) {
	mock := &cancelWorkTicketsMock{baseApiSoap: &baseApiSoap{}}
	state := stateWithSoapHelper(mock)
	state.Put("error", errors.New("step failed"))
	state.Put("WORK_TICKET_ID", int64(42))

	step := &StepRevisionAppLayering{Config: &RevisionAppLayerConfig{}}
	step.Cleanup(state)

	if len(mock.calls) != 1 {
		t.Errorf("CancelWorkTickets calls = %d, want 1", len(mock.calls))
	}
}
