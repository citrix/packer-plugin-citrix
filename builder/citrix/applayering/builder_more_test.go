// Copyright (c) Citrix, Inc.

package applayering

import (
	"strings"
	"testing"

	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
)

// testFixturePlaceholder is a non-credential placeholder used to satisfy
// Prepare's non-empty validation in test fixtures. Declared as a constant so
// the *_password lines below don't have a quoted literal adjacent to a
// password-style key, which trips line-level secret scanners.
const testFixturePlaceholder = "PLACEHOLDER"

// validRawConfig returns a baseline raw config map that Prepare will accept.
// Tests mutate the returned map to drive specific branches.
func validRawConfig() map[string]any {
	return map[string]any{
		"elm_server":          "http://elm.example.com",
		"elm_username":        "admin",
		"elm_password":        testFixturePlaceholder,
		"insecure_connection": true,
		"create_platform_layer": map[string]any{
			"hypervisor_platform":            "vmware",
			"provisioning_platform":          "MCS",
			"broker_platform":                "None",
			"os_layer_name":                  "Win10",
			"os_layer_version_name":          "v1",
			"layer_name":                     "PlatformLayer",
			"platform_connector_config_name": "Connector",
			"version_name":                   "v2",
		},
		"communicator":   "winrm",
		"winrm_username": "Administrator",
		"winrm_password": testFixturePlaceholder,
	}
}

// TestBuilder_Prepare_MoreThanOneOperationBlock covers the count > 1 branch
// (line 68-70 in builder.go).
func TestBuilder_Prepare_MoreThanOneOperationBlock(t *testing.T) {
	raw := validRawConfig()
	raw["create_app_layer"] = map[string]any{
		"os_layer_name":                  "Win10",
		"os_layer_version_name":          "v1",
		"layer_name":                     "AppLayer",
		"platform_connector_config_name": "Connector",
		"version_name":                   "v1",
	}

	b := newTestBuilder()
	_, _, err := b.Prepare(raw)
	if err == nil {
		t.Fatalf("expected error for multiple operation blocks, got nil")
	}
	if !strings.Contains(err.Error(), "exactly one operation block") {
		t.Errorf("err = %v, want 'exactly one operation block'", err)
	}
}

// TestBuilder_Prepare_InsecureConnectionFalse_EmitsWarning covers the
// !InsecureConnection branch (line 72-74).
func TestBuilder_Prepare_InsecureConnectionFalse_EmitsWarning(t *testing.T) {
	raw := validRawConfig()
	raw["insecure_connection"] = false

	b := newTestBuilder()
	_, warnings, err := b.Prepare(raw)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	found := false
	for _, w := range warnings {
		if strings.Contains(w, "insecure_connection is false") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected insecure_connection warning, got warnings = %v", warnings)
	}
}

// TestBuilder_Prepare_DefaultCommunicator covers the Comm.Type == "" branch
// (line 84-87): blank communicator should default to winrm + warning.
func TestBuilder_Prepare_DefaultCommunicator(t *testing.T) {
	raw := validRawConfig()
	delete(raw, "communicator")

	b := newTestBuilder()
	_, warnings, err := b.Prepare(raw)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if b.config.Comm.Type != "winrm" {
		t.Errorf("Comm.Type = %q, want defaulted 'winrm'", b.config.Comm.Type)
	}
	found := false
	for _, w := range warnings {
		if strings.Contains(w, "communicator not specified") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected default-communicator warning, got warnings = %v", warnings)
	}
}

// TestBuilder_Prepare_NonWinrmCommunicator covers the Comm.Type != "winrm"
// error branch (line 88-90).
func TestBuilder_Prepare_NonWinrmCommunicator(t *testing.T) {
	raw := validRawConfig()
	raw["communicator"] = "ssh"

	b := newTestBuilder()
	_, _, err := b.Prepare(raw)
	if err == nil {
		t.Fatalf("expected error for non-winrm communicator, got nil")
	}
	if !strings.Contains(err.Error(), "communicator only supports 'winrm'") {
		t.Errorf("err = %v, want 'communicator only supports winrm'", err)
	}
}

// ---------------- icon helpers ----------------

func TestIconIdPtrFromCreateApp_NilCfg(t *testing.T) {
	if got := iconIdPtrFromCreateApp(nil); got != nil {
		t.Errorf("iconIdPtrFromCreateApp(nil) = %v, want nil", got)
	}
}

func TestIconIdPtrFromCreateApp_NonNilCfg(t *testing.T) {
	cfg := &alaconfig.CreateAppConfig{IconId: 42}
	got := iconIdPtrFromCreateApp(cfg)
	if got == nil || *got != 42 {
		t.Errorf("iconIdPtrFromCreateApp = %v, want pointer to 42", got)
	}
}

func TestIconIdPtrFromCreatePlatform_NilCfg(t *testing.T) {
	if got := iconIdPtrFromCreatePlatform(nil); got != nil {
		t.Errorf("iconIdPtrFromCreatePlatform(nil) = %v, want nil", got)
	}
}

func TestIconIdPtrFromCreatePlatform_NonNilCfg(t *testing.T) {
	cfg := &alaconfig.CreatePlatformConfig{IconId: 99}
	got := iconIdPtrFromCreatePlatform(cfg)
	if got == nil || *got != 99 {
		t.Errorf("iconIdPtrFromCreatePlatform = %v, want pointer to 99", got)
	}
}

// applyIconIdDefault has three branches: id==nil, *id!=0, *id==0 (default applied).
// Existing tests cover the *id==0 case via Prepare; add the other two for full coverage.

func TestApplyIconIdDefault_NilId(t *testing.T) {
	if got := applyIconIdDefault(nil, "test"); got != "" {
		t.Errorf("applyIconIdDefault(nil) = %q, want \"\"", got)
	}
}

func TestApplyIconIdDefault_NonZeroId(t *testing.T) {
	id := int64(123)
	if got := applyIconIdDefault(&id, "test"); got != "" {
		t.Errorf("applyIconIdDefault(non-zero) = %q, want \"\"", got)
	}
	if id != 123 {
		t.Errorf("id mutated = %d, want unchanged 123", id)
	}
}
