// Copyright (c) Citrix, Inc.

package elmsoap

import (
	"errors"
	"strings"
	"testing"
)

// --- Mocks: one per Client method exercised. Each embeds *baseApiSoap so any
// SOAP call the helper does NOT make would panic loudly via baseApiSoap. ---

type queryPlatformConnectorConfigSummaryMock struct {
	*baseApiSoap
	resp *QueryPlatformConnectorConfigSummaryResponse
	err  error
}

func (m *queryPlatformConnectorConfigSummaryMock) QueryPlatformConnectorConfigSummary(req *QueryPlatformConnectorConfigSummary) (*QueryPlatformConnectorConfigSummaryResponse, error) {
	return m.resp, m.err
}

type queryOsLayersMock struct {
	*baseApiSoap
	resp *QueryOsLayersResponse
	err  error
}

func (m *queryOsLayersMock) QueryOsLayers(req *QueryOsLayers) (*QueryOsLayersResponse, error) {
	return m.resp, m.err
}

type queryPlatformLayersMock struct {
	*baseApiSoap
	resp *QueryPlatformLayersResponse
	err  error
}

func (m *queryPlatformLayersMock) QueryPlatformLayers(req *QueryPlatformLayers) (*QueryPlatformLayersResponse, error) {
	return m.resp, m.err
}

type queryApplicationLayersMock struct {
	*baseApiSoap
	resp *QueryApplicationLayersResponse
	err  error
}

func (m *queryApplicationLayersMock) QueryApplicationLayers(req *QueryApplicationLayers) (*QueryApplicationLayersResponse, error) {
	return m.resp, m.err
}

type queryRemoteFileSharesMock struct {
	*baseApiSoap
	resp *QueryRemoteFileSharesResponse
	err  error
}

func (m *queryRemoteFileSharesMock) QueryRemoteFileShares(req *QueryRemoteFileShares) (*QueryRemoteFileSharesResponse, error) {
	return m.resp, m.err
}

// --- helper builders ---

func layerSummary(id int64, name string) *LayerEntitySummary {
	return &LayerEntitySummary{EntityNode: &EntityNode{Id: id, Name: name}}
}

func webResultErr(msg string) *WebResultBase {
	return &WebResultBase{ResultBase: &ResultBase{Error: &ApplicationError{Message: msg}}}
}

// ---------------- GetPlatformConnectorConfigId ----------------

func TestSoapHelper_GetPlatformConnectorConfigId_Happy(t *testing.T) {
	mock := &queryPlatformConnectorConfigSummaryMock{
		baseApiSoap: &baseApiSoap{},
		resp: &QueryPlatformConnectorConfigSummaryResponse{
			QueryPlatformConnectorConfigSummaryResult: &PlatformConnectorConfigSummaryResult{
				Configurations: &ArrayOfPlatformConnectorConfigSummary{
					PlatformConnectorConfigSummary: []*PlatformConnectorConfigSummary{
						{Id: "ID-1", Name: "OtherConfig"},
						{Id: "ID-2", Name: "MyConfig"},
					},
				},
			},
		},
	}
	s := &SoapHelper{Client: mock}
	got, err := s.GetPlatformConnectorConfigId("MyConfig")
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if got != "ID-2" {
		t.Errorf("id = %q, want ID-2", got)
	}
}

func TestSoapHelper_GetPlatformConnectorConfigId_SoapError(t *testing.T) {
	mock := &queryPlatformConnectorConfigSummaryMock{baseApiSoap: &baseApiSoap{}, err: errors.New("boom")}
	s := &SoapHelper{Client: mock}
	_, err := s.GetPlatformConnectorConfigId("X")
	if err == nil || !strings.Contains(err.Error(), "QueryPlatformConnectorConfigSummary") {
		t.Errorf("err = %v, want wrapped SOAP error", err)
	}
}

func TestSoapHelper_GetPlatformConnectorConfigId_WebResultError(t *testing.T) {
	mock := &queryPlatformConnectorConfigSummaryMock{
		baseApiSoap: &baseApiSoap{},
		resp: &QueryPlatformConnectorConfigSummaryResponse{
			QueryPlatformConnectorConfigSummaryResult: &PlatformConnectorConfigSummaryResult{
				WebResultBase: webResultErr("denied"),
				Configurations: &ArrayOfPlatformConnectorConfigSummary{},
			},
		},
	}
	s := &SoapHelper{Client: mock}
	_, err := s.GetPlatformConnectorConfigId("X")
	if err == nil || !strings.Contains(err.Error(), "denied") {
		t.Errorf("err = %v, want wrapped WebResult error", err)
	}
}

func TestSoapHelper_GetPlatformConnectorConfigId_NotFound(t *testing.T) {
	mock := &queryPlatformConnectorConfigSummaryMock{
		baseApiSoap: &baseApiSoap{},
		resp: &QueryPlatformConnectorConfigSummaryResponse{
			QueryPlatformConnectorConfigSummaryResult: &PlatformConnectorConfigSummaryResult{
				Configurations: &ArrayOfPlatformConnectorConfigSummary{
					PlatformConnectorConfigSummary: []*PlatformConnectorConfigSummary{
						{Id: "ID-1", Name: "Foo"},
					},
				},
			},
		},
	}
	s := &SoapHelper{Client: mock}
	_, err := s.GetPlatformConnectorConfigId("MissingName")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Errorf("err = %v, want not-found error", err)
	}
}

// ---------------- GetOsLayerId (canonical pattern: covers all 6 branches) ----------------

func TestSoapHelper_GetOsLayerId_AllBranches(t *testing.T) {
	t.Run("happy: matching layer returns Id", func(t *testing.T) {
		mock := &queryOsLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryOsLayersResponse{
				QueryOsLayersResult: &OsLayersResult{
					OsLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{
							layerSummary(42, "Other"),
							layerSummary(7, "Win10"),
						},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetOsLayerId("Win10")
		if err != nil || id != 7 {
			t.Errorf("got (%d,%v), want (7,nil)", id, err)
		}
	})
	t.Run("soap error wraps", func(t *testing.T) {
		mock := &queryOsLayersMock{baseApiSoap: &baseApiSoap{}, err: errors.New("conn refused")}
		s := &SoapHelper{Client: mock}
		_, err := s.GetOsLayerId("X")
		if err == nil || !strings.Contains(err.Error(), "QueryOsLayers") {
			t.Errorf("err = %v, want wrapped SOAP error", err)
		}
	})
	t.Run("nil top-level Result returns not found", func(t *testing.T) {
		mock := &queryOsLayersMock{baseApiSoap: &baseApiSoap{}, resp: &QueryOsLayersResponse{}}
		s := &SoapHelper{Client: mock}
		_, err := s.GetOsLayerId("X")
		if err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
	t.Run("WebResult error wraps", func(t *testing.T) {
		mock := &queryOsLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryOsLayersResponse{
				QueryOsLayersResult: &OsLayersResult{WebResultBase: webResultErr("bad cookie")},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetOsLayerId("X")
		if err == nil || !strings.Contains(err.Error(), "bad cookie") {
			t.Errorf("err = %v, want wrapped WebResult error", err)
		}
	})
	t.Run("nil OsLayers sublist returns not found", func(t *testing.T) {
		mock := &queryOsLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryOsLayersResponse{
				QueryOsLayersResult: &OsLayersResult{}, // OsLayers nil
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetOsLayerId("X")
		if err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
	t.Run("name does not match any entry returns not found", func(t *testing.T) {
		mock := &queryOsLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryOsLayersResponse{
				QueryOsLayersResult: &OsLayersResult{
					OsLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{layerSummary(1, "WrongName")},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetOsLayerId("Win10")
		if err == nil || !strings.Contains(err.Error(), "Win10") {
			t.Errorf("err = %v, want not-found error mentioning name", err)
		}
	})
}

// ---------------- GetPlatformLayerId (smoke: structurally identical to GetOsLayerId) ----------------

func TestSoapHelper_GetPlatformLayerId_HappyAndError(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &queryPlatformLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryPlatformLayersResponse{
				QueryPlatformLayersResult: &PlatformLayersResult{
					PlatformLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{layerSummary(99, "PlatformX")},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetPlatformLayerId("PlatformX")
		if err != nil || id != 99 {
			t.Errorf("got (%d,%v), want (99,nil)", id, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &queryPlatformLayersMock{baseApiSoap: &baseApiSoap{}, err: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerId("X"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("not found", func(t *testing.T) {
		mock := &queryPlatformLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryPlatformLayersResponse{
				QueryPlatformLayersResult: &PlatformLayersResult{
					PlatformLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{layerSummary(1, "Other")},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetPlatformLayerId("Missing"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}

// ---------------- GetAppLayerId (smoke) ----------------

func TestSoapHelper_GetAppLayerId_HappyAndError(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		mock := &queryApplicationLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryApplicationLayersResponse{
				QueryApplicationLayersResult: &AppLayersResult{
					AppLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{layerSummary(33, "MyApp")},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetAppLayerId("MyApp")
		if err != nil || id != 33 {
			t.Errorf("got (%d,%v), want (33,nil)", id, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &queryApplicationLayersMock{baseApiSoap: &baseApiSoap{}, err: errors.New("x")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerId("X"); err == nil {
			t.Errorf("want err, got nil")
		}
	})
	t.Run("not found", func(t *testing.T) {
		mock := &queryApplicationLayersMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryApplicationLayersResponse{
				QueryApplicationLayersResult: &AppLayersResult{
					AppLayers: &ArrayOfLayerEntitySummary{
						LayerEntitySummary: []*LayerEntitySummary{layerSummary(1, "Other")},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetAppLayerId("Missing"); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}

// ---------------- GetDefaultFileShareId ----------------

func TestSoapHelper_GetDefaultFileShareId_AllBranches(t *testing.T) {
	t.Run("happy: returns first share Id", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryRemoteFileSharesResponse{
				QueryRemoteFileSharesResult: &RemoteFileSharesResult{
					RemoteShares: &ArrayOfRemoteFileShareSummary{
						RemoteFileShareSummary: []*RemoteFileShareSummary{
							{Id: 10, SharePath: `\\srv\share1`},
							{Id: 20, SharePath: `\\srv\share2`},
						},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetDefaultFileShareId()
		if err != nil || id != 10 {
			t.Errorf("got (%d,%v), want (10,nil)", id, err)
		}
	})
	t.Run("soap error", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{baseApiSoap: &baseApiSoap{}, err: errors.New("net err")}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetDefaultFileShareId(); err == nil || !strings.Contains(err.Error(), "QueryRemoteFileShares") {
			t.Errorf("err = %v, want wrapped soap err", err)
		}
	})
	t.Run("nil top-level result returns helpful err", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{baseApiSoap: &baseApiSoap{}, resp: &QueryRemoteFileSharesResponse{}}
		s := &SoapHelper{Client: mock}
		_, err := s.GetDefaultFileShareId()
		if err == nil || !strings.Contains(err.Error(), "no file shares found") {
			t.Errorf("err = %v, want no-file-shares err", err)
		}
	})
	t.Run("WebResult error wraps", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryRemoteFileSharesResponse{
				QueryRemoteFileSharesResult: &RemoteFileSharesResult{WebResultBase: webResultErr("perm denied")},
			},
		}
		s := &SoapHelper{Client: mock}
		if _, err := s.GetDefaultFileShareId(); err == nil || !strings.Contains(err.Error(), "perm denied") {
			t.Errorf("err = %v, want wrapped err", err)
		}
	})
	t.Run("nil RemoteShares returns helpful err", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryRemoteFileSharesResponse{
				QueryRemoteFileSharesResult: &RemoteFileSharesResult{},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetDefaultFileShareId()
		if err == nil || !strings.Contains(err.Error(), "no file shares found") {
			t.Errorf("err = %v, want no-file-shares err", err)
		}
	})
	t.Run("empty share list returns helpful err", func(t *testing.T) {
		mock := &queryRemoteFileSharesMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryRemoteFileSharesResponse{
				QueryRemoteFileSharesResult: &RemoteFileSharesResult{
					RemoteShares: &ArrayOfRemoteFileShareSummary{},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetDefaultFileShareId()
		if err == nil || !strings.Contains(err.Error(), "no file shares found") {
			t.Errorf("err = %v, want no-file-shares err", err)
		}
	})
}
