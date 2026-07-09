// Copyright (c) Citrix, Inc.

package common

import (
	"regexp"
	"testing"
)

// timestampPattern matches the prefix "YYYY-MM-DD HH:MM:SS.mmm ".
var timestampPattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} `)

func TestTimestampedUi_Say(t *testing.T) {
	inner := &recordingUi{}
	tu := &TimestampedUi{Ui: inner}

	tu.Say("hello")

	if len(inner.said) != 1 {
		t.Fatalf("len(said) = %d, want 1", len(inner.said))
	}
	got := inner.said[0]
	if !timestampPattern.MatchString(got) {
		t.Errorf("Say output %q missing timestamp prefix", got)
	}
	if got[len(got)-5:] != "hello" {
		t.Errorf("Say output %q does not end with 'hello'", got)
	}
}

func TestTimestampedUi_Sayf(t *testing.T) {
	inner := &recordingUi{}
	tu := &TimestampedUi{Ui: inner}

	tu.Sayf("count=%d name=%s", 7, "abc")

	if len(inner.said) != 1 {
		t.Fatalf("len(said) = %d, want 1 (Sayf delegates to underlying Say)", len(inner.said))
	}
	got := inner.said[0]
	if !timestampPattern.MatchString(got) {
		t.Errorf("Sayf output %q missing timestamp prefix", got)
	}
	if got[len(got)-len("count=7 name=abc"):] != "count=7 name=abc" {
		t.Errorf("Sayf output %q does not end with formatted message", got)
	}
}
