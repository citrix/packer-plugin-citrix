// Copyright (c) Citrix, Inc.

package applayering

import (
	"strings"
	"testing"
)

func TestArtifact_BuilderId(t *testing.T) {
	a := &Artifact{}
	if got := a.BuilderId(); got != BuilderId {
		t.Errorf("BuilderId() = %q, want %q", got, BuilderId)
	}
}

func TestArtifact_Files(t *testing.T) {
	a := &Artifact{}
	files := a.Files()
	if len(files) != 0 {
		t.Errorf("Files() = %v, want empty slice", files)
	}
}

func TestArtifact_Id(t *testing.T) {
	a := &Artifact{}
	if got := a.Id(); got != "" {
		t.Errorf("Id() = %q, want empty string", got)
	}
}

func TestArtifact_String(t *testing.T) {
	a := &Artifact{
		OperationType:          "CREATE_APP_LAYER",
		TargetLayerName:        "MyApp",
		TargetLayerVersionName: "v1.0",
	}
	got := a.String()
	for _, want := range []string{"CREATE_APP_LAYER", "MyApp", "v1.0"} {
		if !strings.Contains(got, want) {
			t.Errorf("String() = %q, missing %q", got, want)
		}
	}
}

func TestArtifact_State(t *testing.T) {
	a := &Artifact{
		StateData: map[string]interface{}{
			"generated_data": "abc",
			"other":          42,
		},
	}
	if got := a.State("generated_data"); got != "abc" {
		t.Errorf("State(generated_data) = %v, want abc", got)
	}
	if got := a.State("other"); got != 42 {
		t.Errorf("State(other) = %v, want 42", got)
	}
	if got := a.State("missing"); got != nil {
		t.Errorf("State(missing) = %v, want nil", got)
	}
}

func TestArtifact_Destroy(t *testing.T) {
	a := &Artifact{}
	if err := a.Destroy(); err != nil {
		t.Errorf("Destroy() = %v, want nil", err)
	}
}
