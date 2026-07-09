// Copyright (c) Citrix, Inc.

package common

import (
	"strings"
	"testing"
)

// TestConnectConfig_Prepare exercises all four branches of ConnectConfig.Prepare:
// fully-populated config (zero errors) and each of the three required-field cases.
func TestConnectConfig_Prepare_HappyPath(t *testing.T) {
	c := &ConnectConfig{
		ELMServer:   "elm.example.com",
		ELMUsername: "u",
		ELMPassword: "p",
	}
	if errs := c.Prepare(); len(errs) != 0 {
		t.Errorf("Prepare = %v, want no errors", errs)
	}
}

func TestConnectConfig_Prepare_MissingUsername(t *testing.T) {
	c := &ConnectConfig{ELMServer: "x", ELMPassword: "p"}
	errs := c.Prepare()
	if len(errs) != 1 || !strings.Contains(errs[0].Error(), "elm_username") {
		t.Errorf("Prepare = %v, want single 'elm_username is required'", errs)
	}
}

func TestConnectConfig_Prepare_MissingPassword(t *testing.T) {
	c := &ConnectConfig{ELMServer: "x", ELMUsername: "u"}
	errs := c.Prepare()
	if len(errs) != 1 || !strings.Contains(errs[0].Error(), "elm_password") {
		t.Errorf("Prepare = %v, want single 'elm_password is required'", errs)
	}
}

func TestConnectConfig_Prepare_MissingServer(t *testing.T) {
	c := &ConnectConfig{ELMUsername: "u", ELMPassword: "p"}
	errs := c.Prepare()
	if len(errs) != 1 || !strings.Contains(errs[0].Error(), "elm_server") {
		t.Errorf("Prepare = %v, want single 'elm_server is required'", errs)
	}
}

func TestConnectConfig_Prepare_AllMissing(t *testing.T) {
	c := &ConnectConfig{}
	errs := c.Prepare()
	if len(errs) != 3 {
		t.Errorf("Prepare returned %d errors, want 3", len(errs))
	}
}

// TestStepConnect_Cleanup just ensures the no-op Cleanup doesn't panic.
func TestStepConnect_Cleanup_NoOp(t *testing.T) {
	s := &StepConnect{Config: &ConnectConfig{}}
	s.Cleanup(nil) // must not panic
}
