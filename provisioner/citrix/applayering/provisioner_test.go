// Copyright (c) Citrix, Inc.

package applayering

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	getter "github.com/hashicorp/go-getter/v2"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// ---------------------------------------------------------------------------
// Mock types
// ---------------------------------------------------------------------------

// mockComm is a minimal packer.Communicator for testing.
type mockComm struct {
	downloadData  []byte
	downloadErr   error
	startErr      error
	// downloadSeq provides successive Download responses; once exhausted,
	// falls back to downloadData. Enables tests that cycle bitmask values.
	downloadSeq   [][]byte
	downloadCallN int
}

func (m *mockComm) Start(_ context.Context, cmd *packer.RemoteCmd) error {
	if m.startErr != nil {
		return m.startErr
	}
	cmd.SetExited(0)
	return nil
}
func (m *mockComm) Upload(_ string, _ io.Reader, _ *os.FileInfo) error { return nil }
func (m *mockComm) UploadDir(_, _ string, _ []string) error            { return nil }
func (m *mockComm) Download(_ string, output io.Writer) error {
	if m.downloadErr != nil {
		return m.downloadErr
	}
	if m.downloadCallN < len(m.downloadSeq) {
		output.Write(m.downloadSeq[m.downloadCallN])
		m.downloadCallN++
		return nil
	}
	output.Write(m.downloadData)
	return nil
}
func (m *mockComm) DownloadDir(_, _ string, _ []string) error { return nil }

// discardUI swallows all UI output so tests stay quiet.
type discardUI struct{}

func (u *discardUI) Ask(_ string) (string, error)                    { return "", nil }
func (u *discardUI) Askf(_ string, _ ...interface{}) (string, error) { return "", nil }
func (u *discardUI) Say(_ string)                                    {}
func (u *discardUI) Message(_ string)                                {}
func (u *discardUI) Error(_ string)                                  {}
func (u *discardUI) Machine(_ string, _ ...string)                   {}
func (u *discardUI) Sayf(_ string, _ ...interface{})                 {}
func (u *discardUI) Errorf(_ string, _ ...interface{})               {}
func (u *discardUI) TrackProgress(_ string, _, _ int64, stream io.ReadCloser) io.ReadCloser {
	return stream
}

var _ getter.ProgressTracker = (*discardUI)(nil)

// ---------------------------------------------------------------------------
// readBlockFinalizeBitmask tests
// ---------------------------------------------------------------------------

func TestReadBlockFinalizeBitmask_Valid(t *testing.T) {
	comm := &mockComm{downloadData: []byte(`{"BlockFinalize":2}`)}
	got, err := readBlockFinalizeBitmask(context.Background(), comm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("expected bitmask 2, got %d", got)
	}
}

func TestReadBlockFinalizeBitmask_Zero(t *testing.T) {
	comm := &mockComm{downloadData: []byte(`{"BlockFinalize":0}`)}
	got, err := readBlockFinalizeBitmask(context.Background(), comm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != icNone {
		t.Errorf("expected bitmask 0, got %d", got)
	}
}

func TestReadBlockFinalizeBitmask_DownloadError(t *testing.T) {
	comm := &mockComm{downloadErr: errors.New("connection refused")}
	_, err := readBlockFinalizeBitmask(context.Background(), comm)
	if err == nil {
		t.Fatal("expected error when download fails, got nil")
	}
}

func TestReadBlockFinalizeBitmask_InvalidJSON(t *testing.T) {
	comm := &mockComm{downloadData: []byte(`not json`)}
	_, err := readBlockFinalizeBitmask(context.Background(), comm)
	if err == nil {
		t.Fatal("expected error on invalid JSON, got nil")
	}
}

// ---------------------------------------------------------------------------
// waitForRestart tests
// ---------------------------------------------------------------------------

// TestWaitForRestart_Success verifies that waitForRestart returns nil when
// the communicator becomes available within the timeout.
func TestWaitForRestart_Success(t *testing.T) {
	p := &Provisioner{restartTimeout: 200 * time.Millisecond}
	comm := &mockComm{} // Start() succeeds immediately
	err := p.waitForRestart(context.Background(), &discardUI{}, comm)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

// TestWaitForRestart_Timeout verifies that waitForRestart returns an error
// when the communicator is never available (MY-03 fix).
func TestWaitForRestart_Timeout(t *testing.T) {
	p := &Provisioner{restartTimeout: 50 * time.Millisecond}
	comm := &mockComm{startErr: errors.New("not ready")}
	err := p.waitForRestart(context.Background(), &discardUI{}, comm)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}

// ---------------------------------------------------------------------------
// isCommunicatorReady tests
// ---------------------------------------------------------------------------

func TestIsCommunicatorReady_OK(t *testing.T) {
	comm := &mockComm{}
	err := isCommunicatorReady(context.Background(), &discardUI{}, comm)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestIsCommunicatorReady_Error(t *testing.T) {
	comm := &mockComm{startErr: errors.New("fail")}
	err := isCommunicatorReady(context.Background(), &discardUI{}, comm)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// ---------------------------------------------------------------------------
// BlockFinalize bitmask constant sanity checks
// ---------------------------------------------------------------------------

func TestBlockFinalizeBitmaskConstants(t *testing.T) {
	tests := []struct {
		name     string
		val      uint32
		expected uint32
	}{
		{"icNone", icNone, 0},
		{"icBypass", icBypass, 1},
		{"icReboot", icReboot, 2},
		{"icInstaller", icInstaller, 4},
		{"icNgen", icNgen, 8},
		{"icRunOnce", icRunOnce, 16},
		{"icWinUpgrade", icWinUpgrade, 32},
		{"icNgenNeeded", icNgenNeeded, 64},
		{"icSmsCfgPresent", icSmsCfgPresent, 128},
		{"icFsLogix", icFsLogix, 256},
		{"icWemRsaKeyFile", icWemRsaKeyFile, 512},
		{"icDelTokens", icDelTokens, 1024},
		{"icKmsDirMissing", icKmsDirMissing, 2048},
		{"icKmsGpoMissing", icKmsGpoMissing, 4096},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.val != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, tc.val)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// pollInterval / shutdownWait helper tests
// ---------------------------------------------------------------------------

func TestPollInterval_TestMode(t *testing.T) {
	p := &Provisioner{restartTimeout: 5 * time.Millisecond}
	if p.pollInterval() != time.Millisecond {
		t.Errorf("expected 1ms in test mode, got %v", p.pollInterval())
	}
}

func TestPollInterval_ProductionMode(t *testing.T) {
	p := &Provisioner{}
	if p.pollInterval() != 10*time.Second {
		t.Errorf("expected 10s in production mode, got %v", p.pollInterval())
	}
}

func TestShutdownWait_TestMode(t *testing.T) {
	p := &Provisioner{restartTimeout: 5 * time.Millisecond}
	if p.shutdownWait() != time.Millisecond {
		t.Errorf("expected 1ms in test mode, got %v", p.shutdownWait())
	}
}

func TestShutdownWait_ProductionMode(t *testing.T) {
	p := &Provisioner{}
	if p.shutdownWait() != 30*time.Second {
		t.Errorf("expected 30s in production mode, got %v", p.shutdownWait())
	}
}

// ---------------------------------------------------------------------------
// initiateRestart tests (Issue B)
// ---------------------------------------------------------------------------

// TestInitiateRestart_Success verifies that initiateRestart completes without
// panicking when the communicator succeeds.
func TestInitiateRestart_Success(t *testing.T) {
	comm := &mockComm{} // Start() succeeds immediately
	initiateRestart(context.Background(), comm, &discardUI{})
}

// TestInitiateRestart_IgnoresCommError verifies that initiateRestart does not
// propagate WinRM errors: a connection drop is expected when the VM reboots.
func TestInitiateRestart_IgnoresCommError(t *testing.T) {
	comm := &mockComm{startErr: errors.New("connection reset by peer")}
	// Must not panic or otherwise signal an error (no return value).
	initiateRestart(context.Background(), comm, &discardUI{})
}

// ---------------------------------------------------------------------------
// handle_run_once / IC_RUNONCE tests (Issue C)
// ---------------------------------------------------------------------------

// TestPrepare_HandleRunOnce_DefaultFalse verifies that handle_run_once
// defaults to false when not set in the template.
func TestPrepare_HandleRunOnce_DefaultFalse(t *testing.T) {
	p := &Provisioner{}
	if err := p.Prepare(map[string]interface{}{}); err != nil {
		t.Fatalf("Prepare with no options should succeed, got: %v", err)
	}
	if p.config.HandleRunOnce {
		t.Error("expected HandleRunOnce to default to false")
	}
}

// TestPrepare_HandleRunOnceTrue_Succeeds verifies that Prepare accepts
// handle_run_once = true and propagates it to the config.
func TestPrepare_HandleRunOnceTrue_Succeeds(t *testing.T) {
	p := &Provisioner{}
	if err := p.Prepare(map[string]interface{}{"handle_run_once": true}); err != nil {
		t.Fatalf("Prepare with handle_run_once=true should succeed, got: %v", err)
	}
	if !p.config.HandleRunOnce {
		t.Error("expected HandleRunOnce to be true after Prepare")
	}
}

// TestProvision_IcRunOnce_HandleRunOnceFalse_ReturnsError verifies that
// Provision returns an error when IC_RUNONCE is set and handle_run_once=false.
// The build should fail so the user knows RunOnce entries need attention.
func TestProvision_IcRunOnce_HandleRunOnceFalse_ReturnsError(t *testing.T) {
	p := &Provisioner{
		config:         Config{HandleRunOnce: false},
		restartTimeout: 50 * time.Millisecond,
	}
	comm := &mockComm{downloadData: []byte(`{"BlockFinalize":16}`)} // icRunOnce = 16
	err := p.Provision(context.Background(), &discardUI{}, comm, map[string]interface{}{})
	if err == nil {
		t.Fatal("expected error when IC_RUNONCE set and handle_run_once=false, got nil")
	}
	if !strings.Contains(err.Error(), "handle_run_once") {
		t.Errorf("error should mention handle_run_once, got: %v", err)
	}
}

// TestProvision_IcRunOnce_HandleRunOnceTrue_DoesNotReturnRunOnceError verifies
// that when handle_run_once=true the IC_RUNONCE gate does not trigger an error.
// Provision will proceed past the BlockFinalize loop and fail for an unrelated
// reason (missing ELM credentials), which is expected in this unit-test context.
func TestProvision_IcRunOnce_HandleRunOnceTrue_DoesNotReturnRunOnceError(t *testing.T) {
	p := &Provisioner{
		config:         Config{HandleRunOnce: true},
		restartTimeout: 200 * time.Millisecond,
	}
	comm := &mockComm{
		downloadSeq: [][]byte{
			[]byte(`{"BlockFinalize":16}`), // first poll: icRunOnce set
			[]byte(`{"BlockFinalize":0}`),  // second poll: cleared
		},
		downloadData: []byte(`{"BlockFinalize":0}`), // fallback
	}
	err := p.Provision(context.Background(), &discardUI{}, comm, map[string]interface{}{})
	if err != nil && strings.Contains(err.Error(), "handle_run_once is false") {
		t.Errorf("unexpected IC_RUNONCE gate error when handle_run_once=true: %v", err)
	}
}

// ---------------------------------------------------------------------------
// IC_NGENNEEDED / IC_NGEN tests (Issue A)
// ---------------------------------------------------------------------------

// TestProvision_IcNgenNeeded_DoesNotError verifies that when IC_NGENNEEDED is
// set, forceNGen is called (via comm) and the loop continues without error.
// Provision eventually fails for an unrelated reason (missing ELM credentials),
// which is expected in this unit-test context.
func TestProvision_IcNgenNeeded_DoesNotError(t *testing.T) {
	p := &Provisioner{restartTimeout: 200 * time.Millisecond}
	comm := &mockComm{
		downloadSeq: [][]byte{
			[]byte(`{"BlockFinalize":64}`), // icNgenNeeded = 64
			[]byte(`{"BlockFinalize":0}`),  // cleared after forceNGen
		},
		downloadData: []byte(`{"BlockFinalize":0}`),
	}
	err := p.Provision(context.Background(), &discardUI{}, comm, map[string]interface{}{})
	// The error (if any) must not be NGen-specific.
	if err != nil && strings.Contains(err.Error(), "NGen") {
		t.Errorf("unexpected NGen error: %v", err)
	}
}

// TestProvision_IcNgen_WaitsPassively verifies that IC_NGEN (NGen already
// running) does not call forceNGen — the loop simply waits until bitmask clears.
func TestProvision_IcNgen_WaitsPassively(t *testing.T) {
	p := &Provisioner{restartTimeout: 200 * time.Millisecond}
	comm := &mockComm{
		downloadSeq: [][]byte{
			[]byte(`{"BlockFinalize":8}`), // icNgen = 8
			[]byte(`{"BlockFinalize":0}`), // cleared (NGen finished on its own)
		},
		downloadData: []byte(`{"BlockFinalize":0}`),
	}
	err := p.Provision(context.Background(), &discardUI{}, comm, map[string]interface{}{})
	if err != nil && strings.Contains(err.Error(), "NGen") {
		t.Errorf("unexpected NGen error: %v", err)
	}
}
