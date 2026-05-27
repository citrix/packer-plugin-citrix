# Copyright (c) Citrix, Inc.
# All supported use cases in one file. Use "packer build -only=citrix-applayering.all_xxx ..."
# to run a specific source.

packer {
  required_plugins {
    citrix = {
      version = ">= 1.0.0"
      source  = "github.com/citrix/citrix"
    }
  }
}

variable "elm_password" {
  type      = string
  sensitive = true
}

variable "vm_password" {
  type      = string
  sensitive = true
}

# ─── Create App Layer (with platform) ────────────────────────────────────────

source "citrix-applayering" "all_create_app" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  create_app_layer {
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    platform_layer_name            = "MyPlatformLayer"
    platform_layer_version_name    = "v1.0"
    layer_name                     = "MyAppLayer"
    packaging_disk_file_name       = "MyAppLayer_v1.0"
    platform_connector_config_name = "MyConnectorConfig"
    icon_id                        = 196608
    version_name                   = "v1.0"
    version_description            = "Application Layer v1.0"
    version_size_gb                = 10
    comment                        = "Initial app layer creation"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

# ─── Create App Layer (without platform) ─────────────────────────────────────

source "citrix-applayering" "all_create_app_no_platform" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  create_app_layer {
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    # platform_layer_name / platform_layer_version_name omitted — no platform layer
    layer_name                     = "MyAppLayerNoPlatform"
    packaging_disk_file_name       = "MyAppLayerNoPlatform_v1.0"
    platform_connector_config_name = "MyConnectorConfig"
    icon_id                        = 196608
    version_name                   = "v1.0"
    version_description            = "Application Layer v1.0 (no platform)"
    version_size_gb                = 10
    comment                        = "Initial app layer creation (no platform)"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

# ─── Create Platform Layer ───────────────────────────────────────────────────

source "citrix-applayering" "all_create_platform" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  create_platform_layer {
    # platform type IDs (required; use "None" to omit provisioning/broker)
    hypervisor_platform            = "vsphere"
    provisioning_platform          = "MCS"
    broker_platform                = "None"
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    layer_name                     = "MyPlatformLayer"
    packaging_disk_file_name       = "MyPlatformLayer_v1.0"
    platform_connector_config_name = "MyConnectorConfig"
    icon_id                        = 196608
    version_name                   = "v1.0"
    version_description            = "Platform Layer v1.0"
    version_size_gb                = 10
    comment                        = "Initial platform layer creation"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

# ─── Revision OS Layer ───────────────────────────────────────────────────────

source "citrix-applayering" "all_revision_os" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  revision_os_layer {
    layer_name                     = "MyOsLayer"
    base_version_name              = "v1.0"
    packaging_disk_file_name       = "MyOsLayer_v2.0"
    platform_connector_config_name = "MyConnectorConfig"
    version_name                   = "v2.0"
    version_description            = "OS Layer v2.0"
    version_size_gb                = 25
    comment                        = "OS layer revision v2.0"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
  winrm_timeout  = "30m"
}

# ─── Revision Platform Layer ────────────────────────────────────────────────

source "citrix-applayering" "all_revision_platform" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  revision_platform_layer {
    hypervisor_platform            = "vsphere"
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    layer_name                     = "MyPlatformLayer"
    base_version_name              = "v1.0"
    packaging_disk_file_name       = "MyPlatformLayer_v2.0"
    platform_connector_config_name = "MyConnectorConfig"
    version_name                   = "v2.0"
    version_description            = "Platform Layer v2.0"
    version_size_gb                = 10
    comment                        = "Platform layer revision v2.0"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

# ─── Revision App Layer ─────────────────────────────────────────────────────

source "citrix-applayering" "all_revision_app" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  revision_app_layer {
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    platform_layer_name            = "MyPlatformLayer"
    platform_layer_version_name    = "v1.0"
    layer_name                     = "MyAppLayer"
    # base_version_name is optional — omit to auto-select latest
    base_version_name              = "v1.0"
    packaging_disk_file_name       = "MyAppLayer_v2.0"
    platform_connector_config_name = "MyConnectorConfig"
    version_name                   = "v2.0"
    version_description            = "App Layer v2.0"
    version_size_gb                = 10
    comment                        = "App layer revision v2.0"
    # skip_cleanup_on_failure = true  # Set to true to skip cancelling the ELM work
    # ticket on failure. By default (false), the work ticket is cancelled automatically
    # when the build fails or is cancelled by the user.
  }

  communicator   = "winrm"
  winrm_username = "Administrator"
  winrm_password = var.vm_password
}

# ─── Build ───────────────────────────────────────────────────────────────────

build {
  sources = [
    "source.citrix-applayering.all_create_app",
    "source.citrix-applayering.all_create_app_no_platform",
    "source.citrix-applayering.all_create_platform",
    "source.citrix-applayering.all_revision_os",
    "source.citrix-applayering.all_revision_platform",
    "source.citrix-applayering.all_revision_app"
  ]

  provisioner "citrix-layerartifact" {
    # handle_run_once = true  # Set to true to auto-process RunOnce registry entries.
    # When false (default), the provisioner exits successfully if RunOnce entries
    # are detected, leaving the user to handle them manually via the CE VM.
  }
}
