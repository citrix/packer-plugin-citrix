// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// revisionPlatformLayerHappyMock is the composite mock for
// StepRevisionPlatformLayering.Run. In addition to the canonical OS-layer
// stack, it also mocks QueryPlatformLayers / QueryPlatformLayerDetails
// (for GetPlatformLayerRevisionId + GetPlatformLayerId) and the terminal
// CreatePlatformLayerRevision returning workTicketId=666.
type revisionPlatformLayerHappyMock struct {
	*baseApiSoap
	createErr        error
	createWebMessage string
}

func (m *revisionPlatformLayerHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *revisionPlatformLayerHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *revisionPlatformLayerHappyMock) QueryPlatformLayers(*elmsoap.QueryPlatformLayers) (*elmsoap.QueryPlatformLayersResponse, error) {
	return &elmsoap.QueryPlatformLayersResponse{
		QueryPlatformLayersResult: &elmsoap.PlatformLayersResult{
			PlatformLayers: &elmsoap.ArrayOfLayerEntitySummary{
				LayerEntitySummary: []*elmsoap.LayerEntitySummary{
					{EntityNode: &elmsoap.EntityNode{Id: 20, Name: "MyPlatform"}},
				},
			},
		},
	}, nil
}

func (m *revisionPlatformLayerHappyMock) QueryPlatformLayerDetails(*elmsoap.QueryPlatformLayerDetails) (*elmsoap.QueryPlatformLayerDetailsResponse, error) {
	return &elmsoap.QueryPlatformLayerDetailsResponse{
		QueryPlatformLayerDetailsResult: &elmsoap.PlatformLayerDetailsResult{
			LayerDetailsResultOfPlatformLayerRevisionDetail: &elmsoap.LayerDetailsResultOfPlatformLayerRevisionDetail{
				Revisions: &elmsoap.ArrayOfPlatformLayerRevisionDetail{
					PlatformLayerRevisionDetail: []*elmsoap.PlatformLayerRevisionDetail{
						{LayerRevisionDetail: &elmsoap.LayerRevisionDetail{Id: 200, DisplayedVersion: "v1"}},
					},
				},
			},
		},
	}, nil
}

func (m *revisionPlatformLayerHappyMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
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

func (m *revisionPlatformLayerHappyMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
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

func (m *revisionPlatformLayerHappyMock) CreatePlatformLayerRevision(*elmsoap.CreatePlatformLayerRevision) (*elmsoap.CreatePlatformLayerRevisionResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	res := &elmsoap.CreateRevisionResult{WorkTicketId: 666}
	if m.createWebMessage != "" {
		res.WebResultBase = &elmsoap.WebResultBase{
			ResultBase: &elmsoap.ResultBase{
				Error: &elmsoap.ApplicationError{Message: m.createWebMessage},
			},
		}
	}
	return &elmsoap.CreatePlatformLayerRevisionResponse{
		CreatePlatformLayerRevisionResult: res,
	}, nil
}

// VersionSizeGb=1 skips the GetPlatformLayerRevisionSizeMiB branch and the
// "None" platform types skip the GetPlatformLayerRevisionDetailByName lookup.
func newRevisionPlatformConfig() *RevisionPlatformLayerConfig {
	return &RevisionPlatformLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyPlatform",
		BaseVersionName:             "v1",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v2",
		VersionSizeGb:               1,
		HypervisorPlatform:          "None",
		ProvisioningPlatform:        "None",
		BrokerPlatform:              "None",
	}
}

// ---------------- tests ----------------

func TestStepRevisionPlatformLayering_Run_HappyPath(t *testing.T) {
	mock := &revisionPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionPlatformLayering{Config: newRevisionPlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	got, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		t.Fatalf("WORK_TICKET_ID not set on success")
	}
	if got.(int64) != 666 {
		t.Errorf("WORK_TICKET_ID = %v, want 666", got)
	}
}

func TestStepRevisionPlatformLayering_Run_HaltsOnCreateSoapError(t *testing.T) {
	mock := &revisionPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}, createErr: errors.New("boom")}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionPlatformLayering{Config: newRevisionPlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

func TestStepRevisionPlatformLayering_Run_HaltsOnCreateWebResultError(t *testing.T) {
	mock := &revisionPlatformLayerHappyMock{baseApiSoap: &baseApiSoap{}, createWebMessage: "denied"}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionPlatformLayering{Config: newRevisionPlatformConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}
