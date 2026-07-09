// Copyright (c) Citrix, Inc.

package applayering

import (
	"reflect"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	common "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/common"
	alaconfig "github.com/citrix/packer-plugin-citrix/builder/citrix/applayering/config"
	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

// stateWithGenerated returns a populated state bag for Artifact tests.
func stateWithGenerated(value any) multistep.StateBag {
	s := new(multistep.BasicStateBag)
	s.Put("generated_data", value)
	return s
}

// ----------------------------------------------------------------------------
// NewBuilderStrategy dispatch (covers installer.go).
// ----------------------------------------------------------------------------

func TestNewBuilderStrategy_Dispatch(t *testing.T) {
	cases := []struct {
		name string
		cfg  *alaconfig.Config
		want BuilderStrategy
	}{
		{"CreateApp", &alaconfig.Config{CreateApp: &alaconfig.CreateAppConfig{}}, &createAppStrategy{}},
		{"CreatePlatform", &alaconfig.Config{CreatePlatform: &alaconfig.CreatePlatformConfig{}}, &createPlatformStrategy{}},
		{"RevisionOs", &alaconfig.Config{RevisionOs: &alaconfig.RevisionOsConfig{}}, &revisionOsStrategy{}},
		{"RevisionPlatform", &alaconfig.Config{RevisionPlatform: &alaconfig.RevisionPlatformConfig{}}, &revisionPlatformStrategy{}},
		{"RevisionApp", &alaconfig.Config{RevisionApp: &alaconfig.RevisionAppConfig{}}, &revisionAppStrategy{}},
		{"ConnectOnly", &alaconfig.Config{ConnectOnly: &alaconfig.ConnectOnlyConfig{}}, &connectOnlyStrategy{}},
		{"NoBlocks_fallsThroughToConnectOnly", &alaconfig.Config{}, &connectOnlyStrategy{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBuilderStrategy(tc.cfg)
			if reflect.TypeOf(got) != reflect.TypeOf(tc.want) {
				t.Errorf("got %T, want %T", got, tc.want)
			}
		})
	}
}

// ----------------------------------------------------------------------------
// createAppStrategy
// ----------------------------------------------------------------------------

func TestCreateAppStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		CreateApp: &alaconfig.CreateAppConfig{
			OsLayerName:                 "Win10",
			OsLayerVersionName:          "v1",
			PlatformLayerName:           "Plat",
			PlatformLayerVersionName:    "v2",
			PrerequisiteLayers:          []string{"L1:v1"},
			PackagingDiskFileName:       "disk.vhd",
			PlatformConnectorConfigName: "Connector",
			LayerName:                   "AppLayer",
			IconId:                      42,
			VersionName:                 "v3",
			VersionDescription:          "desc",
			VersionSizeGb:               20,
			Comment:                     "reason",
			SkipCleanupOnFailure:        true,
		},
	}
	s := &createAppStrategy{cfg: cfg}

	if got, want := string(s.OperationType()), elmsoap.CREATE_APP_LAYER; got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepCreateAppLayer)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepCreateAppLayer", steps[0])
	}
	wantStepCfg := &common.CreateAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		PlatformLayerName:           "Plat",
		PlatformLayerVersionName:    "v2",
		PrerequisiteLayers:          []string{"L1:v1"},
		PackagingDiskFileName:       "disk.vhd",
		PlatformConnectorConfigName: "Connector",
		LayerName:                   "AppLayer",
		IconId:                      42,
		VersionName:                 "v3",
		VersionDescription:          "desc",
		VersionSizeGb:               20,
		Comment:                     "reason",
		SkipCleanupOnFailure:        true,
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-value"))
	if art.OperationType != string(elmsoap.CREATE_APP_LAYER) {
		t.Errorf("Artifact.OperationType = %v, want %v", art.OperationType, string(elmsoap.CREATE_APP_LAYER))
	}
	if art.TargetLayerName != "AppLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want AppLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "v3" {
		t.Errorf("Artifact.TargetLayerVersionName = %v, want v3", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-value" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-value", got)
	}
}

// ----------------------------------------------------------------------------
// createPlatformStrategy
// ----------------------------------------------------------------------------

func TestCreatePlatformStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		CreatePlatform: &alaconfig.CreatePlatformConfig{
			HypervisorPlatform:          "vmware",
			ProvisioningPlatform:        "MCS",
			BrokerPlatform:              "None",
			OsLayerName:                 "Win10",
			OsLayerVersionName:          "v1",
			PackagingDiskFileName:       "disk.vhd",
			PlatformConnectorConfigName: "Connector",
			LayerName:                   "PlatformLayer",
			IconId:                      456,
			VersionName:                 "RevInfo",
			VersionDescription:          "Revision Description",
			VersionSizeGb:               10,
			Comment:                     "Reason",
			SkipCleanupOnFailure:        false,
		},
	}
	s := &createPlatformStrategy{cfg: cfg}

	if got, want := string(s.OperationType()), elmsoap.CREATE_PLATFORM_LAYER; got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepCreatePlatformLayering)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepCreatePlatformLayering", steps[0])
	}
	wantStepCfg := &common.CreatePlatformLayerConfig{
		HypervisorPlatform:          "vmware",
		ProvisioningPlatform:        "MCS",
		BrokerPlatform:              "None",
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		PackagingDiskFileName:       "disk.vhd",
		PlatformConnectorConfigName: "Connector",
		LayerName:                   "PlatformLayer",
		IconId:                      456,
		VersionName:                 "RevInfo",
		VersionDescription:          "Revision Description",
		VersionSizeGb:               10,
		Comment:                     "Reason",
		SkipCleanupOnFailure:        false,
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-create-plat"))
	if art.OperationType != string(elmsoap.CREATE_PLATFORM_LAYER) {
		t.Errorf("Artifact.OperationType = %v, want %v", art.OperationType, string(elmsoap.CREATE_PLATFORM_LAYER))
	}
	if art.TargetLayerName != "PlatformLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want PlatformLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "RevInfo" {
		t.Errorf("Artifact.TargetLayerVersionName = %v, want RevInfo", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-create-plat" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-create-plat", got)
	}
}

// ----------------------------------------------------------------------------
// revisionOsStrategy
// ----------------------------------------------------------------------------

func TestRevisionOsStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		RevisionOs: &alaconfig.RevisionOsConfig{
			LayerName:                   "OsLayer",
			BaseVersionName:             "base-v1",
			PackagingDiskFileName:       "disk.vhd",
			PlatformConnectorConfigName: "Connector",
			VersionName:                 "rev-v2",
			VersionDescription:          "desc",
			VersionSizeGb:               15,
			Comment:                     "reason",
			SkipCleanupOnFailure:        true,
		},
	}
	s := &revisionOsStrategy{cfg: cfg}

	if got, want := string(s.OperationType()), elmsoap.REVISION_OS_LAYER; got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepRevisionOsLayering)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepRevisionOsLayering", steps[0])
	}
	wantStepCfg := &common.RevisionOsLayerConfig{
		LayerName:                   "OsLayer",
		BaseVersionName:             "base-v1",
		PackagingDiskFileName:       "disk.vhd",
		PlatformConnectorConfigName: "Connector",
		VersionName:                 "rev-v2",
		VersionDescription:          "desc",
		VersionSizeGb:               15,
		Comment:                     "reason",
		SkipCleanupOnFailure:        true,
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-rev-os"))
	if art.OperationType != string(elmsoap.REVISION_OS_LAYER) {
		t.Errorf("Artifact.OperationType = %v, want %v", art.OperationType, string(elmsoap.REVISION_OS_LAYER))
	}
	if art.TargetLayerName != "OsLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want OsLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "rev-v2" {
		t.Errorf("Artifact.TargetLayerVersionName = %v, want rev-v2", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-rev-os" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-rev-os", got)
	}
}

// ----------------------------------------------------------------------------
// revisionPlatformStrategy
// ----------------------------------------------------------------------------

func TestRevisionPlatformStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		RevisionPlatform: &alaconfig.RevisionPlatformConfig{
			HypervisorPlatform:          "vmware",
			ProvisioningPlatform:        "MCS",
			BrokerPlatform:              "None",
			OsLayerName:                 "Win10",
			OsLayerVersionName:          "v1",
			LayerName:                   "PlatLayer",
			BaseVersionName:             "base-v1",
			PackagingDiskFileName:       "disk.vhd",
			PlatformConnectorConfigName: "Connector",
			VersionName:                 "rev-v3",
			VersionDescription:          "desc",
			VersionSizeGb:               12,
			Comment:                     "reason",
			SkipCleanupOnFailure:        false,
		},
	}
	s := &revisionPlatformStrategy{cfg: cfg}

	if got, want := string(s.OperationType()), elmsoap.REVISION_PLATFORM_LAYER; got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepRevisionPlatformLayering)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepRevisionPlatformLayering", steps[0])
	}
	wantStepCfg := &common.RevisionPlatformLayerConfig{
		HypervisorPlatform:          "vmware",
		ProvisioningPlatform:        "MCS",
		BrokerPlatform:              "None",
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		LayerName:                   "PlatLayer",
		BaseVersionName:             "base-v1",
		PackagingDiskFileName:       "disk.vhd",
		PlatformConnectorConfigName: "Connector",
		VersionName:                 "rev-v3",
		VersionDescription:          "desc",
		VersionSizeGb:               12,
		Comment:                     "reason",
		SkipCleanupOnFailure:        false,
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-rev-plat"))
	if art.OperationType != string(elmsoap.REVISION_PLATFORM_LAYER) {
		t.Errorf("Artifact.OperationType = %v, want %v", art.OperationType, string(elmsoap.REVISION_PLATFORM_LAYER))
	}
	if art.TargetLayerName != "PlatLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want PlatLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "rev-v3" {
		t.Errorf("Artifact.TargetLayerVersionName = %v, want rev-v3", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-rev-plat" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-rev-plat", got)
	}
}

// ----------------------------------------------------------------------------
// revisionAppStrategy
// ----------------------------------------------------------------------------

func TestRevisionAppStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		RevisionApp: &alaconfig.RevisionAppConfig{
			OsLayerName:                 "Win10",
			OsLayerVersionName:          "v1",
			PlatformLayerName:           "Plat",
			PlatformLayerVersionName:    "v2",
			LayerName:                   "AppLayer",
			BaseVersionName:             "base-v1",
			PrerequisiteLayers:          []string{"L1:v1", "L2:v2"},
			PackagingDiskFileName:       "disk.vhd",
			PlatformConnectorConfigName: "Connector",
			VersionName:                 "rev-v4",
			VersionDescription:          "desc",
			VersionSizeGb:               25,
			Comment:                     "reason",
			SkipCleanupOnFailure:        true,
		},
	}
	s := &revisionAppStrategy{cfg: cfg}

	if got, want := string(s.OperationType()), elmsoap.REVISION_APP_LAYER; got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepRevisionAppLayering)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepRevisionAppLayering", steps[0])
	}
	wantStepCfg := &common.RevisionAppLayerConfig{
		OsLayerName:                 "Win10",
		OsLayerVersionName:          "v1",
		PlatformLayerName:           "Plat",
		PlatformLayerVersionName:    "v2",
		LayerName:                   "AppLayer",
		BaseVersionName:             "base-v1",
		PrerequisiteLayers:          []string{"L1:v1", "L2:v2"},
		PackagingDiskFileName:       "disk.vhd",
		PlatformConnectorConfigName: "Connector",
		VersionName:                 "rev-v4",
		VersionDescription:          "desc",
		VersionSizeGb:               25,
		Comment:                     "reason",
		SkipCleanupOnFailure:        true,
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-rev-app"))
	if art.OperationType != string(elmsoap.REVISION_APP_LAYER) {
		t.Errorf("Artifact.OperationType = %v, want %v", art.OperationType, string(elmsoap.REVISION_APP_LAYER))
	}
	if art.TargetLayerName != "AppLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want AppLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "rev-v4" {
		t.Errorf("Artifact.TargetLayerVersionName = %v, want rev-v4", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-rev-app" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-rev-app", got)
	}
}

// ----------------------------------------------------------------------------
// connectOnlyStrategy
// ----------------------------------------------------------------------------

func TestConnectOnlyStrategy(t *testing.T) {
	cfg := &alaconfig.Config{
		ConnectOnly: &alaconfig.ConnectOnlyConfig{
			OperationType: "CONNECT_CREATE_APP_VM_ONLY",
			LayerName:     "ActiveAppLayer",
		},
	}
	s := &connectOnlyStrategy{cfg: cfg}

	if got, want := s.OperationType(), elmsoap.ApplayeringOperationType("CONNECT_CREATE_APP_VM_ONLY"); got != want {
		t.Errorf("OperationType = %v, want %v", got, want)
	}

	steps := s.OperationSteps()
	if len(steps) != 1 {
		t.Fatalf("OperationSteps len = %d, want 1", len(steps))
	}
	step, ok := steps[0].(*common.StepGetWorkTicketId)
	if !ok {
		t.Fatalf("step type = %T, want *common.StepGetWorkTicketId", steps[0])
	}
	wantStepCfg := &common.GetWorkTicketIdConfig{
		OperationType: "CONNECT_CREATE_APP_VM_ONLY",
		LayerName:     "ActiveAppLayer",
	}
	if !reflect.DeepEqual(step.Config, wantStepCfg) {
		t.Errorf("step.Config = %+v, want %+v", step.Config, wantStepCfg)
	}

	art := s.Artifact(stateWithGenerated("gd-connect"))
	if art.OperationType != "CONNECT_CREATE_APP_VM_ONLY" {
		t.Errorf("Artifact.OperationType = %v, want CONNECT_CREATE_APP_VM_ONLY", art.OperationType)
	}
	if art.TargetLayerName != "ActiveAppLayer" {
		t.Errorf("Artifact.TargetLayerName = %v, want ActiveAppLayer", art.TargetLayerName)
	}
	if art.TargetLayerVersionName != "" {
		t.Errorf("Artifact.TargetLayerVersionName = %q, want empty (ConnectOnly has no version)", art.TargetLayerVersionName)
	}
	if got := art.StateData["generated_data"]; got != "gd-connect" {
		t.Errorf("Artifact.StateData[generated_data] = %v, want gd-connect", got)
	}
}
