# Copyright (c) Citrix, Inc.

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

source "citrix-applayering" "revision_app_final" {
  elm_server          = "your-elm-server.example.com"
  elm_username        = "Administrator"
  elm_password        = var.elm_password
  insecure_connection = true

  revision_app_layer {
    os_layer_name                  = "MyOsLayer"
    os_layer_version_name          = "v1.0"
    # platform_layer_name / platform_layer_version_name are optional; a platform layer is not required
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

build {
  sources = [
    "source.citrix-applayering.revision_app_final"
  ]

  provisioner "citrix-layerartifact" {
    # handle_run_once = true  # Set to true to auto-process RunOnce registry entries.
    # When false (default), the provisioner exits successfully if RunOnce entries
    # are detected, leaving the user to handle them manually via the CE VM.
  }
}
