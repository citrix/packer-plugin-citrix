// Copyright (c) Citrix, Inc.

package elmsoap

import (
	"errors"
	"strings"
	"testing"
)

// --- Mocks: each composes one *Layers query method + its *LayerDetails companion.
// All embed *baseApiSoap so unrelated SOAP calls panic. ---

type osLayerRevisionMock struct {
	*baseApiSoap
	layersResp  *QueryOsLayersResponse
	layersErr   error
	detailsResp *QueryOsLayerDetailsResponse
	detailsErr  error
	detailCalls int
}

func (m *osLayerRevisionMock) QueryOsLayers(req *QueryOsLayers) (*QueryOsLayersResponse, error) {
	return m.layersResp, m.layersErr
}
func (m *osLayerRevisionMock) QueryOsLayerDetails(req *QueryOsLayerDetails) (*QueryOsLayerDetailsResponse, error) {
	m.detailCalls++
	return m.detailsResp, m.detailsErr
}

type platformLayerRevisionMock struct {
	*baseApiSoap
	layersResp  *QueryPlatformLayersResponse
	layersErr   error
	detailsResp *QueryPlatformLayerDetailsResponse
	detailsErr  error
	detailCalls int
}

func (m *platformLayerRevisionMock) QueryPlatformLayers(req *QueryPlatformLayers) (*QueryPlatformLayersResponse, error) {
	return m.layersResp, m.layersErr
}
func (m *platformLayerRevisionMock) QueryPlatformLayerDetails(req *QueryPlatformLayerDetails) (*QueryPlatformLayerDetailsResponse, error) {
	m.detailCalls++
	return m.detailsResp, m.detailsErr
}

type appLayerRevisionMock struct {
	*baseApiSoap
	layersResp  *QueryApplicationLayersResponse
	layersErr   error
	detailsResp *QueryApplicationLayerDetailsResponse
	detailsErr  error
	detailCalls int
}

func (m *appLayerRevisionMock) QueryApplicationLayers(req *QueryApplicationLayers) (*QueryApplicationLayersResponse, error) {
	return m.layersResp, m.layersErr
}
func (m *appLayerRevisionMock) QueryApplicationLayerDetails(req *QueryApplicationLayerDetails) (*QueryApplicationLayerDetailsResponse, error) {
	m.detailCalls++
	return m.detailsResp, m.detailsErr
}

// --- response builders ---

func osLayersResp(entries ...*LayerEntitySummary) *QueryOsLayersResponse {
	return &QueryOsLayersResponse{
		QueryOsLayersResult: &OsLayersResult{
			OsLayers: &ArrayOfLayerEntitySummary{LayerEntitySummary: entries},
		},
	}
}
func platformLayersResp(entries ...*LayerEntitySummary) *QueryPlatformLayersResponse {
	return &QueryPlatformLayersResponse{
		QueryPlatformLayersResult: &PlatformLayersResult{
			PlatformLayers: &ArrayOfLayerEntitySummary{LayerEntitySummary: entries},
		},
	}
}
func appLayersResp(entries ...*LayerEntitySummary) *QueryApplicationLayersResponse {
	return &QueryApplicationLayersResponse{
		QueryApplicationLayersResult: &AppLayersResult{
			AppLayers: &ArrayOfLayerEntitySummary{LayerEntitySummary: entries},
		},
	}
}

func osDetailsResp(revs ...*OsLayerRevisionDetail) *QueryOsLayerDetailsResponse {
	return &QueryOsLayerDetailsResponse{
		QueryOsLayerDetailsResult: &OsLayerDetailsResult{
			LayerDetailsResultOfOsLayerRevisionDetail: &LayerDetailsResultOfOsLayerRevisionDetail{
				Revisions: &ArrayOfOsLayerRevisionDetail{OsLayerRevisionDetail: revs},
			},
		},
	}
}
func platformDetailsResp(revs ...*PlatformLayerRevisionDetail) *QueryPlatformLayerDetailsResponse {
	return &QueryPlatformLayerDetailsResponse{
		QueryPlatformLayerDetailsResult: &PlatformLayerDetailsResult{
			LayerDetailsResultOfPlatformLayerRevisionDetail: &LayerDetailsResultOfPlatformLayerRevisionDetail{
				Revisions: &ArrayOfPlatformLayerRevisionDetail{PlatformLayerRevisionDetail: revs},
			},
		},
	}
}
func appDetailsResp(revs ...*AppLayerRevisionDetail) *QueryApplicationLayerDetailsResponse {
	return &QueryApplicationLayerDetailsResponse{
		QueryApplicationLayerDetailsResult: &AppLayerDetailsResult{
			LayerDetailsResultOfAppLayerRevisionDetail: &LayerDetailsResultOfAppLayerRevisionDetail{
				Revisions: &ArrayOfAppLayerRevisionDetail{AppLayerRevisionDetail: revs},
			},
		},
	}
}

func osRev(id int64, version string, sizeMegs int32) *OsLayerRevisionDetail {
	return &OsLayerRevisionDetail{LayerRevisionDetail: &LayerRevisionDetail{
		Id: id, DisplayedVersion: version, SizeMegs: sizeMegs,
	}}
}
func platRev(id int64, version string, sizeMegs int32) *PlatformLayerRevisionDetail {
	return &PlatformLayerRevisionDetail{LayerRevisionDetail: &LayerRevisionDetail{
		Id: id, DisplayedVersion: version, SizeMegs: sizeMegs,
	}}
}
func appRev(id int64, version string, sizeMegs int32, revisionNum int32) *AppLayerRevisionDetail {
	return &AppLayerRevisionDetail{LayerRevisionDetail: &LayerRevisionDetail{
		Id: id, DisplayedVersion: version, SizeMegs: sizeMegs, Revision: revisionNum,
	}}
}

// ---------------- GetOsLayerRevisionId (canonical pattern: all branches) ----------------

func TestSoapHelper_GetOsLayerRevisionId_AllBranches(t *testing.T) {
	t.Run("happy: layer + version match returns revision Id", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  osLayersResp(layerSummary(7, "Win10")),
			detailsResp: osDetailsResp(osRev(100, "v1", 0), osRev(101, "v2", 0)),
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetOsLayerRevisionId("Win10", "v2")
		if err != nil || id != 101 {
			t.Errorf("got (%d,%v), want (101,nil)", id, err)
		}
		if mock.detailCalls != 1 {
			t.Errorf("detail calls = %d, want 1", mock.detailCalls)
		}
	})

	t.Run("QueryOsLayers SOAP error wraps", func(t *testing.T) {
		mock := &osLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("boom")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("X", "v1"); err == nil || !strings.Contains(err.Error(), "QueryOsLayers") {
			t.Errorf("err = %v, want wrapped layers err", err)
		}
	})

	t.Run("nil top-level layers Result returns not found", func(t *testing.T) {
		mock := &osLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersResp: &QueryOsLayersResponse{}}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("X", "v1"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})

	t.Run("WebResult error on layers wraps", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp: &QueryOsLayersResponse{
				QueryOsLayersResult: &OsLayersResult{WebResultBase: webResultErr("bad cookie")},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("X", "v1"); err == nil || !strings.Contains(err.Error(), "bad cookie") {
			t.Errorf("err = %v, want wrapped WebResult err", err)
		}
	})

	t.Run("nil OsLayers sublist returns not found", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  &QueryOsLayersResponse{QueryOsLayersResult: &OsLayersResult{}},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("X", "v1"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})

	t.Run("QueryOsLayerDetails SOAP error wraps", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  osLayersResp(layerSummary(7, "Win10")),
			detailsErr:  errors.New("detail boom"),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("Win10", "v1"); err == nil || !strings.Contains(err.Error(), "QueryOsLayerDetails") {
			t.Errorf("err = %v, want wrapped details err", err)
		}
	})

	t.Run("WebResult error on details wraps", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  osLayersResp(layerSummary(7, "Win10")),
			detailsResp: &QueryOsLayerDetailsResponse{
				QueryOsLayerDetailsResult: &OsLayerDetailsResult{
					LayerDetailsResultOfOsLayerRevisionDetail: &LayerDetailsResultOfOsLayerRevisionDetail{
						WebResultBase: webResultErr("denied"),
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("Win10", "v1"); err == nil || !strings.Contains(err.Error(), "denied") {
			t.Errorf("err = %v, want wrapped details WebResult err", err)
		}
	})

	t.Run("matched layer but no matching version returns not found", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  osLayersResp(layerSummary(7, "Win10")),
			detailsResp: osDetailsResp(osRev(100, "v1", 0)),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionId("Win10", "v999"); err == nil || !strings.Contains(err.Error(), "v999") {
			t.Errorf("err = %v, want not-found mentioning version", err)
		}
	})
}

// ---------------- GetPlatformLayerRevisionId (smoke; structurally identical) ----------------

func TestSoapHelper_GetPlatformLayerRevisionId(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "P1")),
			detailsResp: platformDetailsResp(platRev(200, "v1", 0)),
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetPlatformLayerRevisionId("P1", "v1")
		if err != nil || id != 200 {
			t.Errorf("got (%d,%v), want (200,nil)", id, err)
		}
	})
	t.Run("layers soap error", func(t *testing.T) {
		mock := &platformLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionId("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("details soap error", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "P1")),
			detailsErr:  errors.New("x"),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionId("P1", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("not found", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "Other")),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionId("Missing", "v1"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}

// ---------------- GetAppLayerRevisionId (smoke) ----------------

func TestSoapHelper_GetAppLayerRevisionId(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "A1")),
			detailsResp: appDetailsResp(appRev(300, "v1", 0, 1)),
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetAppLayerRevisionId("A1", "v1")
		if err != nil || id != 300 {
			t.Errorf("got (%d,%v), want (300,nil)", id, err)
		}
	})
	t.Run("layers soap error", func(t *testing.T) {
		mock := &appLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerRevisionId("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("details soap error", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "A1")),
			detailsErr:  errors.New("x"),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerRevisionId("A1", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("not found", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "Other")),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerRevisionId("Missing", "v1"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}

// ---------------- Get*RevisionSizeMiB smoke ----------------

func TestSoapHelper_GetOsLayerRevisionSizeMiB(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &osLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  osLayersResp(layerSummary(7, "Win10")),
			detailsResp: osDetailsResp(osRev(100, "v1", 4096)),
		}
		s := &SoapHelper{Client: mock}
		size, err := s.GetOsLayerRevisionSizeMiB("Win10", "v1")
		if err != nil || size != 4096 {
			t.Errorf("got (%d,%v), want (4096,nil)", size, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &osLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetOsLayerRevisionSizeMiB("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
}

func TestSoapHelper_GetPlatformLayerRevisionSizeMiB(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "P1")),
			detailsResp: platformDetailsResp(platRev(200, "v1", 2048)),
		}
		s := &SoapHelper{Client: mock}
		size, err := s.GetPlatformLayerRevisionSizeMiB("P1", "v1")
		if err != nil || size != 2048 {
			t.Errorf("got (%d,%v), want (2048,nil)", size, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &platformLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionSizeMiB("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
}

func TestSoapHelper_GetAppLayerRevisionSizeMiB(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "A1")),
			detailsResp: appDetailsResp(appRev(300, "v1", 1024, 1)),
		}
		s := &SoapHelper{Client: mock}
		size, err := s.GetAppLayerRevisionSizeMiB("A1", "v1")
		if err != nil || size != 1024 {
			t.Errorf("got (%d,%v), want (1024,nil)", size, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &appLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerRevisionSizeMiB("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
}

// ---------------- GetLatestAppLayerRevision (picks highest Revision) ----------------

func TestSoapHelper_GetLatestAppLayerRevision(t *testing.T) {
	t.Run("happy: highest Revision number wins, regardless of order", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "A1")),
			detailsResp: appDetailsResp(
				appRev(100, "v1", 0, 1),
				appRev(300, "v3", 0, 3),
				appRev(200, "v2", 0, 2),
			),
		}
		s := &SoapHelper{Client: mock}
		rev, err := s.GetLatestAppLayerRevision("A1")
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if rev.Id != 300 || rev.DisplayedVersion != "v3" {
			t.Errorf("latest = (Id=%d, DisplayedVersion=%q), want (300, v3)", rev.Id, rev.DisplayedVersion)
		}
	})
	t.Run("soap error on layers", func(t *testing.T) {
		mock := &appLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetLatestAppLayerRevision("X"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("layer not found returns not-found", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "Other")),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetLatestAppLayerRevision("Missing"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
	t.Run("matched layer but no revisions returns helpful err", func(t *testing.T) {
		mock := &appLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  appLayersResp(layerSummary(9, "A1")),
			detailsResp: &QueryApplicationLayerDetailsResponse{
				QueryApplicationLayerDetailsResult: &AppLayerDetailsResult{
					LayerDetailsResultOfAppLayerRevisionDetail: &LayerDetailsResultOfAppLayerRevisionDetail{
						Revisions: &ArrayOfAppLayerRevisionDetail{}, // empty list
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetLatestAppLayerRevision("A1"); err == nil || !strings.Contains(err.Error(), "no revisions") {
			t.Errorf("err = %v, want no-revisions err", err)
		}
	})
}

// ---------------- GetPlatformLayerRevisionDetailByName ----------------

func TestSoapHelper_GetPlatformLayerRevisionDetailByName(t *testing.T) {
	t.Run("happy: returns matching revision detail pointer", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "P1")),
			detailsResp: platformDetailsResp(platRev(200, "v1", 0), platRev(201, "v2", 0)),
		}
		s := &SoapHelper{Client: mock}
		rev, err := s.GetPlatformLayerRevisionDetailByName("P1", "v2")
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if rev.Id != 201 {
			t.Errorf("Id = %d, want 201", rev.Id)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &platformLayerRevisionMock{baseApiSoap: &baseApiSoap{}, layersErr: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionDetailByName("X", "v1"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("version not found", func(t *testing.T) {
		mock := &platformLayerRevisionMock{
			baseApiSoap: &baseApiSoap{},
			layersResp:  platformLayersResp(layerSummary(8, "P1")),
			detailsResp: platformDetailsResp(platRev(200, "v1", 0)),
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerRevisionDetailByName("P1", "v999"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}
