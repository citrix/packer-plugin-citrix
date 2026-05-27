// Copyright (c) Citrix, Inc.

package applayering

import (
	"github.com/hashicorp/packer-plugin-sdk/multistep"

	common "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/common"
	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// BuilderStrategy is implemented by each supported operation type.
// Separating operation-specific step assembly into strategies keeps
// Builder.Run stable as new operations are added.
type BuilderStrategy interface {
	// OperationSteps returns the multistep steps that perform the layer operation.
	OperationSteps() []multistep.Step
	// OperationType returns the ELM operation type identifier.
	OperationType() elmsoap.ApplayeringOperationType
	// Artifact constructs the packer Artifact for this operation (MY-11).
	// state is used to carry generated_data through to post-processors.
	// Pass new(multistep.BasicStateBag) when only layer/version name is needed.
	Artifact(state multistep.StateBag) *Artifact
}

// --- createAppStrategy ---

type createAppStrategy struct{ cfg *alaconfig.Config }

func (s *createAppStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.CreateApp
	return []multistep.Step{
		&common.StepCreateAppLayer{Config: &common.CreateAppLayerConfig{
			OsLayerName:                 c.OsLayerName,
			OsLayerVersionName:          c.OsLayerVersionName,
			PlatformLayerName:           c.PlatformLayerName,
			PlatformLayerVersionName:    c.PlatformLayerVersionName,
			PrerequisiteLayers:          c.PrerequisiteLayers,
			PackagingDiskFileName:       c.PackagingDiskFileName,
			PlatformConnectorConfigName: c.PlatformConnectorConfigName,
			LayerName:                   c.LayerName,
			IconId:                      c.IconId,
			VersionName:                 c.VersionName,
			VersionDescription:          c.VersionDescription,
			VersionSizeGb:               c.VersionSizeGb,
			Comment:                     c.Comment,
			SkipCleanupOnFailure:        c.SkipCleanupOnFailure,
		}},
	}
}

func (s *createAppStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.CREATE_APP_LAYER
}

func (s *createAppStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          string(elmsoap.CREATE_APP_LAYER),
		TargetLayerName:        s.cfg.CreateApp.LayerName,
		TargetLayerVersionName: s.cfg.CreateApp.VersionName,
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}

// --- createPlatformStrategy ---

type createPlatformStrategy struct{ cfg *alaconfig.Config }

func (s *createPlatformStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.CreatePlatform
	return []multistep.Step{
		&common.StepCreatePlatformLayering{Config: &common.CreatePlatformLayerConfig{
			HypervisorPlatform:          c.HypervisorPlatform,
			ProvisioningPlatform:        c.ProvisioningPlatform,
			BrokerPlatform:              c.BrokerPlatform,
			OsLayerName:                 c.OsLayerName,
			OsLayerVersionName:          c.OsLayerVersionName,
			PackagingDiskFileName:       c.PackagingDiskFileName,
			PlatformConnectorConfigName: c.PlatformConnectorConfigName,
			LayerName:                   c.LayerName,
			IconId:                      c.IconId,
			VersionName:                 c.VersionName,
			VersionDescription:          c.VersionDescription,
			VersionSizeGb:               c.VersionSizeGb,
			Comment:                     c.Comment,
			SkipCleanupOnFailure:        c.SkipCleanupOnFailure,
		}},
	}
}

func (s *createPlatformStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.CREATE_PLATFORM_LAYER
}

func (s *createPlatformStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          string(elmsoap.CREATE_PLATFORM_LAYER),
		TargetLayerName:        s.cfg.CreatePlatform.LayerName,
		TargetLayerVersionName: s.cfg.CreatePlatform.VersionName,
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}

// --- revisionOsStrategy ---

type revisionOsStrategy struct{ cfg *alaconfig.Config }

func (s *revisionOsStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.RevisionOs
	return []multistep.Step{
		&common.StepRevisionOsLayering{Config: &common.RevisionOsLayerConfig{
			LayerName:                   c.LayerName,
			BaseVersionName:             c.BaseVersionName,
			PackagingDiskFileName:       c.PackagingDiskFileName,
			PlatformConnectorConfigName: c.PlatformConnectorConfigName,
			VersionName:                 c.VersionName,
			VersionDescription:          c.VersionDescription,
			VersionSizeGb:               c.VersionSizeGb,
			Comment:                     c.Comment,
			SkipCleanupOnFailure:        c.SkipCleanupOnFailure,
		}},
	}
}

func (s *revisionOsStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.REVISION_OS_LAYER
}

func (s *revisionOsStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          string(elmsoap.REVISION_OS_LAYER),
		TargetLayerName:        s.cfg.RevisionOs.LayerName,
		TargetLayerVersionName: s.cfg.RevisionOs.VersionName,
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}

// --- revisionPlatformStrategy ---

type revisionPlatformStrategy struct{ cfg *alaconfig.Config }

func (s *revisionPlatformStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.RevisionPlatform
	return []multistep.Step{
		&common.StepRevisionPlatformLayering{Config: &common.RevisionPlatformLayerConfig{
			HypervisorPlatform:          c.HypervisorPlatform,
			ProvisioningPlatform:        c.ProvisioningPlatform,
			BrokerPlatform:              c.BrokerPlatform,
			OsLayerName:                 c.OsLayerName,
			OsLayerVersionName:          c.OsLayerVersionName,
			LayerName:                   c.LayerName,
			BaseVersionName:             c.BaseVersionName,
			PackagingDiskFileName:       c.PackagingDiskFileName,
			PlatformConnectorConfigName: c.PlatformConnectorConfigName,
			VersionName:                 c.VersionName,
			VersionDescription:          c.VersionDescription,
			VersionSizeGb:               c.VersionSizeGb,
			Comment:                     c.Comment,
			SkipCleanupOnFailure:        c.SkipCleanupOnFailure,
		}},
	}
}

func (s *revisionPlatformStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.REVISION_PLATFORM_LAYER
}

func (s *revisionPlatformStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          string(elmsoap.REVISION_PLATFORM_LAYER),
		TargetLayerName:        s.cfg.RevisionPlatform.LayerName,
		TargetLayerVersionName: s.cfg.RevisionPlatform.VersionName,
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}

// --- revisionAppStrategy ---

type revisionAppStrategy struct{ cfg *alaconfig.Config }

func (s *revisionAppStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.RevisionApp
	return []multistep.Step{
		&common.StepRevisionAppLayering{Config: &common.RevisionAppLayerConfig{
			OsLayerName:                 c.OsLayerName,
			OsLayerVersionName:          c.OsLayerVersionName,
			PlatformLayerName:           c.PlatformLayerName,
			PlatformLayerVersionName:    c.PlatformLayerVersionName,
			LayerName:                   c.LayerName,
			BaseVersionName:             c.BaseVersionName,
			PrerequisiteLayers:          c.PrerequisiteLayers,
			PackagingDiskFileName:       c.PackagingDiskFileName,
			PlatformConnectorConfigName: c.PlatformConnectorConfigName,
			VersionName:                 c.VersionName,
			VersionDescription:          c.VersionDescription,
			VersionSizeGb:               c.VersionSizeGb,
			Comment:                     c.Comment,
			SkipCleanupOnFailure:        c.SkipCleanupOnFailure,
		}},
	}
}

func (s *revisionAppStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.REVISION_APP_LAYER
}

func (s *revisionAppStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          string(elmsoap.REVISION_APP_LAYER),
		TargetLayerName:        s.cfg.RevisionApp.LayerName,
		TargetLayerVersionName: s.cfg.RevisionApp.VersionName,
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}

// --- connectOnlyStrategy ---

type connectOnlyStrategy struct{ cfg *alaconfig.Config }

func (s *connectOnlyStrategy) OperationSteps() []multistep.Step {
	c := s.cfg.ConnectOnly
	return []multistep.Step{
		&common.StepGetWorkTicketId{Config: &common.GetWorkTicketIdConfig{
			OperationType: c.OperationType,
			LayerName:     c.LayerName,
		}},
	}
}

func (s *connectOnlyStrategy) OperationType() elmsoap.ApplayeringOperationType {
	return elmsoap.ApplayeringOperationType(s.cfg.ConnectOnly.OperationType)
}

func (s *connectOnlyStrategy) Artifact(state multistep.StateBag) *Artifact {
	return &Artifact{
		OperationType:          s.cfg.ConnectOnly.OperationType,
		TargetLayerName:        s.cfg.ConnectOnly.LayerName,
		TargetLayerVersionName: "",
		StateData:              map[string]any{"generated_data": state.Get("generated_data")},
	}
}
