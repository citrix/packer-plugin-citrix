// Copyright (c) Citrix, Inc.

package common

import (
	"strings"
	"testing"
)

func TestUiSay(t *testing.T) {
	ui := &recordingUi{}

	UiSay(ui, "hello world")

	if len(ui.said) != 1 {
		t.Fatalf("len(said) = %d, want 1", len(ui.said))
	}
	got := ui.said[0]
	if !timestampPattern.MatchString(got) {
		t.Errorf("UiSay output %q missing timestamp prefix", got)
	}
	if !strings.HasSuffix(got, "hello world") {
		t.Errorf("UiSay output %q does not end with 'hello world'", got)
	}
}

func TestUiSayf(t *testing.T) {
	ui := &recordingUi{}

	UiSayf(ui, "id=%d name=%s", 42, "foo")

	if len(ui.said) != 1 {
		t.Fatalf("len(said) = %d, want 1", len(ui.said))
	}
	got := ui.said[0]
	if !timestampPattern.MatchString(got) {
		t.Errorf("UiSayf output %q missing timestamp prefix", got)
	}
	if !strings.HasSuffix(got, "id=42 name=foo") {
		t.Errorf("UiSayf output %q does not end with formatted message", got)
	}
}
