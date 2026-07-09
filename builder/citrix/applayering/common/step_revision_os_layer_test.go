// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// revisionOsLayerHappyMock is the composite mock for
// StepRevisionOsLayering.Run. The OS layer SOAP responses match the canonical
// layer-step fixtures (Id=10/Name="Win10", revision Id=100/v1); the terminal
// call is CreateOsLayerRevision returning workTicketId=555.
type revisionOsLayerHappyMock struct {
	*baseApiSoap
	createErr        error
	createWebMessage string
}

func (m *revisionOsLayerHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *revisionOsLayerHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *revisionOsLayerHappyMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
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

func (m *revisionOsLayerHappyMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
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

func (m *revisionOsLayerHappyMock) CreateOsLayerRevision(*elmsoap.CreateOsLayerRevision) (*elmsoap.CreateOsLayerRevisionResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	res := &elmsoap.CreateRevisionResult{WorkTicketId: 555}
	if m.createWebMessage != "" {
		res.WebResultBase = &elmsoap.WebResultBase{
			ResultBase: &elmsoap.ResultBase{
				Error: &elmsoap.ApplicationError{Message: m.createWebMessage},
			},
		}
	}
	return &elmsoap.CreateOsLayerRevisionResponse{
		CreateOsLayerRevisionResult: res,
	}, nil
}

// VersionSizeGb=1 keeps the test out of the GetOsLayerRevisionSizeMiB branch
// (which would re-issue QueryOsLayers/QueryOsLayerDetails just to look up size).
func newRevisionOsConfig() *RevisionOsLayerConfig {
	return &RevisionOsLayerConfig{
		LayerName:                   "Win10",
		BaseVersionName:             "v1",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v2",
		VersionSizeGb:               1,
	}
}

// ---------------- tests ----------------

func TestStepRevisionOsLayering_Run_HappyPath(t *testing.T) {
	mock := &revisionOsLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionOsLayering{Config: newRevisionOsConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	got, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		t.Fatalf("WORK_TICKET_ID not set on success")
	}
	if got.(int64) != 555 {
		t.Errorf("WORK_TICKET_ID = %v, want 555", got)
	}
}

func TestStepRevisionOsLayering_Run_HaltsOnCreateSoapError(t *testing.T) {
	mock := &revisionOsLayerHappyMock{baseApiSoap: &baseApiSoap{}, createErr: errors.New("boom")}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionOsLayering{Config: newRevisionOsConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

func TestStepRevisionOsLayering_Run_HaltsOnCreateWebResultError(t *testing.T) {
	mock := &revisionOsLayerHappyMock{baseApiSoap: &baseApiSoap{}, createWebMessage: "denied"}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionOsLayering{Config: newRevisionOsConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}
