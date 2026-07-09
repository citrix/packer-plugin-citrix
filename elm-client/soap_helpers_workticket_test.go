// Copyright (c) Citrix, Inc.

package elmsoap

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ---------------- Mocks ----------------

type queryWorkTicketsMock struct {
	*baseApiSoap
	resp *QueryWorkTicketsResponse
	err  error
}

func (m *queryWorkTicketsMock) QueryWorkTickets(req *QueryWorkTickets) (*QueryWorkTicketsResponse, error) {
	return m.resp, m.err
}

type queryWorkTicketsAsPendingOpMock struct {
	*baseApiSoap
	resp *QueryWorkTicketsAsPendingOpResponse
	err  error
}

func (m *queryWorkTicketsAsPendingOpMock) QueryWorkTicketsAsPendingOp(req *QueryWorkTicketsAsPendingOp) (*QueryWorkTicketsAsPendingOpResponse, error) {
	return m.resp, m.err
}

type cancelWorkTicketsMock struct {
	*baseApiSoap
	resp *CancelWorkTicketsResponse
	err  error
}

func (m *cancelWorkTicketsMock) CancelWorkTickets(req *CancelWorkTickets) (*CancelWorkTicketsResponse, error) {
	return m.resp, m.err
}

// ---------------- helper builders ----------------

func workItemStatePtr(s WorkItemState) *WorkItemState { return &s }
func entityTypePtr(s string) *EntityType {
	et := EntityType(s)
	return &et
}
func strPtr(s string) *string { return &s }

// ---------------- GetWorkTicketIdByOperationTypeAndLayerName ----------------

func TestSoapHelper_GetWorkTicketIdByOperationTypeAndLayerName(t *testing.T) {
	t.Run("happy: returns first work item with ActionRequired state", func(t *testing.T) {
		mock := &queryWorkTicketsMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryWorkTicketsResponse{
				QueryWorkTicketsResult: &WorkTicketsResult{
					WorkTickets: &ArrayOfWorkTicketResult{
						WorkTicketResult: []*WorkTicketResult{
							{
								Id: 7,
								WorkItems: &ArrayOfWorkItemResult{
									WorkItemResult: []*WorkItemResult{
										{Id: 100, State: workItemStatePtr(WorkItemStatePending)},
										{Id: 101, State: workItemStatePtr(WorkItemStateActionRequired)},
									},
								},
							},
						},
					},
				},
			},
		}
		s := &SoapHelper{Client: mock}
		id, err := s.GetWorkTicketIdByOperationTypeAndLayerName(CREATE_APP_LAYER, "MyApp")
		if err != nil || id != 101 {
			t.Errorf("got (%d,%v), want (101,nil)", id, err)
		}
	})
	t.Run("soap error wraps", func(t *testing.T) {
		mock := &queryWorkTicketsMock{baseApiSoap: &baseApiSoap{}, err: errors.New("boom")}
		s := &SoapHelper{Client: mock}
		_, err := s.GetWorkTicketIdByOperationTypeAndLayerName(CREATE_APP_LAYER, "L")
		if err == nil || !strings.Contains(err.Error(), "QueryWorkTickets") {
			t.Errorf("err = %v, want wrapped SOAP err", err)
		}
	})
	t.Run("nil top-level result returns not found", func(t *testing.T) {
		mock := &queryWorkTicketsMock{baseApiSoap: &baseApiSoap{}, resp: &QueryWorkTicketsResponse{}}
		s := &SoapHelper{Client: mock}
		_, err := s.GetWorkTicketIdByOperationTypeAndLayerName(CREATE_APP_LAYER, "L")
		if err == nil || !strings.Contains(err.Error(), "no work ticket found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
	t.Run("WebResult error wraps", func(t *testing.T) {
		mock := &queryWorkTicketsMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryWorkTicketsResponse{
				QueryWorkTicketsResult: &WorkTicketsResult{WebResultBase: webResultErr("denied")},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetWorkTicketIdByOperationTypeAndLayerName(CREATE_APP_LAYER, "L")
		if err == nil || !strings.Contains(err.Error(), "denied") {
			t.Errorf("err = %v, want wrapped WebResult err", err)
		}
	})
	t.Run("nil WorkTickets returns not found", func(t *testing.T) {
		mock := &queryWorkTicketsMock{
			baseApiSoap: &baseApiSoap{},
			resp: &QueryWorkTicketsResponse{
				QueryWorkTicketsResult: &WorkTicketsResult{},
			},
		}
		s := &SoapHelper{Client: mock}
		_, err := s.GetWorkTicketIdByOperationTypeAndLayerName(CREATE_APP_LAYER, "L")
		if err == nil || !strings.Contains(err.Error(), "no work ticket found") {
			t.Errorf("err = %v, want not-found", err)
		}
	})
}

// ---------------- CancelWorkTicket WebResult-error branch (uplifts existing coverage) ----------------

func TestSoapHelper_CancelWorkTicket_WebResultError(t *testing.T) {
	mock := &cancelWorkTicketsMock{
		baseApiSoap: &baseApiSoap{},
		resp: &CancelWorkTicketsResponse{
			CancelWorkTicketsResult: &CancelWorkTicketsResult{WebResultBase: webResultErr("rejected")},
		},
	}
	s := &SoapHelper{Client: mock}
	err := s.CancelWorkTicket(99)
	if err == nil || !strings.Contains(err.Error(), "rejected") {
		t.Errorf("err = %v, want wrapped WebResult err", err)
	}
}

// ---------------- GetWorkTicketId branch uplift (currently 86.7%) ----------------

func TestSoapHelper_GetWorkTicketId_NilOperationResult(t *testing.T) {
	mock := &queryWorkTicketsAsPendingOpMock{
		baseApiSoap: &baseApiSoap{},
		resp: &QueryWorkTicketsAsPendingOpResponse{
			QueryWorkTicketsAsPendingOpResult: &PendingOperationResult{}, // OperationResult is nil
		},
	}
	s := &SoapHelper{Client: mock}
	_, err := s.GetWorkTicketId(nil, CREATE_APP_LAYER, "MyApp")
	if err == nil || !strings.Contains(err.Error(), "no operation result found") {
		t.Errorf("err = %v, want no-operation-result err", err)
	}
}

func TestSoapHelper_GetWorkTicketId_WebResultError(t *testing.T) {
	mock := &queryWorkTicketsAsPendingOpMock{
		baseApiSoap: &baseApiSoap{},
		resp: &QueryWorkTicketsAsPendingOpResponse{
			QueryWorkTicketsAsPendingOpResult: &PendingOperationResult{
				WebResultBase: webResultErr("perm denied"),
			},
		},
	}
	s := &SoapHelper{Client: mock}
	_, err := s.GetWorkTicketId(nil, CREATE_APP_LAYER, "MyApp")
	if err == nil || !strings.Contains(err.Error(), "perm denied") {
		t.Errorf("err = %v, want wrapped WebResult err", err)
	}
}

// ---------------- Raw-HTTP helpers via httptest ----------------
// queryWorkTicketByFilter, GetTaskByWorkTicketId, GetTaskCompletedFilter,
// GetTaskStateActiveFilter, GetIpByWorkTicketId all bypass the ApiSoap interface
// and POST directly to s.URL. We point at a TLS test server with InsecureSkipVerify.

const workTicketFilterRespTpl = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <QueryWorkTicketsAsPendingOpResponse xmlns="http://www.unidesk.com/">
      <QueryWorkTicketsAsPendingOpResult>
        <OperationResult>
          <WorkTickets>%s</WorkTickets>
        </OperationResult>
      </QueryWorkTicketsAsPendingOpResult>
    </QueryWorkTicketsAsPendingOpResponse>
  </soap:Body>
</soap:Envelope>`

// xmlWorkTicket renders a minimal <WorkTicketResult> string for stuffing into
// the template above. Fields are kept to what the helpers actually read.
func xmlWorkTicket(id int64, state string, items string) string {
	stateXML := ""
	if state != "" {
		stateXML = "<State>" + state + "</State>"
	}
	return `<WorkTicketResult><Id>` + itoa(id) + `</Id>` + stateXML + items + `</WorkTicketResult>`
}

func xmlWorkItem(id int64, itemType, state, status string) string {
	return `<WorkItems><WorkItemResult>` +
		`<Id>` + itoa(id) + `</Id>` +
		`<ItemType>` + itemType + `</ItemType>` +
		`<State>` + state + `</State>` +
		`<Status>` + status + `</Status>` +
		`</WorkItemResult></WorkItems>`
}

// itoa: tiny dep-free int64 → string (avoids importing strconv just for this).
func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var b [20]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

// soapServer starts a TLS test server returning the given response body on
// any POST. Returns the server URL and a teardown.
func soapServer(t *testing.T, body string) (*httptest.Server, *SoapHelper) {
	t.Helper()
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("server got %s, want POST", r.Method)
		}
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	s := &SoapHelper{URL: srv.URL, InsecureSkipVerify: true}
	return srv, s
}

func TestSoapHelper_queryWorkTicketByFilter_Happy(t *testing.T) {
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "ActionRequired", ""), 1)
	_, s := soapServer(t, body)

	got, err := s.GetTaskByWorkTicketId(42)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if got.Id != 42 {
		t.Errorf("Id = %d, want 42", got.Id)
	}
}

func TestSoapHelper_queryWorkTicketByFilter_NotFound(t *testing.T) {
	// Response contains a different work ticket, so helper should return notFoundErr.
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(1, "ActionRequired", ""), 1)
	_, s := soapServer(t, body)

	_, err := s.GetTaskByWorkTicketId(42)
	if !errors.Is(err, ErrWorkTicketNotInActiveFilter) {
		t.Errorf("err = %v, want ErrWorkTicketNotInActiveFilter", err)
	}
}

func TestSoapHelper_queryWorkTicketByFilter_HttpError(t *testing.T) {
	// Point at an unreachable URL — http.Client.Do should fail.
	s := &SoapHelper{URL: "https://127.0.0.1:1/never-listen", InsecureSkipVerify: true}
	if _, err := s.GetTaskByWorkTicketId(42); err == nil {
		t.Errorf("want HTTP err, got nil")
	}
}

func TestSoapHelper_queryWorkTicketByFilter_BadXML(t *testing.T) {
	_, s := soapServer(t, "<<<not-valid-xml")
	if _, err := s.GetTaskByWorkTicketId(42); err == nil {
		t.Errorf("want XML parse err, got nil")
	}
}

func TestSoapHelper_GetTaskCompletedFilter(t *testing.T) {
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "ActionRequired", ""), 1)
	_, s := soapServer(t, body)

	got, err := s.GetTaskCompletedFilter(42)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if got.Id != 42 {
		t.Errorf("Id = %d, want 42", got.Id)
	}
}

func TestSoapHelper_GetTaskCompletedFilter_NotFound(t *testing.T) {
	// Helper composes its own not-found error wrapping the ticket id.
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(1, "ActionRequired", ""), 1)
	_, s := soapServer(t, body)

	_, err := s.GetTaskCompletedFilter(42)
	if err == nil || !strings.Contains(err.Error(), "42") {
		t.Errorf("err = %v, want err mentioning ticket 42", err)
	}
}

func TestSoapHelper_GetTaskStateActiveFilter(t *testing.T) {
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "ActionRequired", ""), 1)
	_, s := soapServer(t, body)

	state, err := s.GetTaskStateActiveFilter(42)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if state != "ActionRequired" {
		t.Errorf("state = %q, want ActionRequired", state)
	}
}

func TestSoapHelper_GetTaskStateActiveFilter_Error(t *testing.T) {
	s := &SoapHelper{URL: "https://127.0.0.1:1/never-listen", InsecureSkipVerify: true}
	if _, err := s.GetTaskStateActiveFilter(42); err == nil {
		t.Errorf("want err, got nil")
	}
}

// ---------------- GetIpByWorkTicketId ----------------

func TestSoapHelper_GetIpByWorkTicketId_ExtractsFirstIp(t *testing.T) {
	// One ActionRequired item whose Status contains an IP address.
	items := `<WorkItems>` +
		`<WorkItemResult><Id>200</Id><ItemType>VirtualMachine</ItemType><State>ActionRequired</State><Status>VM ready at 10.20.30.40 listening</Status></WorkItemResult>` +
		`</WorkItems>`
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "ActionRequired", items), 1)
	_, s := soapServer(t, body)

	ip, err := s.GetIpByWorkTicketId(42)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if ip != "10.20.30.40" {
		t.Errorf("ip = %q, want 10.20.30.40", ip)
	}
}

func TestSoapHelper_GetIpByWorkTicketId_SkipsWorkItemType(t *testing.T) {
	// Helper skips items whose ItemType is exactly "WorkItem". Only the second item
	// (VirtualMachine + ActionRequired + IP in Status) should match.
	items := `<WorkItems>` +
		`<WorkItemResult><Id>100</Id><ItemType>WorkItem</ItemType><State>ActionRequired</State><Status>ignored 1.2.3.4</Status></WorkItemResult>` +
		`<WorkItemResult><Id>200</Id><ItemType>VirtualMachine</ItemType><State>ActionRequired</State><Status>VM at 9.8.7.6</Status></WorkItemResult>` +
		`</WorkItems>`
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "ActionRequired", items), 1)
	_, s := soapServer(t, body)

	ip, err := s.GetIpByWorkTicketId(42)
	if err != nil || ip != "9.8.7.6" {
		t.Errorf("got (%q,%v), want (9.8.7.6,nil)", ip, err)
	}
}

func TestSoapHelper_GetIpByWorkTicketId_NoActionRequiredReturnsEmpty(t *testing.T) {
	items := `<WorkItems>` +
		`<WorkItemResult><Id>200</Id><ItemType>VirtualMachine</ItemType><State>Pending</State><Status>VM at 9.8.7.6</Status></WorkItemResult>` +
		`</WorkItems>`
	body := strings.Replace(workTicketFilterRespTpl, "%s",
		xmlWorkTicket(42, "Pending", items), 1)
	_, s := soapServer(t, body)

	ip, err := s.GetIpByWorkTicketId(42)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
	if ip != "" {
		t.Errorf("ip = %q, want empty", ip)
	}
}

func TestSoapHelper_GetIpByWorkTicketId_TaskLookupError(t *testing.T) {
	s := &SoapHelper{URL: "https://127.0.0.1:1/never-listen", InsecureSkipVerify: true}
	if _, err := s.GetIpByWorkTicketId(42); err == nil {
		t.Errorf("want err, got nil")
	}
}

// Suppress unused-builder warnings for helpers we re-export for future tests.
var (
	_ = entityTypePtr
	_ = strPtr
)
