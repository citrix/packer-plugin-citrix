// Copyright (c) Citrix, Inc.

package applayering

import (
	"strings"
	"testing"
)

// testFixturePlaceholder is a non-credential placeholder used to satisfy
// non-empty validation in test fixtures. Declared as a constant so lines
// below don't have a quoted literal adjacent to a *_password / PASSWORD key,
// which trips line-level secret scanners.
const testFixturePlaceholder = "PLACEHOLDER"

// ---------------------------------------------------------------------------
// ConfigSpec — trivial; just verify it returns a non-nil object spec.
// ---------------------------------------------------------------------------

func TestConfigSpec_NotNil(t *testing.T) {
	p := &Provisioner{}
	spec := p.ConfigSpec()
	if spec == nil {
		t.Errorf("ConfigSpec() returned nil")
	}
}

// ---------------------------------------------------------------------------
// uiSay / uiSayf / logUi.Say / logUi.Sayf — exercise the log/UI shims so
// coverage doesn't sit at 0% for these one-liners.
// ---------------------------------------------------------------------------

// recordingUI captures the last Say message so tests can assert on the
// timestamped prefix the helpers prepend.
type recordingUI struct {
	discardUI
	lastSay string
}

func (r *recordingUI) Say(msg string) { r.lastSay = msg }

func TestUiSay_PrependsTimestamp(t *testing.T) {
	ui := &recordingUI{}
	uiSay(ui, "hello")
	if !strings.HasSuffix(ui.lastSay, " hello") {
		t.Errorf("uiSay result = %q, want suffix \" hello\"", ui.lastSay)
	}
	if len(ui.lastSay) < len("2006-01-02 15:04:05.000 hello") {
		t.Errorf("uiSay result %q shorter than timestamped form", ui.lastSay)
	}
}

func TestUiSayf_FormatsAndPrependsTimestamp(t *testing.T) {
	ui := &recordingUI{}
	uiSayf(ui, "ticket %d state %s", 42, "ActionRequired")
	if !strings.HasSuffix(ui.lastSay, " ticket 42 state ActionRequired") {
		t.Errorf("uiSayf result = %q, want formatted suffix", ui.lastSay)
	}
}

func TestLogUi_SayAndSayfDoNotPanic(t *testing.T) {
	// logUi.Say / Sayf only call log.Printf; just make sure they run.
	lu := &logUi{Ui: &discardUI{}}
	lu.Say("trace message")
	lu.Sayf("trace %s %d", "x", 1)
}

// ---------------------------------------------------------------------------
// monitorTaskState — error branches that return before any SOAP call.
// generatedData is the only input; cover the four "missing key" failures.
// ---------------------------------------------------------------------------

func TestMonitorTaskState_MissingElmServer(t *testing.T) {
	p := &Provisioner{}
	err := p.monitorTaskState(&discardUI{}, map[string]any{})
	if err == nil || !strings.Contains(err.Error(), "elm server") {
		t.Errorf("err = %v, want elm server missing", err)
	}
}

func TestMonitorTaskState_MissingUserName(t *testing.T) {
	p := &Provisioner{}
	err := p.monitorTaskState(&discardUI{}, map[string]any{
		"ELM_SERVER": "elm.example.com",
	})
	if err == nil || !strings.Contains(err.Error(), "user name") {
		t.Errorf("err = %v, want user name missing", err)
	}
}

func TestMonitorTaskState_MissingPassword(t *testing.T) {
	p := &Provisioner{}
	err := p.monitorTaskState(&discardUI{}, map[string]any{
		"ELM_SERVER": "elm.example.com",
		"USER_NAME":  "admin",
	})
	if err == nil || !strings.Contains(err.Error(), "password") {
		t.Errorf("err = %v, want password missing", err)
	}
}

// TestMonitorTaskState_MissingWorkTicketId covers the 4th early-return
// branch. Reaching it requires passing ELM_TOKEN to skip the Login2 SOAP
// call (which would otherwise fail before we get to the WORK_TICKET_ID check).
func TestMonitorTaskState_MissingWorkTicketId(t *testing.T) {
	p := &Provisioner{}
	err := p.monitorTaskState(&discardUI{}, map[string]any{
		"ELM_SERVER": "elm.example.com",
		"USER_NAME":  "admin",
		"PASSWORD":   testFixturePlaceholder,
		"ELM_TOKEN":  "test-token",
	})
	if err == nil || !strings.Contains(err.Error(), "work ticket id") {
		t.Errorf("err = %v, want work ticket id missing", err)
	}
}
