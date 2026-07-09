// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// queryOsLayersErrMock fails on the first QueryOsLayers call — useful for
// driving any preflight strategy down its first error branch. All other SOAP
// methods panic via baseApiSoap.
type queryOsLayersErrMock struct {
	*baseApiSoap
	err   error
	calls int
}

func (m *queryOsLayersErrMock) QueryOsLayers(req *elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
	m.calls++
	return nil, m.err
}

// stateForStepPreflightCheck builds a state bag pre-populated with ui +
// a SoapHelper backed by the given Client mock.
func stateForStepPreflightCheck(client elmsoap.ApiSoap) (multistep.StateBag, *recordingUi) {
	state := new(multistep.BasicStateBag)
	ui := &recordingUi{}
	state.Put("ui", packersdk.Ui(ui))
	state.Put("soap_helper", &elmsoap.SoapHelper{Client: client})
	return state, ui
}

func TestNewPreflightStrategy_Dispatch(t *testing.T) {
	cases := []struct {
		op   elmsoap.ApplayeringOperationType
		want any
	}{
		{elmsoap.CREATE_PLATFORM_LAYER, &createPlatformLayerPreflight{}},
		{elmsoap.CREATE_APP_LAYER, &createAppLayerPreflight{}},
		{elmsoap.REVISION_OS_LAYER, &revisionOsLayerPreflight{}},
		{elmsoap.REVISION_PLATFORM_LAYER, &revisionPlatformLayerPreflight{}},
		{elmsoap.REVISION_APP_LAYER, &revisionAppLayerPreflight{}},
		{elmsoap.ApplayeringOperationType("UNKNOWN_OP"), &noOpPreflight{}},
		{elmsoap.ApplayeringOperationType(""), &noOpPreflight{}},
	}
	for _, tc := range cases {
		t.Run(string(tc.op), func(t *testing.T) {
			got := newPreflightStrategy(tc.op)
			if reflect.TypeOf(got) != reflect.TypeOf(tc.want) {
				t.Errorf("newPreflightStrategy(%q) = %T, want %T", tc.op, got, tc.want)
			}
		})
	}
}

func TestNoOpPreflight_Validate_ReturnsNil(t *testing.T) {
	s := &noOpPreflight{}
	if err := s.Validate(nil, nil); err != nil {
		t.Errorf("noOpPreflight.Validate = %v, want nil", err)
	}
}

func TestStepPreflightCheck_Run_NoOpType_Continues(t *testing.T) {
	// When no operation Config field is set, op type stays as the zero value
	// — newPreflightStrategy returns noOp, which always succeeds.
	// SOAP mock without overrides ensures no client call happens.
	mock := &queryOsLayersErrMock{baseApiSoap: &baseApiSoap{}}
	state, _ := stateForStepPreflightCheck(mock)

	step := &StepPreflightCheck{Config: &alaconfig.Config{}}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue (noOp should succeed without SOAP)", action)
	}
	if mock.calls != 0 {
		t.Errorf("QueryOsLayers calls = %d, want 0 (noOp should not call SOAP)", mock.calls)
	}
	if _, hasErr := state.GetOk("error"); hasErr {
		t.Errorf("error should not be set when noOp succeeds")
	}
}

func TestStepPreflightCheck_Run_OpTypeDispatch_HaltsOnStrategyError(t *testing.T) {
	// Each op type's strategy hits QueryOsLayers as its first SOAP call.
	// Driving that into error proves the right strategy was selected and that
	// the step halts + records the error.
	cases := []struct {
		name       string
		cfg        *alaconfig.Config
		errSubstr  string // expected substring in error message
	}{
		{
			name:      "CreateApp",
			cfg:       &alaconfig.Config{CreateApp: &alaconfig.CreateAppConfig{}},
			errSubstr: "preflight: validate OS layer version",
		},
		{
			name:      "CreatePlatform",
			cfg:       &alaconfig.Config{CreatePlatform: &alaconfig.CreatePlatformConfig{}},
			errSubstr: "preflight: validate OS layer version",
		},
		{
			name:      "RevisionOs",
			cfg:       &alaconfig.Config{RevisionOs: &alaconfig.RevisionOsConfig{}},
			errSubstr: "preflight: validate OS layer", // hits GetOsLayerId first
		},
		{
			name:      "RevisionPlatform",
			cfg:       &alaconfig.Config{RevisionPlatform: &alaconfig.RevisionPlatformConfig{}},
			errSubstr: "preflight: validate OS layer version",
		},
		{
			name:      "RevisionApp",
			cfg:       &alaconfig.Config{RevisionApp: &alaconfig.RevisionAppConfig{}},
			errSubstr: "preflight: validate OS layer version",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &queryOsLayersErrMock{
				baseApiSoap: &baseApiSoap{},
				err:         errors.New("fake SOAP failure"),
			}
			state, _ := stateForStepPreflightCheck(mock)
			step := &StepPreflightCheck{Config: tc.cfg}

			action := step.Run(context.Background(), state)

			if action != multistep.ActionHalt {
				t.Fatalf("action = %v, want ActionHalt", action)
			}
			if mock.calls < 1 {
				t.Errorf("QueryOsLayers should be called at least once; got %d", mock.calls)
			}
			err, ok := state.GetOk("error")
			if !ok {
				t.Fatalf("error should be set on strategy failure")
			}
			msg := err.(error).Error()
			if !strings.Contains(msg, tc.errSubstr) {
				t.Errorf("error = %q, want substring %q", msg, tc.errSubstr)
			}
		})
	}
}

func TestStepPreflightCheck_Cleanup_NoOp(t *testing.T) {
	step := &StepPreflightCheck{Config: &alaconfig.Config{}}
	step.Cleanup(new(multistep.BasicStateBag)) // must not panic
}
