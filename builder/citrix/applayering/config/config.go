// Copyright (c) Citrix, Inc.

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,CreateAppConfig,CreatePlatformConfig,RevisionOsConfig,RevisionPlatformConfig,RevisionAppConfig,ConnectOnlyConfig

// Package config defines the HCL2-decodable configuration types for the
// citrix-applayering builder. Exactly one nested operation block must be set
// in a given source block; the builder dispatches to the matching strategy.
package config

import (
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// Config is the top-level builder configuration.
// Shared connection fields live here; exactly one operation block must be set.
type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	// ELMServer is the URL (or hostname) of the ELM server.
	ELMServer string `mapstructure:"elm_server" validate:"required"`

	// ELMUsername is the ELM login username.
	ELMUsername string `mapstructure:"elm_username" validate:"required"`

	// ELMPassword is the ELM login password.
	ELMPassword string `mapstructure:"elm_password" validate:"required"`

	// InsecureConnection skips TLS certificate verification. Use only for
	// testing with self-signed certificates.
	InsecureConnection bool `mapstructure:"insecure_connection"`

	// WaitForIpTimeout is the maximum number of seconds to wait for the
	// packaging machine to report its IP address. Defaults to 600.
	WaitForIpTimeout int `mapstructure:"wait_for_ip_timeout"`

	// Comm holds the communicator (SSH/WinRM) configuration. Defaults to WinRM.
	Comm communicator.Config `mapstructure:",squash"`

	Ctx interpolate.Context `mapstructure:"-"`

	// Exactly one of the following operation blocks must be present.

	// CreateApp configures a CREATE_APP_LAYER operation.
	CreateApp *CreateAppConfig `mapstructure:"create_app_layer" validate:"required_without_all=CreatePlatform RevisionOs RevisionPlatform RevisionApp ConnectOnly"`

	// CreatePlatform configures a CREATE_PLATFORM_LAYER operation.
	CreatePlatform *CreatePlatformConfig `mapstructure:"create_platform_layer" validate:"required_without_all=CreateApp RevisionOs RevisionPlatform RevisionApp ConnectOnly"`

	// RevisionOs configures a REVISION_OS_LAYER operation.
	RevisionOs *RevisionOsConfig `mapstructure:"revision_os_layer" validate:"required_without_all=CreateApp CreatePlatform RevisionPlatform RevisionApp ConnectOnly"`

	// RevisionPlatform configures a REVISION_PLATFORM_LAYER operation.
	RevisionPlatform *RevisionPlatformConfig `mapstructure:"revision_platform_layer" validate:"required_without_all=CreateApp CreatePlatform RevisionOs RevisionApp ConnectOnly"`

	// RevisionApp configures a REVISION_APP_LAYER operation.
	RevisionApp *RevisionAppConfig `mapstructure:"revision_app_layer" validate:"required_without_all=CreateApp CreatePlatform RevisionOs RevisionPlatform ConnectOnly"`

	// ConnectOnly configures a CONNECT_* debug operation that connects to the
	// packaging VM without executing a layer operation.
	ConnectOnly *ConnectOnlyConfig `mapstructure:"connect_only" validate:"required_without_all=CreateApp CreatePlatform RevisionOs RevisionPlatform RevisionApp"`
}

// CreateAppConfig holds fields for the create_app_layer operation.
type CreateAppConfig struct {
	// OsLayerName is the name of the OS layer that hosts the new app layer.
	OsLayerName string `mapstructure:"os_layer_name" validate:"required"`

	// OsLayerVersionName is the version of the OS layer.
	OsLayerVersionName string `mapstructure:"os_layer_version_name" validate:"required"`

	// PlatformLayerName is the optional platform layer to associate.
	PlatformLayerName string `mapstructure:"platform_layer_name"`

	// PlatformLayerVersionName is the version of the optional platform layer.
	PlatformLayerVersionName string `mapstructure:"platform_layer_version_name"`

	// PrerequisiteLayers is an optional list of prerequisite app layers in
	// "LayerName:VersionName" format.
	PrerequisiteLayers []string `mapstructure:"prerequisite_layers"`

	// LayerName is the name of the new app layer.
	LayerName string `mapstructure:"layer_name" validate:"required"`

	// PackagingDiskFileName is the optional packaging disk filename. Defaults
	// to LayerName.
	PackagingDiskFileName string `mapstructure:"packaging_disk_file_name"`

	// PlatformConnectorConfigName is the name of the connector configuration
	// defined in the ELM that controls how the packaging machine is deployed
	// to your hypervisor.
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name" validate:"required"`

	// IconId is the icon ID. Defaults to 196608.
	IconId int64 `mapstructure:"icon_id"`

	// VersionName is the name of the new version.
	VersionName string `mapstructure:"version_name" validate:"required"`

	// VersionDescription is an optional description for the new version.
	VersionDescription string `mapstructure:"version_description"`

	// VersionSizeGb is the packaging disk size in GB. Defaults to 10.
	VersionSizeGb int32 `mapstructure:"version_size_gb"`

	// Comment is an optional reason/comment for the operation.
	Comment string `mapstructure:"comment"`

	// SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
	// on failure. By default (false) the work ticket is cancelled automatically.
	SkipCleanupOnFailure bool `mapstructure:"skip_cleanup_on_failure"`
}

// CreatePlatformConfig holds fields for the create_platform_layer operation.
type CreatePlatformConfig struct {
	// HypervisorPlatform is the hypervisor platform type (e.g. "vmware").
	HypervisorPlatform string `mapstructure:"hypervisor_platform" validate:"required"`

	// ProvisioningPlatform is the provisioning platform (e.g. "MCS" or "None").
	ProvisioningPlatform string `mapstructure:"provisioning_platform" validate:"required"`

	// BrokerPlatform is the broker platform (e.g. "Citrix VAD" or "None").
	BrokerPlatform string `mapstructure:"broker_platform" validate:"required"`

	// OsLayerName is the OS layer that the new platform layer will target.
	OsLayerName string `mapstructure:"os_layer_name" validate:"required"`

	// OsLayerVersionName is the version of the OS layer.
	OsLayerVersionName string `mapstructure:"os_layer_version_name" validate:"required"`

	// LayerName is the name of the new platform layer.
	LayerName string `mapstructure:"layer_name" validate:"required"`

	// PackagingDiskFileName is the optional packaging disk filename. Defaults
	// to LayerName.
	PackagingDiskFileName string `mapstructure:"packaging_disk_file_name"`

	// PlatformConnectorConfigName is the name of the connector configuration
	// defined in the ELM that controls how the packaging machine is deployed
	// to your hypervisor.
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name" validate:"required"`

	// IconId is the icon ID. Defaults to 196608.
	IconId int64 `mapstructure:"icon_id"`

	// VersionName is the name of the new version.
	VersionName string `mapstructure:"version_name" validate:"required"`

	// VersionDescription is an optional description.
	VersionDescription string `mapstructure:"version_description"`

	// VersionSizeGb is the packaging disk size in GB. Defaults to 10.
	VersionSizeGb int32 `mapstructure:"version_size_gb"`

	// Comment is an optional reason/comment for the operation.
	Comment string `mapstructure:"comment"`

	// SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
	// on failure. By default (false) the work ticket is cancelled automatically.
	SkipCleanupOnFailure bool `mapstructure:"skip_cleanup_on_failure"`
}

// RevisionOsConfig holds fields for the revision_os_layer operation.
type RevisionOsConfig struct {
	// LayerName is the name of the existing OS layer to revise.
	LayerName string `mapstructure:"layer_name" validate:"required"`

	// BaseVersionName is the version to base the revision on.
	BaseVersionName string `mapstructure:"base_version_name" validate:"required"`

	// PackagingDiskFileName is the optional packaging disk filename. Defaults
	// to LayerName.
	PackagingDiskFileName string `mapstructure:"packaging_disk_file_name"`

	// PlatformConnectorConfigName is the name of the connector configuration
	// defined in the ELM that controls how the packaging machine is deployed
	// to your hypervisor.
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name" validate:"required"`

	// VersionName is the name of the new revision.
	VersionName string `mapstructure:"version_name" validate:"required"`

	// VersionDescription is an optional description.
	VersionDescription string `mapstructure:"version_description"`

	// VersionSizeGb is the packaging disk size in GB. Defaults to 10.
	VersionSizeGb int32 `mapstructure:"version_size_gb"`

	// Comment is an optional reason/comment for the operation.
	Comment string `mapstructure:"comment"`

	// SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
	// on failure. By default (false) the work ticket is cancelled automatically.
	SkipCleanupOnFailure bool `mapstructure:"skip_cleanup_on_failure"`
}

// RevisionPlatformConfig holds fields for the revision_platform_layer operation.
type RevisionPlatformConfig struct {
	// HypervisorPlatform is the optional override for the hypervisor platform
	// type. Defaults to the base version's value.
	HypervisorPlatform string `mapstructure:"hypervisor_platform"`

	// ProvisioningPlatform is the optional override for the provisioning
	// platform. Defaults to the base version's value.
	ProvisioningPlatform string `mapstructure:"provisioning_platform"`

	// BrokerPlatform is the optional override for the broker platform.
	// Defaults to the base version's value.
	BrokerPlatform string `mapstructure:"broker_platform"`

	// OsLayerName is the OS layer to reference.
	OsLayerName string `mapstructure:"os_layer_name" validate:"required"`

	// OsLayerVersionName is the version of the OS layer.
	OsLayerVersionName string `mapstructure:"os_layer_version_name" validate:"required"`

	// LayerName is the name of the existing platform layer to revise.
	LayerName string `mapstructure:"layer_name" validate:"required"`

	// BaseVersionName is the version to base the revision on.
	BaseVersionName string `mapstructure:"base_version_name" validate:"required"`

	// PackagingDiskFileName is the optional packaging disk filename. Defaults
	// to LayerName.
	PackagingDiskFileName string `mapstructure:"packaging_disk_file_name"`

	// PlatformConnectorConfigName is the name of the connector configuration
	// defined in the ELM that controls how the packaging machine is deployed
	// to your hypervisor.
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name" validate:"required"`

	// VersionName is the name of the new revision.
	VersionName string `mapstructure:"version_name" validate:"required"`

	// VersionDescription is an optional description.
	VersionDescription string `mapstructure:"version_description"`

	// VersionSizeGb is the packaging disk size in GB. Defaults to 10.
	VersionSizeGb int32 `mapstructure:"version_size_gb"`

	// Comment is an optional reason/comment for the operation.
	Comment string `mapstructure:"comment"`

	// SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
	// on failure. By default (false) the work ticket is cancelled automatically.
	SkipCleanupOnFailure bool `mapstructure:"skip_cleanup_on_failure"`
}

// RevisionAppConfig holds fields for the revision_app_layer operation.
type RevisionAppConfig struct {
	// OsLayerName is the OS layer to reference.
	OsLayerName string `mapstructure:"os_layer_name" validate:"required"`

	// OsLayerVersionName is the version of the OS layer.
	OsLayerVersionName string `mapstructure:"os_layer_version_name" validate:"required"`

	// PlatformLayerName is the optional platform layer to associate.
	PlatformLayerName string `mapstructure:"platform_layer_name"`

	// PlatformLayerVersionName is the version of the optional platform layer.
	PlatformLayerVersionName string `mapstructure:"platform_layer_version_name"`

	// LayerName is the name of the existing app layer to revise.
	LayerName string `mapstructure:"layer_name" validate:"required"`

	// BaseVersionName is the version to base the revision on. Empty means
	// auto-latest.
	BaseVersionName string `mapstructure:"base_version_name"`

	// PrerequisiteLayers is an optional list of prerequisite app layers in
	// "LayerName:VersionName" format.
	PrerequisiteLayers []string `mapstructure:"prerequisite_layers"`

	// PackagingDiskFileName is the optional packaging disk filename. Defaults
	// to LayerName.
	PackagingDiskFileName string `mapstructure:"packaging_disk_file_name"`

	// PlatformConnectorConfigName is the name of the connector configuration
	// defined in the ELM that controls how the packaging machine is deployed
	// to your hypervisor.
	PlatformConnectorConfigName string `mapstructure:"platform_connector_config_name" validate:"required"`

	// VersionName is the name of the new revision.
	VersionName string `mapstructure:"version_name" validate:"required"`

	// VersionDescription is an optional description.
	VersionDescription string `mapstructure:"version_description"`

	// VersionSizeGb is the packaging disk size in GB. Defaults to 10.
	VersionSizeGb int32 `mapstructure:"version_size_gb"`

	// Comment is an optional reason/comment for the operation.
	Comment string `mapstructure:"comment"`

	// SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
	// on failure. By default (false) the work ticket is cancelled automatically.
	SkipCleanupOnFailure bool `mapstructure:"skip_cleanup_on_failure"`
}

// ConnectOnlyConfig holds fields for CONNECT_* debug operations.
// These operations skip the layer creation step and instead connect to a
// packaging machine that is already running in the ELM — useful for recovering
// from a failed Packer run or for manual debugging of an active packaging session.
type ConnectOnlyConfig struct {
	// OperationType identifies which type of active packaging session to connect
	// to. It must correspond to the operation that originally started the session:
	// CONNECT_CREATE_APP_VM_ONLY — reconnect to a create_app_layer session,
	// CONNECT_CREATE_PLATFORM_VM_ONLY — reconnect to a create_platform_layer session,
	// CONNECT_REVISION_APP_VM_ONLY — reconnect to a revision_app_layer session,
	// CONNECT_REVISION_PLATFORM_VM_ONLY — reconnect to a revision_platform_layer session,
	// CONNECT_REVISION_OS_VM_ONLY — reconnect to a revision_os_layer session.
	OperationType string `mapstructure:"operation_type" validate:"required"`

	// LayerName is the name of the layer with an active packaging session in
	// the ELM. This is used to locate the open work ticket for the session.
	LayerName string `mapstructure:"layer_name" validate:"required"`
}
