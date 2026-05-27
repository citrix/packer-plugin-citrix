package common

import (
	"context"
	"fmt"

	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// OperationPreflightStrategy validates preconditions for one operation type.
// Each operation type has its own strategy; adding a new type means adding a new Strategy, not modifying existing code.
type OperationPreflightStrategy interface {
	Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error
}

// newPreflightStrategy returns the preflight strategy for the given operation type.
// Operations with no preconditions return noOpPreflight.
func newPreflightStrategy(opType elmsoap.ApplayeringOperationType) OperationPreflightStrategy {
	switch opType {
	case elmsoap.CREATE_PLATFORM_LAYER:
		return &createPlatformLayerPreflight{}
	case elmsoap.CREATE_APP_LAYER:
		return &createAppLayerPreflight{}
	case elmsoap.REVISION_OS_LAYER:
		return &revisionOsLayerPreflight{}
	case elmsoap.REVISION_PLATFORM_LAYER:
		return &revisionPlatformLayerPreflight{}
	case elmsoap.REVISION_APP_LAYER:
		return &revisionAppLayerPreflight{}
	default:
		return &noOpPreflight{}
	}
}

// noOpPreflight is used for operation types that have no preflight preconditions.
type noOpPreflight struct{}

func (s *noOpPreflight) Validate(_ *elmsoap.SoapHelper, _ *alaconfig.Config) error { return nil }

// createPlatformLayerPreflight validates CREATE_PLATFORM_LAYER preconditions:
// - referenced OS layer version exists
// - target platform layer name does not already exist
type createPlatformLayerPreflight struct{}

func (s *createPlatformLayerPreflight) Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error {
	c := cfg.CreatePlatform
	_, err := helper.GetOsLayerRevisionId(c.OsLayerName, c.OsLayerVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer version: %w", err)
	}
	existingId, _ := helper.GetPlatformLayerId(c.LayerName)
	if existingId != 0 {
		return fmt.Errorf("layer already exists: %s", c.LayerName)
	}
	return nil
}

// createAppLayerPreflight validates CREATE_APP_LAYER preconditions:
// - referenced OS layer version exists
// - referenced platform layer version exists (optional; skipped when empty)
// - target app layer name does not already exist
type createAppLayerPreflight struct{}

func (s *createAppLayerPreflight) Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error {
	c := cfg.CreateApp
	_, err := helper.GetOsLayerRevisionId(c.OsLayerName, c.OsLayerVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer version: %w", err)
	}
	if c.PlatformLayerName != "" { // optional
		_, err := helper.GetPlatformLayerRevisionId(c.PlatformLayerName, c.PlatformLayerVersionName)
		if err != nil {
			return fmt.Errorf("preflight: validate platform layer version: %w", err)
		}
	}
	existingId, _ := helper.GetAppLayerId(c.LayerName)
	if existingId != 0 {
		return fmt.Errorf("layer already exists: %s", c.LayerName)
	}
	return nil
}

// revisionOsLayerPreflight validates REVISION_OS_LAYER preconditions:
// - OS layer exists
// - base version exists
// - target version does not already exist
type revisionOsLayerPreflight struct{}

func (s *revisionOsLayerPreflight) Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error {
	c := cfg.RevisionOs
	_, err := helper.GetOsLayerId(c.LayerName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer: %w", err)
	}
	_, err = helper.GetOsLayerRevisionId(c.LayerName, c.BaseVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer version: %w", err)
	}
	targetId, _ := helper.GetOsLayerRevisionId(c.LayerName, c.VersionName)
	if targetId != 0 {
		return fmt.Errorf("target OS layer version already exists: %s / %s", c.LayerName, c.VersionName)
	}
	return nil
}

// revisionPlatformLayerPreflight validates REVISION_PLATFORM_LAYER preconditions:
// - referenced OS layer version exists
// - platform layer exists
// - base version exists
// - target version does not already exist
type revisionPlatformLayerPreflight struct{}

func (s *revisionPlatformLayerPreflight) Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error {
	c := cfg.RevisionPlatform
	_, err := helper.GetOsLayerRevisionId(c.OsLayerName, c.OsLayerVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer version: %w", err)
	}
	_, err = helper.GetPlatformLayerId(c.LayerName)
	if err != nil {
		return fmt.Errorf("preflight: validate platform layer: %w", err)
	}
	_, err = helper.GetPlatformLayerRevisionId(c.LayerName, c.BaseVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate platform layer version: %w", err)
	}
	targetId, _ := helper.GetPlatformLayerRevisionId(c.LayerName, c.VersionName)
	if targetId != 0 {
		return fmt.Errorf("target platform layer version already exists: %s / %s", c.LayerName, c.VersionName)
	}
	return nil
}

// revisionAppLayerPreflight validates REVISION_APP_LAYER preconditions:
// - referenced OS layer version exists
// - referenced platform layer version exists (optional; skipped when empty)
// - app layer exists
// - base version exists (optional; empty = auto-latest)
// - target version does not already exist
type revisionAppLayerPreflight struct{}

func (s *revisionAppLayerPreflight) Validate(helper *elmsoap.SoapHelper, cfg *alaconfig.Config) error {
	c := cfg.RevisionApp
	_, err := helper.GetOsLayerRevisionId(c.OsLayerName, c.OsLayerVersionName)
	if err != nil {
		return fmt.Errorf("preflight: validate OS layer version: %w", err)
	}
	if c.PlatformLayerName != "" { // optional
		_, err := helper.GetPlatformLayerRevisionId(c.PlatformLayerName, c.PlatformLayerVersionName)
		if err != nil {
			return fmt.Errorf("preflight: validate platform layer version: %w", err)
		}
	}
	_, err = helper.GetAppLayerId(c.LayerName)
	if err != nil {
		return fmt.Errorf("preflight: validate app layer: %w", err)
	}
	// BaseVersionName is optional (empty = auto-latest)
	if c.BaseVersionName != "" {
		_, revErr := helper.GetAppLayerRevisionId(c.LayerName, c.BaseVersionName)
		if revErr != nil {
			return fmt.Errorf("preflight: validate app layer version: %w", revErr)
		}
	}
	targetId, _ := helper.GetAppLayerRevisionId(c.LayerName, c.VersionName)
	if targetId != 0 {
		return fmt.Errorf("target app layer version already exists: %s / %s", c.LayerName, c.VersionName)
	}
	return nil
}

type StepPreflightCheck struct {
	Config *alaconfig.Config
}

// Run validates preconditions for the configured operation type using the Strategy pattern.
func (s *StepPreflightCheck) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)

	var operationType elmsoap.ApplayeringOperationType
	switch {
	case s.Config.CreateApp != nil:
		operationType = elmsoap.CREATE_APP_LAYER
	case s.Config.CreatePlatform != nil:
		operationType = elmsoap.CREATE_PLATFORM_LAYER
	case s.Config.RevisionOs != nil:
		operationType = elmsoap.REVISION_OS_LAYER
	case s.Config.RevisionPlatform != nil:
		operationType = elmsoap.REVISION_PLATFORM_LAYER
	case s.Config.RevisionApp != nil:
		operationType = elmsoap.REVISION_APP_LAYER
	}

	strategy := newPreflightStrategy(operationType)
	if err := strategy.Validate(helper, s.Config); err != nil {
		ui.Errorf("[ERR] %v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepPreflightCheck) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
