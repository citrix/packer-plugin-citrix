// Copyright (c) Citrix, Inc.

package common

import (
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestCommHost_FixedHost(t *testing.T) {
	fn := CommHost("10.0.0.1")
	state := new(multistep.BasicStateBag)

	got, err := fn(state)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "10.0.0.1" {
		t.Errorf("got %q, want 10.0.0.1", got)
	}
}

func TestCommHost_FixedHost_IgnoresStateIp(t *testing.T) {
	fn := CommHost("10.0.0.1")
	state := new(multistep.BasicStateBag)
	state.Put("ip", "192.168.1.1")

	got, _ := fn(state)
	if got != "10.0.0.1" {
		t.Errorf("got %q, want 10.0.0.1 (state ip should be ignored when host is set)", got)
	}
}

func TestCommHost_EmptyHost_UsesStateIp(t *testing.T) {
	fn := CommHost("")
	state := new(multistep.BasicStateBag)
	state.Put("ip", "192.168.1.1")

	got, err := fn(state)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "192.168.1.1" {
		t.Errorf("got %q, want 192.168.1.1", got)
	}
}
