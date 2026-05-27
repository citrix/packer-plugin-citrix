Type: `citrix-applayering`

The citrix-applayering builder connects to a Citrix App Layering (ELM) appliance
via its SOAP API. It creates a Compositing Engine (CE) VM for layer packaging,
waits for the VM to obtain an IP address, and connects via WinRM so that Packer
provisioners can install software inside the CE VM. After provisioning,
the citrix-layerartifact provisioner finalizes the layer version in the ELM.

Exactly one operation block must be present in a source block. The supported
operations are:

| Operation block            | Description                                    |
|----------------------------|------------------------------------------------|
| `create_app_layer`         | Create a new app layer                         |
| `create_platform_layer`    | Create a new platform layer                    |
| `revision_os_layer`        | Add a new version to an existing OS layer      |
| `revision_platform_layer`  | Add a new version to an existing platform layer|
| `revision_app_layer`       | Add a new version to an existing app layer     |
| `connect_only`             | Debug: connect to an existing CE VM            |

## Configuration Reference

<!-- Code generated from the comments of the Config struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

Config is the top-level builder configuration.
Shared connection fields live here; exactly one operation block must be set.

<!-- End of code generated from the comments of the Config struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the Config struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `elm_server` (string) - ELMServer is the URL (or hostname) of the ELM server.

- `elm_username` (string) - ELMUsername is the ELM login username.

- `elm_password` (string) - ELMPassword is the ELM login password.

- `insecure_connection` (bool) - InsecureConnection skips TLS certificate verification. Use only for
  testing with self-signed certificates.

- `wait_for_ip_timeout` (int) - WaitForIpTimeout is the maximum number of seconds to wait for the
  packaging machine to report its IP address. Defaults to 600.

- `create_app_layer` (\*CreateAppConfig) - CreateApp configures a CREATE_APP_LAYER operation.

- `create_platform_layer` (\*CreatePlatformConfig) - CreatePlatform configures a CREATE_PLATFORM_LAYER operation.

- `revision_os_layer` (\*RevisionOsConfig) - RevisionOs configures a REVISION_OS_LAYER operation.

- `revision_platform_layer` (\*RevisionPlatformConfig) - RevisionPlatform configures a REVISION_PLATFORM_LAYER operation.

- `revision_app_layer` (\*RevisionAppConfig) - RevisionApp configures a REVISION_APP_LAYER operation.

- `connect_only` (\*ConnectOnlyConfig) - ConnectOnly configures a CONNECT_* debug operation that connects to the
  packaging VM without executing a layer operation.

<!-- End of code generated from the comments of the Config struct in builder/citrix/applayering/config/config.go; -->


### Create App Layer

<!-- Code generated from the comments of the CreateAppConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

CreateAppConfig holds fields for the create_app_layer operation.

<!-- End of code generated from the comments of the CreateAppConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the CreateAppConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `os_layer_name` (string) - OsLayerName is the name of the OS layer that hosts the new app layer.

- `os_layer_version_name` (string) - OsLayerVersionName is the version of the OS layer.

- `platform_layer_name` (string) - PlatformLayerName is the optional platform layer to associate.

- `platform_layer_version_name` (string) - PlatformLayerVersionName is the version of the optional platform layer.

- `prerequisite_layers` ([]string) - PrerequisiteLayers is an optional list of prerequisite app layers in
  "LayerName:VersionName" format.

- `layer_name` (string) - LayerName is the name of the new app layer.

- `packaging_disk_file_name` (string) - PackagingDiskFileName is the optional packaging disk filename. Defaults
  to LayerName.

- `platform_connector_config_name` (string) - PlatformConnectorConfigName is the name of the connector configuration
  defined in the ELM that controls how the packaging machine is deployed
  to your hypervisor.

- `icon_id` (int64) - IconId is the icon ID. Defaults to 196608.

- `version_name` (string) - VersionName is the name of the new version.

- `version_description` (string) - VersionDescription is an optional description for the new version.

- `version_size_gb` (int32) - VersionSizeGb is the packaging disk size in GB. Defaults to 10.

- `comment` (string) - Comment is an optional reason/comment for the operation.

- `skip_cleanup_on_failure` (bool) - SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
  on failure. By default (false) the work ticket is cancelled automatically.

<!-- End of code generated from the comments of the CreateAppConfig struct in builder/citrix/applayering/config/config.go; -->


### Create Platform Layer

<!-- Code generated from the comments of the CreatePlatformConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

CreatePlatformConfig holds fields for the create_platform_layer operation.

<!-- End of code generated from the comments of the CreatePlatformConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the CreatePlatformConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `hypervisor_platform` (string) - HypervisorPlatform is the hypervisor platform type (e.g. "vmware").

- `provisioning_platform` (string) - ProvisioningPlatform is the provisioning platform (e.g. "MCS" or "None").

- `broker_platform` (string) - BrokerPlatform is the broker platform (e.g. "Citrix VAD" or "None").

- `os_layer_name` (string) - OsLayerName is the OS layer that the new platform layer will target.

- `os_layer_version_name` (string) - OsLayerVersionName is the version of the OS layer.

- `layer_name` (string) - LayerName is the name of the new platform layer.

- `packaging_disk_file_name` (string) - PackagingDiskFileName is the optional packaging disk filename. Defaults
  to LayerName.

- `platform_connector_config_name` (string) - PlatformConnectorConfigName is the name of the connector configuration
  defined in the ELM that controls how the packaging machine is deployed
  to your hypervisor.

- `icon_id` (int64) - IconId is the icon ID. Defaults to 196608.

- `version_name` (string) - VersionName is the name of the new version.

- `version_description` (string) - VersionDescription is an optional description.

- `version_size_gb` (int32) - VersionSizeGb is the packaging disk size in GB. Defaults to 10.

- `comment` (string) - Comment is an optional reason/comment for the operation.

- `skip_cleanup_on_failure` (bool) - SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
  on failure. By default (false) the work ticket is cancelled automatically.

<!-- End of code generated from the comments of the CreatePlatformConfig struct in builder/citrix/applayering/config/config.go; -->


### Revision OS Layer

<!-- Code generated from the comments of the RevisionOsConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

RevisionOsConfig holds fields for the revision_os_layer operation.

<!-- End of code generated from the comments of the RevisionOsConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the RevisionOsConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `layer_name` (string) - LayerName is the name of the existing OS layer to revise.

- `base_version_name` (string) - BaseVersionName is the version to base the revision on.

- `packaging_disk_file_name` (string) - PackagingDiskFileName is the optional packaging disk filename. Defaults
  to LayerName.

- `platform_connector_config_name` (string) - PlatformConnectorConfigName is the name of the connector configuration
  defined in the ELM that controls how the packaging machine is deployed
  to your hypervisor.

- `version_name` (string) - VersionName is the name of the new revision.

- `version_description` (string) - VersionDescription is an optional description.

- `version_size_gb` (int32) - VersionSizeGb is the packaging disk size in GB. Defaults to 10.

- `comment` (string) - Comment is an optional reason/comment for the operation.

- `skip_cleanup_on_failure` (bool) - SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
  on failure. By default (false) the work ticket is cancelled automatically.

<!-- End of code generated from the comments of the RevisionOsConfig struct in builder/citrix/applayering/config/config.go; -->


### Revision Platform Layer

<!-- Code generated from the comments of the RevisionPlatformConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

RevisionPlatformConfig holds fields for the revision_platform_layer operation.

<!-- End of code generated from the comments of the RevisionPlatformConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the RevisionPlatformConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `hypervisor_platform` (string) - HypervisorPlatform is the optional override for the hypervisor platform
  type. Defaults to the base version's value.

- `provisioning_platform` (string) - ProvisioningPlatform is the optional override for the provisioning
  platform. Defaults to the base version's value.

- `broker_platform` (string) - BrokerPlatform is the optional override for the broker platform.
  Defaults to the base version's value.

- `os_layer_name` (string) - OsLayerName is the OS layer to reference.

- `os_layer_version_name` (string) - OsLayerVersionName is the version of the OS layer.

- `layer_name` (string) - LayerName is the name of the existing platform layer to revise.

- `base_version_name` (string) - BaseVersionName is the version to base the revision on.

- `packaging_disk_file_name` (string) - PackagingDiskFileName is the optional packaging disk filename. Defaults
  to LayerName.

- `platform_connector_config_name` (string) - PlatformConnectorConfigName is the name of the connector configuration
  defined in the ELM that controls how the packaging machine is deployed
  to your hypervisor.

- `version_name` (string) - VersionName is the name of the new revision.

- `version_description` (string) - VersionDescription is an optional description.

- `version_size_gb` (int32) - VersionSizeGb is the packaging disk size in GB. Defaults to 10.

- `comment` (string) - Comment is an optional reason/comment for the operation.

- `skip_cleanup_on_failure` (bool) - SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
  on failure. By default (false) the work ticket is cancelled automatically.

<!-- End of code generated from the comments of the RevisionPlatformConfig struct in builder/citrix/applayering/config/config.go; -->


### Revision App Layer

<!-- Code generated from the comments of the RevisionAppConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

RevisionAppConfig holds fields for the revision_app_layer operation.

<!-- End of code generated from the comments of the RevisionAppConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the RevisionAppConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `os_layer_name` (string) - OsLayerName is the OS layer to reference.

- `os_layer_version_name` (string) - OsLayerVersionName is the version of the OS layer.

- `platform_layer_name` (string) - PlatformLayerName is the optional platform layer to associate.

- `platform_layer_version_name` (string) - PlatformLayerVersionName is the version of the optional platform layer.

- `layer_name` (string) - LayerName is the name of the existing app layer to revise.

- `base_version_name` (string) - BaseVersionName is the version to base the revision on. Empty means
  auto-latest.

- `prerequisite_layers` ([]string) - PrerequisiteLayers is an optional list of prerequisite app layers in
  "LayerName:VersionName" format.

- `packaging_disk_file_name` (string) - PackagingDiskFileName is the optional packaging disk filename. Defaults
  to LayerName.

- `platform_connector_config_name` (string) - PlatformConnectorConfigName is the name of the connector configuration
  defined in the ELM that controls how the packaging machine is deployed
  to your hypervisor.

- `version_name` (string) - VersionName is the name of the new revision.

- `version_description` (string) - VersionDescription is an optional description.

- `version_size_gb` (int32) - VersionSizeGb is the packaging disk size in GB. Defaults to 10.

- `comment` (string) - Comment is an optional reason/comment for the operation.

- `skip_cleanup_on_failure` (bool) - SkipCleanupOnFailure, when true, skips cancelling the ELM work ticket
  on failure. By default (false) the work ticket is cancelled automatically.

<!-- End of code generated from the comments of the RevisionAppConfig struct in builder/citrix/applayering/config/config.go; -->


### Connect Only (Debug)

<!-- Code generated from the comments of the ConnectOnlyConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

ConnectOnlyConfig holds fields for CONNECT_* debug operations.
These operations skip the layer creation step and instead connect to a
packaging machine that is already running in the ELM — useful for recovering
from a failed Packer run or for manual debugging of an active packaging session.

<!-- End of code generated from the comments of the ConnectOnlyConfig struct in builder/citrix/applayering/config/config.go; -->


<!-- Code generated from the comments of the ConnectOnlyConfig struct in builder/citrix/applayering/config/config.go; DO NOT EDIT MANUALLY -->

- `operation_type` (string) - OperationType identifies which type of active packaging session to connect
  to. It must correspond to the operation that originally started the session:
  CONNECT_CREATE_APP_VM_ONLY — reconnect to a create_app_layer session,
  CONNECT_CREATE_PLATFORM_VM_ONLY — reconnect to a create_platform_layer session,
  CONNECT_REVISION_APP_VM_ONLY — reconnect to a revision_app_layer session,
  CONNECT_REVISION_PLATFORM_VM_ONLY — reconnect to a revision_platform_layer session,
  CONNECT_REVISION_OS_VM_ONLY — reconnect to a revision_os_layer session.

- `layer_name` (string) - LayerName is the name of the layer with an active packaging session in
  the ELM. This is used to locate the open work ticket for the session.

<!-- End of code generated from the comments of the ConnectOnlyConfig struct in builder/citrix/applayering/config/config.go; -->


## Example Usage

```hcl
variable "elm_password" {
  type      = string
  sensitive = true
}

variable "vm_password" {
  type      = string
  sensitive = true
}

source "citrix-applayering" "my_app" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  create_app_layer {
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    layer_name                     = "MyAppLayer"
    version_name                   = "v1.0"
    version_size_gb                = 10
    platform_connector_config_name = "MyConnectorConfig"
    icon_id                        = 196608
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

build {
  sources = ["source.citrix-applayering.my_app"]

  provisioner "citrix-layerartifact" {}
}
```
