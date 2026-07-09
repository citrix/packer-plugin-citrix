// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// createAppLayerFullMock is the canonical-plus mock for exercising
// StepCreateAppLayer.Run with the optional platform-layer branch and the
// prerequisite-success branch. It mocks every Query* the deeper paths hit.
type createAppLayerFullMock struct {
	*baseApiSoap
}

func (m *createAppLayerFullMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *createAppLayerFullMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *createAppLayerFullMock) QueryPlatformLayers(*elmsoap.QueryPlatformLayers) (*elmsoap.QueryPlatformLayersResponse, error) {
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

func (m *createAppLayerFullMock) QueryPlatformLayerDetails(*elmsoap.QueryPlatformLayerDetails) (*elmsoap.QueryPlatformLayerDetailsResponse, error) {
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

func (m *createAppLayerFullMock) QueryApplicationLayers(*elmsoap.QueryApplicationLayers) (*elmsoap.QueryApplicationLayersResponse, error) {
	return &elmsoap.QueryApplicationLayersResponse{
		QueryApplicationLayersResult: &elmsoap.AppLayersResult{
			AppLayers: &elmsoap.ArrayOfLayerEntitySummary{
				LayerEntitySummary: []*elmsoap.LayerEntitySummary{
					{EntityNode: &elmsoap.EntityNode{Id: 30, Name: "PrereqApp"}},
				},
			},
		},
	}, nil
}

func (m *createAppLayerFullMock) QueryApplicationLayerDetails(*elmsoap.QueryApplicationLayerDetails) (*elmsoap.QueryApplicationLayerDetailsResponse, error) {
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

func (m *createAppLayerFullMock) QueryPlatformConnectorConfigSummary(*elmsoap.QueryPlatformConnectorConfigSummary) (*elmsoap.QueryPlatformConnectorConfigSummaryResponse, error) {
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

func (m *createAppLayerFullMock) QueryRemoteFileShares(*elmsoap.QueryRemoteFileShares) (*elmsoap.QueryRemoteFileSharesResponse, error) {
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

func (m *createAppLayerFullMock) CreateApplicationLayer(*elmsoap.CreateApplicationLayer) (*elmsoap.CreateApplicationLayerResponse, error) {
	return &elmsoap.CreateApplicationLayerResponse{
		CreateApplicationLayerResult: &elmsoap.CreateApplicationLayerResult{
			CreateLayerResult: &elmsoap.CreateLayerResult{WorkTicketId: 1111},
		},
	}, nil
}

func (m *createAppLayerFullMock) CreateAppLayerRevision(*elmsoap.CreateAppLayerRevision) (*elmsoap.CreateAppLayerRevisionResponse, error) {
	return &elmsoap.CreateAppLayerRevisionResponse{
		CreateAppLayerRevisionResult: &elmsoap.CreateRevisionResult{WorkTicketId: 2222},
	}, nil
}

// TestStepCreateAppLayer_Run_WithPlatformLayerAndPrereqs covers both the
// optional PlatformLayerName branch (lines ~54-63) and the prereq-success
// path (lines ~93-121) — the two biggest gaps in StepCreateAppLayer.Run.
func TestStepCreateAppLayer_Run_WithPlatformLayerAndPrereqs(t *testing.T) {
	mock := &createAppLayerFullMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		PlatformLayerName:           "MyPlatform",
		PlatformLayerVersionName:    "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
		PrerequisiteLayers:          []string{"PrereqApp:v1"},
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	if got, _ := state.GetOk("WORK_TICKET_ID"); got.(int64) != 1111 {
		t.Errorf("WORK_TICKET_ID = %v, want 1111", got)
	}
}

// TestStepCreateAppLayer_Run_DuplicatePrereq covers the "duplicate
// prerequisite layer name" branch (line 105-109).
func TestStepCreateAppLayer_Run_DuplicatePrereq(t *testing.T) {
	mock := &createAppLayerFullMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepCreateAppLayer{Config: &CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "MyApp",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v1",
		PrerequisiteLayers:          []string{"PrereqApp:v1", "PrereqApp:v2"},
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Errorf("action = %v, want ActionHalt", action)
	}
}

// TestStepRevisionAppLayering_Run_WithPlatformLayerAndPrereqs covers the
// optional PlatformLayer + prereq-success branches of StepRevisionAppLayering.Run.
func TestStepRevisionAppLayering_Run_WithPlatformLayerAndPrereqs(t *testing.T) {
	mock := &createAppLayerFullMock{baseApiSoap: &baseApiSoap{}}
	state := stateForLayerStepRun(mock)
	step := &StepRevisionAppLayering{Config: &RevisionAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		PlatformLayerName:           "MyPlatform",
		PlatformLayerVersionName:    "v1",
		LayerName:                   "PrereqApp",
		BaseVersionName:             "v1",
		PlatformConnectorConfigName: "MyConn",
		VersionName:                 "v2",
		VersionSizeGb:               1,
		PrerequisiteLayers:          []string{"PrereqApp:v1"},
	}}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Errorf("action = %v, want ActionContinue", action)
	}
	if got, _ := state.GetOk("WORK_TICKET_ID"); got.(int64) != 2222 {
		t.Errorf("WORK_TICKET_ID = %v, want 2222", got)
	}
}
