// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// revisionAppLayerHappyMock is the composite mock for
// StepRevisionAppLayering.Run. In addition to the canonical OS-layer stack,
// it mocks QueryApplicationLayers / QueryApplicationLayerDetails (for
// GetAppLayerId + GetAppLayerRevisionId) and the terminal
// CreateAppLayerRevision returning workTicketId=888.
type revisionAppLayerHappyMock struct {
	*baseApiSoap
	createErr        error
	createWebMessage string
}

func (m *revisionAppLayerHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *revisionAppLayerHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *revisionAppLayerHappyMock) QueryApplicationLayers(*elmsoap.QueryApplicationLayers) (*elmsoap.QueryApplicationLayersResponse, error) {
	return &elmsoap.QueryApplicationLayersResponse{
		QueryApplicationLayersResult: &elmsoap.AppLayersResult{
			AppLayers: &elmsoap.ArrayOfLayerEntitySummary{
				LayerEntitySummary: []*elmsoap.LayerEntitySummary{
					{EntityNode: &elmsoap.EntityNode{Id: 30, Name: "MyApp"}},
				},
			},
		},
	}, nil
}

func (m *revisionAppLayerHappyMock) QueryApplicationLayerDetails(*elmsoap.QueryApplicationLayerDetails) (*elmsoap.QueryApplicationLayerDetailsResponse, error) {
	return &elmsoap.QueryApplicationLayerDetailsResponse{
		QueryApplicationLayerDetailsResult: &elmsoap.AppLayerDetailsResult{
			LayerDetailsResultOfAppLayerRevisionDetail: &elmsoap.LayerDetailsResultOfAppLayerRevisionDetail{
				Revisions: &elmsoap.ArrayOfAppLayerRevisionDetail{
					AppLayerRevisionDetail: []*elmsoap.AppLayerRevisionDetail{
						{LayerRevisionDetail: &elmsoap.LayerRevisionDetail{Id: 400, DisplayedVersion: "v1"}},
					},
				},
			},
		},
	}, nil
}

func (m *revisionAppLayerHappyMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
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

func (m *revisionAppLayerHappyMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
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

func (m *revisionAppLayerHappyMock) CreateAppLayerRevision(*elmsoap.CreateAppLayerRevision) (*elmsoap.CreateAppLayerRevisionResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	res := &elmsoap.CreateRevisionResult{WorkTicketId: 888}
	if m.createWebMessage != "" {
		res.WebResultBase = &elmsoap.WebResultBase{
			ResultBase: &elmsoap.ResultBase{
				Error: &elmsoap.ApplicationError{Message: m.createWebMessage},
			},
		}
	}
	return &elmsoap.CreateAppLayerRevisionResponse{
		CreateAppLayerRevisionResult: res,
	}, nil
}

// BaseVersionName="v1" routes via GetAppLayerRevisionId (not the latest-auto
// path). VersionSizeGb=1 skips GetAppLayerRevisionSizeMiB. Platform-layer
// fields stay empty so GetPlatformLayerRevisionId is also skipped.
func newRevisionAppConfig() *RevisionAppLayerConfig {
	return &RevisionAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		BaseVersionName:             "v1",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v2",
		VersionSizeGb:               1,
	}
}

// ---------------- tests ----------------

func TestStepRevisionAppLayering_Run_HappyPath(t *testing.T) {
	mock := &revisionAppLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionAppLayering{Config: newRevisionAppConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	got, ok := state.GetOk("WORK_TICKET_ID")
	if !ok {
		t.Fatalf("WORK_TICKET_ID not set on success")
	}
	if got.(int64) != 888 {
		t.Errorf("WORK_TICKET_ID = %v, want 888", got)
	}
}

func TestStepRevisionAppLayering_Run_HaltsOnCreateSoapError(t *testing.T) {
	mock := &revisionAppLayerHappyMock{baseApiSoap: &baseApiSoap{}, createErr: errors.New("boom")}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionAppLayering{Config: newRevisionAppConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

func TestStepRevisionAppLayering_Run_HaltsOnCreateWebResultError(t *testing.T) {
	mock := &revisionAppLayerHappyMock{baseApiSoap: &baseApiSoap{}, createWebMessage: "denied"}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionAppLayering{Config: newRevisionAppConfig()}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

// TestStepRevisionAppLayering_Run_BadPrerequisiteFormat exercises the prereq
// parse branch: a malformed entry must halt before the SOAP create call.
func TestStepRevisionAppLayering_Run_BadPrerequisiteFormat(t *testing.T) {
	mock := &revisionAppLayerHappyMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	cfg := newRevisionAppConfig()
	cfg.PrerequisiteLayers = []string{"MissingColon"}
	step := &StepRevisionAppLayering{Config: cfg}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}
