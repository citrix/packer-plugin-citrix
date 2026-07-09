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

// stateForLayerStepRun returns state pre-populated with ui + SoapHelper
// backed by the given Client mock. Used by all the layer-step Run tests.
func stateForLayerStepRun(client elmsoap.ApiSoap) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("ui", packersdk.Ui(&recordingUi{}))
	state.Put("soap_helper", &elmsoap.SoapHelper{Client: client})
	return state
}

// All five layer-step Run methods hit Client.QueryOsLayers as their first SOAP
// call (via helper.GetOsLayerRevisionId or helper.GetOsLayerId). Driving that
// into error covers the state setup + first error branch in each Run. Deeper
// happy-path traversal stays deferred to Wave 3.

func TestStepCreateAppLayer_Run_HaltsOnFirstSoapError(t *testing.T) {
	mock := &queryOsLayersErrMock{
		baseApiSoap: &baseApiSoap{},
		err:         errors.New("ELM unreachable"),
	}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:        "Win10",
		OsLayerVersionName: "v1",
		LayerName:          "MyApp",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("state[\"error\"] should be set on first SOAP failure")
	}
	if _, ok := state.GetOk("WORK_TICKET_ID"); ok {
		t.Errorf("WORK_TICKET_ID should not be set on failure")
	}
}

func TestStepCreatePlatformLayering_Run_HaltsOnFirstSoapError(t *testing.T) {
	mock := &queryOsLayersErrMock{
		baseApiSoap: &baseApiSoap{},
		err:         errors.New("ELM unreachable"),
	}
	state := stateForLayerStepRun(mock)
	step := &StepCreatePlatformLayering{Config: &CreatePlatformLayerConfig{
		OsLayerName:        "Win10",
		OsLayerVersionName: "v1",
		LayerName:          "MyPlatform",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("state[\"error\"] should be set on first SOAP failure")
	}
}

func TestStepRevisionOsLayering_Run_HaltsOnFirstSoapError(t *testing.T) {
	mock := &queryOsLayersErrMock{
		baseApiSoap: &baseApiSoap{},
		err:         errors.New("ELM unreachable"),
	}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionOsLayering{Config: &RevisionOsLayerConfig{
		LayerName:       "Win10",
		BaseVersionName: "v1",
		VersionName:     "v2",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("state[\"error\"] should be set on first SOAP failure")
	}
}

func TestStepRevisionPlatformLayering_Run_HaltsOnFirstSoapError(t *testing.T) {
	mock := &queryOsLayersErrMock{
		baseApiSoap: &baseApiSoap{},
		err:         errors.New("ELM unreachable"),
	}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionPlatformLayering{Config: &RevisionPlatformLayerConfig{
		OsLayerName:        "Win10",
		OsLayerVersionName: "v1",
		LayerName:          "MyPlatform",
		BaseVersionName:    "v1",
		VersionName:        "v2",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("state[\"error\"] should be set on first SOAP failure")
	}
}

func TestStepRevisionAppLayering_Run_HaltsOnFirstSoapError(t *testing.T) {
	mock := &queryOsLayersErrMock{
		baseApiSoap: &baseApiSoap{},
		err:         errors.New("ELM unreachable"),
	}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionAppLayering{Config: &RevisionAppLayerConfig{
		OsLayerName:        "Win10",
		OsLayerVersionName: "v1",
		LayerName:          "MyApp",
		VersionName:        "v2",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Errorf("state[\"error\"] should be set on first SOAP failure")
	}
}
