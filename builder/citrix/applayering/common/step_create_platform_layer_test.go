// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// createPlatformLayerHappyMock is the composite mock for the
// StepCreatePlatformLayering.Run path. The first three SOAP methods are
// identical to createAppLayerHappyMock; only the terminal Create call differs.
type createPlatformLayerHappyMock struct {
	*baseApiSoap
	createErr        error
	createWebMessage string
}

func (m *createPlatformLayerHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *createPlatformLayerHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *createPlatformLayerHappyMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
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

func (m *createPlatformLayerHappyMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
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

func (m *createPlatformLayerHappyMock) CreatePlatformLayer(*elmsoap.CreatePlatformLayer) (*elmsoap.CreatePlatformLayerResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	res := &elmsoap.CreateLayerResult{WorkTicketId: 777}
	if m.createWebMessage != "" {
		res.WebResultBase = &elmsoap.WebResultBase{
			ResultBase: &elmsoap.ResultBase{
				Error: &elmsoap.ApplicationError{Message: m.createWebMessage},
			},
		}
	}
	return &elmsoap.CreatePlatformLayerResponse{
		CreatePlatformLayerResult: &elmsoap.CreatePlatformLayerResult{
			CreateLayerResult: res,
		},
	}, nil
}

func newCreatePlatformConfig() *CreatePlatformLayerConfig {
	return &CreatePlatformLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyPlatform",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
		HypervisorPlatform:          "None",
		ProvisioningPlatform:        "None",
		BrokerPlatform:              "None",
	}
}

// ---------------- tests ----------------

func TestStepCreatePlatformLayering_Run_HappyPath(t *testing.T) {
	mock := &createPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepCreatePlatformLayering{Config: newCreatePlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	got, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		t.Fatalf("WORK_TICKET_ID not set on success")
	}
	if got.(int64) != 777 {
		t.Errorf("WORK_TICKET_ID = %v, want 777", got)
	}
}

func TestStepCreatePlatformLayering_Run_HaltsOnCreateSoapError(t *testing.T) {
	mock := &createPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}, createErr: errors.New("boom")}
	state := stateForLayerStepRun(mock)
	step := &StepCreatePlatformLayering{Config: newCreatePlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

func TestStepCreatePlatformLayering_Run_HaltsOnCreateWebResultError(t *testing.T) {
	mock := &createPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}, createWebMessage: "denied"}
	state := stateForLayerStepRun(mock)
	step := &StepCreatePlatformLayering{Config: newCreatePlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}
