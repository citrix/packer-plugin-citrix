// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// createAppLayerHappyMock is the canonical composite mock for
// StepCreateAppLayer.Run. It overrides every SOAP method touched by the
// happy path (no platform-layer revision, no prerequisites).
type createAppLayerHappyMock struct {
	*baseApiSoap
	createErr error
	createWebMessage string
}

func (m *createAppLayerHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
	return &elmsoap.QueryOsLayersResponse{
		QueryOsLayersResult: &elmsoap.OsLayersResult{
			OsLayers: &elmsoap.ArrayOfLayerEntitySummary{
				LayerEntitySummary: []*elmsoap.LayerEntitySummary{
					{EntityNode: &elmsoap.EntityNode{Id: 10, Name: "Win10"}},
				},
			},
		},
	}, nil
}

func (m *createAppLayerHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
	return &elmsoap.QueryOsLayerDetailsResponse{
		QueryOsLayerDetailsResult: &elmsoap.OsLayerDetailsResult{
			LayerDetailsResultOfOsLayerRevisionDetail: &elmsoap.LayerDetailsResultOfOsLayerRevisionDetail{
				Revisions: &elmsoap.ArrayOfOsLayerRevisionDetail{
					OsLayerRevisionDetail: []*elmsoap.OsLayerRevisionDetail{
						{LayerRevisionDetail: &elmsoap.LayerRevisionDetail{Id: 100, DisplayedVersion: "v1"}},
					},
				},
			},
		},
	}, nil
}

func (m *createAppLayerHappyMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
	return &elmsoap.QueryPlatformConnectorConfigSummaryResponse{
		QueryPlatformConnectorConfigSummaryResult: &elmsoap.PlatformConnectorConfigSummaryResult{
			Configurations: &elmsoap.ArrayOfPlatformConnectorConfigSummary{
				PlatformConnectorConfigSummary: []*elmsoap.PlatformConnectorConfigSummary{
					{Id: "ID-CONN", Name: "MyConn"},
				},
			},
		},
	}, nil
}

func (m *createAppLayerHappyMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
	return &elmsoap.QueryRemoteFileSharesResponse{
		QueryRemoteFileSharesResult: &elmsoap.RemoteFileSharesResult{
			RemoteShares: &elmsoap.ArrayOfRemoteFileShareSummary{
				RemoteFileShareSummary: []*elmsoap.RemoteFileShareSummary{
					{Id: 300, SharePath: `\\fs\share`},
				},
			},
		},
	}, nil
}

func (m *createAppLayerHappyMock) CreateApplicationLayer(*elmsoap.CreateApplicationLayer) (*elmsoap.CreateApplicationLayerResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	res := &elmsoap.CreateLayerResult{WorkTicketId: 999}
	if m.createWebMessage != "" {
		res.WebResultBase = &elmsoap.WebResultBase{
			ResultBase: &elmsoap.ResultBase{
				Error: &elmsoap.ApplicationError{Message: m.createWebMessage},
			},
		}
	}
	return &elmsoap.CreateApplicationLayerResponse{
		CreateApplicationLayerResult: &elmsoap.CreateApplicationLayerResult{
			CreateLayerResult: res,
		},
	}, nil
}

// ---------------- tests ----------------

func TestStepCreateAppLayer_Run_HappyPath(t *testing.T) {
	mock := &createAppLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	got, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		t.Fatalf("WORK_TICKET_ID not set on success")
	}
	if got.(int64) != 999 {
		t.Errorf("WORK_TICKET_ID = %v, want 999", got)
	}
	if _, ok := state.GetOk("error"); ok {
		t.Errorf("state[\"error\"] should not be set on success")
	}
}

func TestStepCreateAppLayer_Run_HaltsOnCreateSoapError(t *testing.T) {
	mock := &createAppLayerHappyMock{baseApiSoap: &baseApiSoap{}, createErr: errors.New("boom")}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("WORK_TICKET_ID"); ok {
		t.Errorf("WORK_TICKET_ID should not be set on Create SOAP failure")
	}
}

func TestStepCreateAppLayer_Run_HaltsOnCreateWebResultError(t *testing.T) {
	mock := &createAppLayerHappyMock{baseApiSoap: &baseApiSoap{}, createWebMessage: "denied"}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("WORK_TICKET_ID"); ok {
		t.Errorf("WORK_TICKET_ID should not be set on WebResult-error failure")
	}
}

// TestStepCreateAppLayer_Run_BadPrerequisiteFormat exercises the prereq parse
// branch: a malformed entry (missing ":") must halt before the SOAP create call.
func TestStepCreateAppLayer_Run_BadPrerequisiteFormat(t *testing.T) {
	mock := &createAppLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
		PrerequisiteLayers:          []string{"MissingColon"},
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("WORK_TICKET_ID"); ok {
		t.Errorf("WORK_TICKET_ID should not be set on bad prereq format")
	}
}
