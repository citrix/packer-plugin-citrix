// Copyright (c) Citrix, Inc.

package applayering

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	packerconfig "github.com/hashicorp/packer-plugin-sdk/template/config"
	common "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/common"
	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
)

const BuilderId = "applayering.builder"

// Builder is the Packer builder for Citrix App Layering operations.
type Builder struct {
	config          alaconfig.Config
	StrategyFactory func(*alaconfig.Config) BuilderStrategy
	runner          multistep.Runner
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

func (b *Builder) Prepare(raws ...any) (_ []string, warnings []string, err error) {
	err = packerconfig.Decode(&b.config,
		&packerconfig.DecodeOpts{
			PluginType:         "packer.builder.applayering",
			Interpolate:        true,
			InterpolateContext: &b.config.Ctx,
		},
		raws...)
	if err != nil {
		return nil, nil, err
	}

	// Validate struct tags (required, required_without_all, etc.)
	if validErr := common.Validate(&b.config); validErr != nil {
		return nil, nil, validErr
	}

	// Enforce exactly one operation block is set.
	count := 0
	if b.config.CreateApp != nil {
		count++
	}
	if b.config.CreatePlatform != nil {
		count++
	}
	if b.config.RevisionOs != nil {
		count++
	}
	if b.config.RevisionPlatform != nil {
		count++
	}
	if b.config.RevisionApp != nil {
		count++
	}
	if b.config.ConnectOnly != nil {
		count++
	}
	if count > 1 {
		return nil, nil, fmt.Errorf("ConfigError: exactly one operation block must be set; found %d", count)
	}

	if !b.config.InsecureConnection {
		warnings = append(warnings, "insecure_connection is false; ignore this only for testing with self-signed certificates")
	}

	// Apply IconId defaults for create operations.
	if warn := applyIconIdDefault(iconIdPtrFromCreateApp(b.config.CreateApp), "create_app_layer"); warn != "" {
		warnings = append(warnings, warn)
	}
	if warn := applyIconIdDefault(iconIdPtrFromCreatePlatform(b.config.CreatePlatform), "create_platform_layer"); warn != "" {
		warnings = append(warnings, warn)
	}

	if b.config.Comm.Type == "" {
		b.config.Comm.Type = "winrm"
		warnings = append(warnings, "communicator not specified, defaulting to 'winrm'")
	}
	if b.config.Comm.Type != "winrm" {
		return nil, warnings, fmt.Errorf("ConfigError: communicator only supports 'winrm'")
	}
	if commErrs := b.config.Comm.Prepare(&b.config.Ctx); len(commErrs) > 0 {
		return nil, warnings, commErrs[0]
	}

	return nil, warnings, nil
}

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	common.UiSay(ui, "Running builder ...")
	log.Print(":: Configuration")
	common.DumpConfig(&b.config, func(s string) { log.Print(s) })

	strategy := b.StrategyFactory(&b.config)
	return b.runStrategy(ctx, ui, hook, strategy)
}

// runStrategy executes the full step chain for the given strategy.
func (b *Builder) runStrategy(ctx context.Context, ui packer.Ui, hook packer.Hook, strategy BuilderStrategy) (packer.Artifact, error) {
	opType := strategy.OperationType()
	common.UiSayf(ui, "[INFO] Operation Type: %s", opType)

	steps := []multistep.Step{
		&common.StepConnect{
			Config: &common.ConnectConfig{
				ELMServer:          b.config.ELMServer,
				ELMUsername:        b.config.ELMUsername,
				ELMPassword:        b.config.ELMPassword,
				InsecureConnection: b.config.InsecureConnection,
			},
		},
		&common.StepPreflightCheck{Config: &b.config},
	}
	steps = append(steps, strategy.OperationSteps()...)

	if b.config.Comm.Type == "winrm" {
		common.UiSay(ui, "Steps add waitforIP step...")
		steps = append(steps, b.waitForIPStep(strategy))
		steps = append(steps,
			&communicator.StepConnect{
				Config:    &b.config.Comm,
				Host:      common.CommHost(b.config.Comm.Host()),
				SSHConfig: b.config.Comm.SSHConfigFunc(),
			},
			&commonsteps.StepProvision{})
	}

	state := new(multistep.BasicStateBag)
	state.Put("hook", hook)
	state.Put("ui", ui)

	generatedData := map[string]interface{}{
		"USER_NAME":           b.config.ELMUsername,
		"PASSWORD":            b.config.ELMPassword,
		"ELM_SERVER":          b.config.ELMServer,
		"INSECURE_CONNECTION": b.config.InsecureConnection,
	}
	state.Put("generated_data", generatedData)

	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)

	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}

	// MY-04: propagate ELM session to provisioner via generatedData so it can
	// reuse the session instead of doing a fresh Login2 call.
	if cookie, ok := state.GetOk("COOKIE"); ok {
		generatedData["ELM_COOKIE"] = cookie
	}
	if token, ok := state.GetOk("UNIDESK_TOKEN"); ok {
		generatedData["ELM_TOKEN"] = token
	}

	return strategy.Artifact(state), nil
}

// waitForIPStep constructs the StepWaitForIp for the given strategy.
func (b *Builder) waitForIPStep(strategy BuilderStrategy) multistep.Step {
	step := common.StepWaitForIp{
		Config: &common.WaitIpConfig{
			OperationType: strategy.OperationType(),
			LayerName:     strategy.Artifact(new(multistep.BasicStateBag)).TargetLayerName,
			WaitTimeout:   b.config.WaitForIpTimeout,
		},
	}
	step.Config.Prepare()
	return &step
}

// iconIdPtrFromCreateApp returns &cfg.IconId, or nil if cfg is nil.
func iconIdPtrFromCreateApp(cfg *alaconfig.CreateAppConfig) *int64 {
	if cfg == nil {
		return nil
	}
	return &cfg.IconId
}

// iconIdPtrFromCreatePlatform returns &cfg.IconId, or nil if cfg is nil.
func iconIdPtrFromCreatePlatform(cfg *alaconfig.CreatePlatformConfig) *int64 {
	if cfg == nil {
		return nil
	}
	return &cfg.IconId
}

// applyIconIdDefault sets *id to defaultIconId when it is zero.
// Returns a warning message if the default was applied.
func applyIconIdDefault(id *int64, opLabel string) string {
	const defaultIconId int64 = 196608
	if id == nil || *id != 0 {
		return ""
	}
	*id = defaultIconId
	return fmt.Sprintf("%s: icon_id not specified, default icon %d will be used", opLabel, defaultIconId)
}
