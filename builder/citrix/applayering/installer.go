// Copyright (c) Citrix, Inc.

package applayering

import (
	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
)

// NewBuilderStrategy returns the BuilderStrategy that corresponds to the
// single non-nil operation block in cfg.
func NewBuilderStrategy(cfg *alaconfig.Config) BuilderStrategy {
	switch {
	case cfg.CreateApp != nil:
		return &createAppStrategy{cfg: cfg}
	case cfg.CreatePlatform != nil:
		return &createPlatformStrategy{cfg: cfg}
	case cfg.RevisionOs != nil:
		return &revisionOsStrategy{cfg: cfg}
	case cfg.RevisionPlatform != nil:
		return &revisionPlatformStrategy{cfg: cfg}
	case cfg.RevisionApp != nil:
		return &revisionAppStrategy{cfg: cfg}
	default:
		return &connectOnlyStrategy{cfg: cfg}
	}
}
