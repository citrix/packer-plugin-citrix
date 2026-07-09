// Copyright (c) Citrix, Inc.

package applayering

import (
	"context"
	"io"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// ---------------- minimal test fakes ----------------

type discardUi struct{}

func (discardUi) Ask(string) (string, error)          { return "", nil }
func (discardUi) Askf(string, ...any) (string, error) { return "", nil }
func (discardUi) Say(string)                          {}
func (discardUi) Sayf(string, ...any)                 {}
func (discardUi) Message(string)                      {}
func (discardUi) Error(string)                        {}
func (discardUi) Errorf(string, ...any)               {}
func (discardUi) Machine(string, ...string)           {}
func (discardUi) TrackProgress(_ string, _, _ int64, stream io.ReadCloser) io.ReadCloser {
	return stream
}

// noopHook is a packer.Hook that does nothing -- runStrategy stores it in state
// but never reaches a hook-firing step in these tests.
type noopHook struct{}

func (noopHook) Run(_ context.Context, _ string, _ packersdk.Ui, _ packersdk.Communicator, _ interface{}) error {
	return nil
}

// fakeStrategy is a BuilderStrategy that returns the supplied steps and a fixed
// operation type / artifact. Tests use this to drive runStrategy without needing
// the real createApp/createPlatform/... factories.
type fakeStrategy struct {
	opType elmsoap.ApplayeringOperationType
	steps  []multistep.Step
}

func (f *fakeStrategy) OperationType() elmsoap.ApplayeringOperationType { return f.opType }
func (f *fakeStrategy) OperationSteps() []multistep.Step                { return f.steps }
func (f *fakeStrategy) Artifact(_ multistep.StateBag) *Artifact {
	return &Artifact{OperationType: string(f.opType), TargetLayerName: "TestLayer"}
}

// ---------------- ConfigSpec ----------------

func TestBuilder_ConfigSpec_NotNil(t *testing.T) {
	b := &Builder{}
	if b.ConfigSpec() == nil {
		t.Errorf("ConfigSpec() returned nil")
	}
}

// ---------------- Run / runStrategy ----------------
//
// Both Run and runStrategy build a step chain that begins with StepConnect.
// We point StepConnect at an unreachable address so its Login call fails fast;
// the multistep runner halts, runStrategy reads state["error"], and returns it.
// This covers the entry path, step-list assembly, the !winrm branch, generated_data
// population, runner invocation, and the error-return tail.

func newBuilderForRunTests() *Builder {
	cfg := alaconfig.Config{
		ELMServer:          "127.0.0.1:1", // unreachable: TCP port 1 + bare host -> https://127.0.0.1:1/...
		ELMUsername:        "u",
		ELMPassword:        "p",
		InsecureConnection: true,
	}
	// Leave Comm.Type empty so the winrm branch is skipped (avoids communicator
	// setup that's irrelevant to the code paths we're covering here).
	return &Builder{
		config: cfg,
		StrategyFactory: func(c *alaconfig.Config) BuilderStrategy {
			return &fakeStrategy{opType: elmsoap.CREATE_APP_LAYER, steps: nil}
		},
	}
}

func TestBuilder_Run_PropagatesConnectError(t *testing.T) {
	b := newBuilderForRunTests()
	_, err := b.Run(context.Background(), discardUi{}, noopHook{})
	if err == nil {
		t.Fatalf("Run err = nil, want StepConnect failure")
	}
}

func TestBuilder_runStrategy_PropagatesConnectError(t *testing.T) {
	b := newBuilderForRunTests()
	strategy := &fakeStrategy{opType: elmsoap.CREATE_APP_LAYER}
	_, err := b.runStrategy(context.Background(), discardUi{}, noopHook{}, strategy)
	if err == nil {
		t.Fatalf("runStrategy err = nil, want StepConnect failure")
	}
}

// ---------------- waitForIPStep ----------------

func TestBuilder_waitForIPStep_NotNil(t *testing.T) {
	b := &Builder{config: alaconfig.Config{WaitForIpTimeout: 0}}
	strategy := &fakeStrategy{opType: elmsoap.CREATE_APP_LAYER}
	if step := b.waitForIPStep(strategy); step == nil {
		t.Errorf("waitForIPStep returned nil")
	}
}
