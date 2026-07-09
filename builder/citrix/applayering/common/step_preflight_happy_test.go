// Copyright (c) Citrix, Inc.

package common

import (
	"testing"

	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// preflightHappyMock returns canonical positive responses for every Query*
// SOAP method any preflight strategy might call. Layers are named Win10 /
// MyPlatform / MyApp with a single revision DisplayedVersion="v1".
// A preflight target version of "v2-new" therefore reads as "doesn't exist"
// (the inner loop never matches), which is the happy-path expectation.
type preflightHappyMock struct {
	*baseApiSoap
}

func (m *preflightHappyMock) QueryOsLayers(*elmsoap.QueryOsLayers) (*elmsoap.QueryOsLayersResponse, error) {
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

func (m *preflightHappyMock) QueryOsLayerDetails(*elmsoap.QueryOsLayerDetails) (*elmsoap.QueryOsLayerDetailsResponse, error) {
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

func (m *preflightHappyMock) QueryPlatformLayers(*elmsoap.QueryPlatformLayers) (*elmsoap.QueryPlatformLayersResponse, error) {
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

func (m *preflightHappyMock) QueryPlatformLayerDetails(*elmsoap.QueryPlatformLayerDetails) (*elmsoap.QueryPlatformLayerDetailsResponse, error) {
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

func (m *preflightHappyMock) QueryApplicationLayers(*elmsoap.QueryApplicationLayers) (*elmsoap.QueryApplicationLayersResponse, error) {
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

func (m *preflightHappyMock) QueryApplicationLayerDetails(*elmsoap.QueryApplicationLayerDetails) (*elmsoap.QueryApplicationLayerDetailsResponse, error) {
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

// ---------------- happy-path tests ----------------

func TestCreatePlatformLayerPreflight_HappyPath(t *testing.T) {
	// LayerName="NewPlatform" doesn't match the mock's "MyPlatform" → existingId=0 → pass.
	s := &createPlatformLayerPreflight{}
	cfg := &alaconfig.Config{
		CreatePlatform: &alaconfig.CreatePlatformConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "NewPlatform",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestCreatePlatformLayerPreflight_LayerAlreadyExists(t *testing.T) {
	// LayerName="MyPlatform" matches the mock → existingId != 0 → error.
	s := &createPlatformLayerPreflight{}
	cfg := &alaconfig.Config{
		CreatePlatform: &alaconfig.CreatePlatformConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyPlatform",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err == nil {
		t.Errorf("Validate err = nil, want 'layer already exists'")
	}
}

func TestCreateAppLayerPreflight_HappyPath_WithPlatformLayer(t *testing.T) {
	// LayerName="NewApp" doesn't match "MyApp" → pass. Platform layer optional
	// branch is exercised because PlatformLayerName is set.
	s := &createAppLayerPreflight{}
	cfg := &alaconfig.Config{
		CreateApp: &alaconfig.CreateAppConfig{
			OsLayerName:              "Win10",
			OsLayerVersionName:       "v1",
			PlatformLayerName:        "MyPlatform",
			PlatformLayerVersionName: "v1",
			LayerName:                "NewApp",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestCreateAppLayerPreflight_LayerAlreadyExists(t *testing.T) {
	s := &createAppLayerPreflight{}
	cfg := &alaconfig.Config{
		CreateApp: &alaconfig.CreateAppConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyApp",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err == nil {
		t.Errorf("Validate err = nil, want 'layer already exists'")
	}
}

func TestRevisionOsLayerPreflight_HappyPath(t *testing.T) {
	// VersionName="v2-new" doesn't match the mock's "v1" → targetId=0 → pass.
	s := &revisionOsLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionOs: &alaconfig.RevisionOsConfig{
			LayerName:       "Win10",
			BaseVersionName: "v1",
			VersionName:     "v2-new",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestRevisionOsLayerPreflight_TargetVersionExists(t *testing.T) {
	s := &revisionOsLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionOs: &alaconfig.RevisionOsConfig{
			LayerName:       "Win10",
			BaseVersionName: "v1",
			VersionName:     "v1", // matches existing
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err == nil {
		t.Errorf("Validate err = nil, want 'target ... already exists'")
	}
}

func TestRevisionPlatformLayerPreflight_HappyPath(t *testing.T) {
	s := &revisionPlatformLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionPlatform: &alaconfig.RevisionPlatformConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyPlatform",
			BaseVersionName:    "v1",
			VersionName:        "v2-new",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestRevisionPlatformLayerPreflight_TargetVersionExists(t *testing.T) {
	s := &revisionPlatformLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionPlatform: &alaconfig.RevisionPlatformConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyPlatform",
			BaseVersionName:    "v1",
			VersionName:        "v1", // matches existing
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err == nil {
		t.Errorf("Validate err = nil, want 'target ... already exists'")
	}
}

func TestRevisionAppLayerPreflight_HappyPath_WithBase(t *testing.T) {
	// Covers the optional PlatformLayer branch and the BaseVersionName != ""
	// branch (GetAppLayerRevisionId for base lookup).
	s := &revisionAppLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionApp: &alaconfig.RevisionAppConfig{
			OsLayerName:              "Win10",
			OsLayerVersionName:       "v1",
			PlatformLayerName:        "MyPlatform",
			PlatformLayerVersionName: "v1",
			LayerName:                "MyApp",
			BaseVersionName:          "v1",
			VersionName:              "v2-new",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestRevisionAppLayerPreflight_HappyPath_AutoLatest(t *testing.T) {
	// BaseVersionName empty → the auto-latest branch is taken (no base lookup).
	s := &revisionAppLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionApp: &alaconfig.RevisionAppConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyApp",
			VersionName:        "v2-new",
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err != nil {
		t.Errorf("Validate err = %v, want nil", err)
	}
}

func TestRevisionAppLayerPreflight_TargetVersionExists(t *testing.T) {
	s := &revisionAppLayerPreflight{}
	cfg := &alaconfig.Config{
		RevisionApp: &alaconfig.RevisionAppConfig{
			OsLayerName:        "Win10",
			OsLayerVersionName: "v1",
			LayerName:          "MyApp",
			VersionName:        "v1", // matches existing
		},
	}
	helper := &elmsoap.SoapHelper{Client: &preflightHappyMock{baseApiSoap: &baseApiSoap{}}}
	if err := s.Validate(helper, cfg); err == nil {
		t.Errorf("Validate err = nil, want 'target ... already exists'")
	}
}
