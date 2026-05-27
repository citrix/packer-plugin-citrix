// Copyright (c) Citrix, Inc.

package applayering

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// layerIntegrityPath is the JSON file written by the ELM agent in the CE VM.
const layerIntegrityPath = `C:\Program Files\Unidesk\Uniservice\Log\LayerIntegrity.txt`

const shutdownForFinalizeCommand = `"C:\Program Files\Unidesk\Uniservice\ShutdownForFinalize.cmd"`

// rebootInitialWait is the time to wait after sending a restart command
// before polling for the VM to come back online.
const rebootInitialWait = 15 * time.Second

// BlockFinalize bitmask constants (ELM agent specification).
const (
	icNone          uint32 = 0
	icBypass        uint32 = 1
	icReboot        uint32 = 2
	icInstaller     uint32 = 4
	icNgen          uint32 = 8
	icRunOnce       uint32 = 16
	icWinUpgrade    uint32 = 32
	icNgenNeeded    uint32 = 64
	icSmsCfgPresent uint32 = 128
	icFsLogix       uint32 = 256
	icWemRsaKeyFile uint32 = 512
	icDelTokens     uint32 = 1024
	icKmsDirMissing uint32 = 2048
	icKmsGpoMissing uint32 = 4096
)

// layerIntegrity maps the JSON structure written by the ELM agent.
// BlockFinalize may be serialized as a JSON string ("2") or number (2).
type layerIntegrity struct {
	BlockFinalize jsonUint32 `json:"BlockFinalize"`
}

// jsonUint32 unmarshals both JSON number and JSON string into uint32.
type jsonUint32 uint32

func (j *jsonUint32) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	var n uint64
	if err := json.Unmarshal([]byte(s), &n); err != nil {
		return err
	}
	*j = jsonUint32(n)
	return nil
}

// Config holds provisioner configuration.
// The shutdown_for_finalize flag has been removed; the provisioner always
// executes the BlockFinalize loop, then ShutdownForFinalize.cmd.
type Config struct {
	ctx interpolate.Context

	// HandleRunOnce controls whether the provisioner will automatically
	// process HKLM RunOnce registry entries when IC_RUNONCE is set in
	// BlockFinalize. Defaults to false.
	//
	// When false (default): if IC_RUNONCE is encountered, the provisioner
	// returns an error, requiring the user to handle RunOnce entries
	// themselves before invoking this provisioner.
	//
	// When true: RunOnce entries are processed one at a time. Note that many
	// RunOnce entries expect an interactive Windows session and may not behave
	// correctly in the headless CE VM environment.
	HandleRunOnce bool `mapstructure:"handle_run_once"`
}

// Provisioner implements packer.Provisioner.
type Provisioner struct {
	config Config

	// restartTimeout overrides the 30-minute default. Set to a small value
	// (e.g. 50ms) in unit tests so they complete instantly.
	restartTimeout time.Duration
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.provisioner.applayering",
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}
	if p.config.HandleRunOnce {
		log.Printf("[WARN] handle_run_once is enabled. RunOnce entries will be executed " +
			"automatically in the headless CE VM. Ensure your RunOnce entries do not " +
			"require an interactive Windows session.")
	}
	return nil
}

// Provision executes the finalization workflow:
//  1. Poll LayerIntegrity.txt until BlockFinalize == 0, handling each
//     bitmask flag (reboot, RunOnce, NGen, etc.).
//  2. Run ShutdownForFinalize.cmd once all blockers are cleared.
//  3. Monitor the ELM work ticket until it completes.
func (p *Provisioner) Provision(ctx context.Context, ui packer.Ui, comm packer.Communicator, generatedData map[string]interface{}) error {
	timeout := 30 * time.Minute
	if p.restartTimeout != 0 {
		timeout = p.restartTimeout
	}
	deadline := time.Now().Add(timeout)

	uiSay(ui, "[INFO] Waiting for BlockFinalize to reach 0...")
	finalizeReady := false
	for time.Now().Before(deadline) {
		bitmask, err := readBlockFinalizeBitmask(ctx, comm)
		if err != nil {
			log.Printf("[WARN] Failed to read LayerIntegrity.txt: %v; retrying...", err)
			time.Sleep(p.pollInterval())
			continue
		}

		if bitmask == icNone {
			finalizeReady = true
			break
		}

		if bitmask&icReboot != 0 {
			uiSay(ui, "[INFO] BlockFinalize IC_REBOOT set; initiating restart...")
			initiateRestart(ctx, comm, ui)
			if err := p.waitForRestart(ctx, ui, comm); err != nil {
				return fmt.Errorf("waitForRestart failed: %w", err)
			}
			continue
		}
		// IC_RUNONCE: CE VM has no interactive user session, so RunOnce entries
		// are processed one at a time explicitly. After each entry we re-check
		// the bitmask; if icReboot is now set the reboot branch above handles it.
		if bitmask&icRunOnce != 0 {
			if !p.config.HandleRunOnce {
				uiSay(ui, "[ERROR] BlockFinalize IC_RUNONCE is set but handle_run_once is false. "+
					"RunOnce entries must be processed before the layer can be finalized. "+
					"Either set handle_run_once = true in the provisioner block, or set "+
					"skip_cleanup_on_failure = true in the builder block to keep the VM alive "+
					"for manual RunOnce processing.")
				return fmt.Errorf("IC_RUNONCE is set but handle_run_once is false; " +
					"set handle_run_once = true to process RunOnce entries automatically, " +
					"or set skip_cleanup_on_failure = true to keep the VM for manual processing")
			}
			ran, err := runOneRunOnceEntry(ctx, comm, ui)
			if err != nil {
				log.Printf("[WARN] RunOnce entry error: %v; continuing...", err)
			} else if !ran {
				// No entries in registry yet — the previous entry's side-effects
				// (Reboot / more RunOnce keys) haven't been written; wait one poll cycle.
				log.Printf("[DEBUG] IC_RUNONCE set but no entries found; waiting...")
			}
			// Regardless of whether entries ran, wait one poll cycle before
			// re-checking the bitmask to avoid spinning before the ELM agent updates it.
			time.Sleep(p.pollInterval())
			continue
		}

		if bitmask&icNgenNeeded != 0 {
			forceNGen(ctx, comm, ui)
			time.Sleep(p.pollInterval())
			continue
		}

		if bitmask&icNgen != 0 {
			log.Printf("[INFO] BlockFinalize IC_NGEN set; NGen in progress, waiting...")
		}
		log.Printf("[INFO] BlockFinalize bitmask: %d; waiting...", bitmask)
		time.Sleep(p.pollInterval())
	}

	if !finalizeReady {
		return fmt.Errorf("timed out waiting for BlockFinalize to reach 0 after %v", timeout)
	}

	uiSay(ui, "[INFO] BlockFinalize is 0; running ShutdownForFinalize...")
	shutdownCmd := &packer.RemoteCmd{Command: shutdownForFinalizeCommand}
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	shutdownCmd.Stdout = &stdoutBuf
	shutdownCmd.Stderr = &stderrBuf
	shutdownCmd.RunWithUi(ctx, comm, &logUi{ui})
	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	maxRetries := 6
	retryCount := 0
	for !strings.Contains(stdoutStr, "Press any key to continue") && retryCount < maxRetries {
		log.Printf("[INFO] 'Press any key to continue' not found in Stdout, sleeping and retrying (%d/%d)...", retryCount+1, maxRetries)
		time.Sleep(p.pollInterval())
		stdoutStr = stdoutBuf.String()
		retryCount++
	}
	time.Sleep(p.shutdownWait())
	if strings.Contains(stdoutStr, "The operation completed successfully") {
		uiSay(ui, "[INFO] Shutdown for finalize completed successfully. Proceeding to next step.")
	} else {
		uiSay(ui, "[INFO] Shutdown for finalize failed.")
		ui.Errorf("Shutdown for finalize command Stdout: %s", stdoutStr)
		ui.Errorf("Shutdown for finalize command stderr: %s", stderrStr)
		return fmt.Errorf("shutdown for finalize command failed")
	}
	if stderrStr != "" {
		ui.Errorf("Shutdown for finalize command stderr: %s", stderrStr)
	}

	uiSay(ui, "[INFO] Monitoring ShutdownForFinalize task state...")
	return p.monitorTaskState(ui, generatedData)
}

// pollInterval returns the retry interval. Short in test mode.
func (p *Provisioner) pollInterval() time.Duration {
	if p.restartTimeout != 0 && p.restartTimeout < 10*time.Second {
		return time.Millisecond
	}
	return 10 * time.Second
}

// shutdownWait returns the post-shutdown additional wait duration.
func (p *Provisioner) shutdownWait() time.Duration {
	if p.restartTimeout != 0 && p.restartTimeout < 10*time.Second {
		return time.Millisecond
	}
	return 30 * time.Second
}

// readBlockFinalizeBitmask downloads LayerIntegrity.txt from the CE VM and
// returns the BlockFinalize integer.
func readBlockFinalizeBitmask(_ context.Context, comm packer.Communicator) (uint32, error) {
	var buf bytes.Buffer
	if err := comm.Download(layerIntegrityPath, &buf); err != nil {
		return 0, fmt.Errorf("download LayerIntegrity.txt: %w", err)
	}
	var li layerIntegrity
	if err := json.Unmarshal(buf.Bytes(), &li); err != nil {
		return 0, fmt.Errorf("parse LayerIntegrity.txt: %w", err)
	}
	return uint32(li.BlockFinalize), nil
}

// initiateRestart sends a restart command to the CE VM via the communicator.
// Any error from the command itself is intentionally ignored: once the OS
// begins rebooting, WinRM will drop the connection, which surfaces as an
// error even when the restart was successfully initiated. The caller must
// follow up with waitForRestart to detect failure.
func initiateRestart(ctx context.Context, comm packer.Communicator, ui packer.Ui) {
	uiSay(ui, "[INFO] Sending restart command to CE VM...")
	cmd := &packer.RemoteCmd{Command: "shutdown /r /t 0"}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.RunWithUi(ctx, comm, &logUi{ui}); err != nil {
		log.Printf("[DEBUG] Restart command returned error (expected if VM is rebooting): %v", err)
	}
}

// waitForRestart waits for the VM communicator to become responsive again.
// Returns an error if the deadline is exceeded (MY-03 fix).
func (p *Provisioner) waitForRestart(ctx context.Context, ui packer.Ui, comm packer.Communicator) error {
	timeout := 30 * time.Minute
	if p.restartTimeout != 0 {
		timeout = p.restartTimeout
	}
	// Skip the initial wait in test mode (short timeout).
	if timeout >= 10*time.Second {
		log.Printf("[INFO] Waiting %v for reboot to start...", rebootInitialWait)
		time.Sleep(rebootInitialWait)
	}
	log.Printf("[INFO] Attempting to reconnect (timeout: %v)...", timeout)
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if err := isCommunicatorReady(ctx, ui, comm); err == nil {
			uiSay(ui, "[INFO] Reconnected to the guest OS.")
			return nil
		}
		log.Printf("[DEBUG] Still waiting for communicator...")
		time.Sleep(p.pollInterval())
	}
	return fmt.Errorf("timed out waiting for VM restart after %v", timeout)
}

// isCommunicatorReady sends a test command to verify the communicator is responding.
func isCommunicatorReady(ctx context.Context, ui packer.Ui, comm packer.Communicator) error {
	cmd := &packer.RemoteCmd{Command: "echo READY"}
	if err := cmd.RunWithUi(ctx, comm, &logUi{ui}); err != nil {
		return err
	}
	return nil
}

func (p *Provisioner) monitorTaskState(ui packer.Ui, generatedData map[string]any) error {
	log.Print("[INFO] Monitor the state of the task.")
	value, ok := generatedData["ELM_SERVER"]
	if !ok {
		return fmt.Errorf("get elm server from generated data failed")
	}
	elm_server := value.(string)
	log.Printf("From generatedData, ELM_SERVER:%s", elm_server)
	value, ok = generatedData["USER_NAME"]
	if !ok {
		return fmt.Errorf("get user name from generated data failed")
	}
	username := value.(string)
	log.Printf("From generatedData, USER_NAME:%s", username)
	value, ok = generatedData["PASSWORD"]
	if !ok {
		return fmt.Errorf("get password data failed")
	}
	password := value.(string)

	unideskurl, err := elmsoap.BuildServerURL(elm_server)
	if err != nil {
		return err
	}
	log.Printf("From generatedData, unideskurl:%s", unideskurl)

	insecureSkipVerify := false
	value, ok = generatedData["INSECURE_CONNECTION"]
	if !ok {
		uiSay(ui, "Get insecure connection data failed, using default value false")
		value = false
	} else {
		insecureSkipVerify = value.(bool)
		log.Printf("From generatedData, INSECURE_CONNECTION:%t", insecureSkipVerify)
	}

	timeout := 30 * time.Minute
	deadline := time.Now().Add(timeout)

	// MY-04: reuse ELM session from the builder when available.
	// Use token-based auth if ELM_TOKEN is present, regardless of whether
	// ELM_COOKIE is set — some servers authenticate via token alone.
	// Only fall back to Login2 if no token was passed from the builder.
	var cookie, token string
	if t, tOk := generatedData["ELM_TOKEN"]; tOk {
		token = t.(string)
		if c, cOk := generatedData["ELM_COOKIE"]; cOk {
			cookie = c.(string)
		}
		log.Printf("[INFO] Reusing ELM session from builder (MY-04)")
	}
	if token == "" {
		cookie, token, err = elmsoap.Login2(username, password, unideskurl, insecureSkipVerify)
		if err != nil {
			ui.Errorf("Login failed due to %v", err)
			return err
		}
	}

	// Create SoapHelper for task monitoring (no SOAP client needed for HTTP-based polling)
	helper := &elmsoap.SoapHelper{
		Cookie:             cookie,
		Token:              token,
		URL:                unideskurl,
		InsecureSkipVerify: insecureSkipVerify,
	}

	value, ok = generatedData["WORK_TICKET_ID"]
	if !ok {
		return fmt.Errorf("get work ticket id from generated data failed")
	}
	workTicketId := value.(int64)
	log.Printf("From generatedData, workTicketId:%d", workTicketId)
	interval := p.pollInterval()
	state := ""
	monitorCnt := 1
	for time.Now().Before(deadline) {
		// MY-05/MY-09: handle GetTaskStateActiveFilter errors with sentinel distinction
		stateTmp, err := helper.GetTaskStateActiveFilter(workTicketId)
		if errors.Is(err, elmsoap.ErrWorkTicketNotInActiveFilter) {
			log.Printf("[INFO] ticket id:%d, tried times: %d, not found in active ticket group, checking completed group.", workTicketId, monitorCnt)
			break
		}
		if err != nil {
			log.Printf("[WARN] GetTaskStateActiveFilter failed (try %d): %v; retrying...", monitorCnt, err)
			time.Sleep(interval)
			monitorCnt++
			continue
		}
		if stateTmp != state {
			state = stateTmp
			log.Printf("Monitor ticket id: %d, tried times: %d, state: %s", workTicketId, monitorCnt, state)
		}
		time.Sleep(interval)
		monitorCnt++
	}
	workTicketResult, err := helper.GetTaskCompletedFilter(workTicketId)
	if err != nil {
		ui.Errorf("GetTaskCompletedFilter failed due to %v", err)
		return err
	}
	state = string(*workTicketResult.State)
	log.Printf("[INFO] Monitor ticket id: %d, tried times: %d, state: %s, workTicketResult: %v. Stop monitor.", workTicketId, monitorCnt, state, workTicketResult)
	uiSayf(ui, "[INFO] Monitor ticket id: %d, state: %s. Stop monitor.", workTicketId, state)
	return nil
}

// uiSay prints a timestamped message to the Packer console (==>).
func uiSay(ui packer.Ui, msg string) {
	ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + msg)
}

// uiSayf prints a timestamped formatted message to the Packer console (==>).
func uiSayf(ui packer.Ui, format string, args ...interface{}) {
	ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + fmt.Sprintf(format, args...))
}

// logUi is a packer.Ui that redirects Say/Sayf to log.Printf instead of the
// console. Pass it to RunWithUi to suppress raw script stdout from packer's
// ==> output while still capturing output in a buffer via MultiWriter.
type logUi struct {
	packer.Ui
}

func (l *logUi) Say(msg string)                    { log.Printf("[TRACE] %s", msg) }
func (l *logUi) Sayf(f string, a ...interface{}) { log.Printf("[TRACE] "+f, a...) }

// runNGenScript invokes ngen.exe executequeueditems for both 32-bit and 64-bit
// .NET frameworks. Paths that do not exist are silently skipped. Written as a
// semicolon-separated one-liner so it can be passed inside
// powershell.exe -Command "..." over WinRM without quoting issues.
const runNGenScript = `` +
	`$ngens = @(` +
	`'C:\Windows\Microsoft.NET\Framework64\v4.0.30319\ngen.exe',` +
	`'C:\Windows\Microsoft.NET\Framework\v4.0.30319\ngen.exe'` +
	`); ` +
	`foreach ($ngen in $ngens) { ` +
	`if (Test-Path $ngen) { ` +
	`Write-Host ('NGen: running ' + $ngen); ` +
	`& $ngen executequeueditems; ` +
	`Write-Host ('NGen: exit ' + $LASTEXITCODE) ` +
	`} else { ` +
	`Write-Host ('NGen: ' + $ngen + ' not found, skipping') ` +
	`} }`

// forceNGen invokes ngen.exe executequeueditems on the CE VM for both 32-bit
// and 64-bit .NET frameworks. Non-zero exit codes and script errors are logged
// as warnings and are not treated as fatal errors.
func forceNGen(ctx context.Context, comm packer.Communicator, ui packer.Ui) {
	uiSay(ui, "[INFO] BlockFinalize IC_NGENNEEDED set; forcing .NET NGen compilation...")
	cmd := &packer.RemoteCmd{Command: `powershell.exe -NonInteractive -Command "` + runNGenScript + `"`}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.RunWithUi(ctx, comm, &logUi{ui}); err != nil {
		log.Printf("[WARN] NGen script returned error: %v (stderr: %s)", err, stderr.String())
		return
	}
	if s := stderr.String(); s != "" {
		log.Printf("[WARN] NGen script stderr: %s", s)
	}
	log.Printf("[TRACE] NGen stdout:\n%s", stdout.String())
}

// runOneRunOnceScript processes a single HKLM RunOnce entry:
//  1. Read the first entry (name + command) and save it.
//  2. Delete that entry from the registry.
//  3. Execute the command.
//
// No reboot is triggered; the caller re-checks the bitmask and handles
// icReboot naturally in the outer BlockFinalize loop.
//
// Written as a semicolon-separated one-liner (no double quotes) so it can be
// passed directly inside  powershell.exe -Command "..."  over WinRM without
// quoting or newline issues.
const runOneRunOnceScript = `` +
	`$key = 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce'; ` +
	`if (-not (Test-Path $key)) { Write-Host 'RunOnce: no key found'; exit 0 }; ` +
	`$props = Get-ItemProperty -Path $key -ErrorAction SilentlyContinue; ` +
	`if (-not $props) { Write-Host 'RunOnce: no properties'; exit 0 }; ` +
	`$entry = $props.PSObject.Properties | Where-Object { $_.Name -notlike 'PS*' } | Select-Object -First 1; ` +
	`if (-not $entry) { Write-Host 'RunOnce: no entries'; exit 0 }; ` +
	`$name = $entry.Name; $cmd = $entry.Value; ` +
	`Write-Host ('RunOnce: saved entry ' + $name + ' = ' + $cmd); ` +
	`Remove-ItemProperty -Path $key -Name $name -Force -ErrorAction SilentlyContinue; ` +
	`Write-Host ('RunOnce: deleted entry ' + $name); ` +
	`Write-Host ('RunOnce: executing ' + $name); ` +
	`cmd /c $cmd; ` +
	`Write-Host ('RunOnce: completed ' + $name)`

// runOneRunOnceEntry reads the first HKLM RunOnce entry, deletes it, then
// executes it. Returns (true, nil) if an entry was found and executed,
// (false, nil) if the registry had no entries (CE VM still clearing the bit),
// or (false, err) on script failure.
func runOneRunOnceEntry(ctx context.Context, comm packer.Communicator, ui packer.Ui) (bool, error) {
	cmd := &packer.RemoteCmd{Command: `powershell.exe -NonInteractive -Command "` + runOneRunOnceScript + `"`}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.RunWithUi(ctx, comm, &logUi{ui}); err != nil {
		return false, fmt.Errorf("RunOnce entry script: %w (stderr: %s)", err, stderr.String())
	}
	if s := stderr.String(); s != "" {
		log.Printf("[WARN] RunOnce entry stderr: %s", s)
	}
	out := stdout.String()
	// The script outputs "RunOnce: no ..." when no entries are found.
	if strings.Contains(out, "RunOnce: no ") {
		return false, nil
	}
	log.Printf("[TRACE] RunOnce stdout:\n%s", out)
	// Show only the first line ("RunOnce: saved entry NAME = ...") on console.
	firstLine := strings.SplitN(strings.TrimSpace(out), "\n", 2)[0]
	uiSay(ui, "[INFO] "+strings.TrimSpace(firstLine))
	return true, nil
}
