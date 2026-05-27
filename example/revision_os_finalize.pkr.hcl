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

source "citrix-applayering" "revision_os_final" {
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

build {
  sources = [
    "source.citrix-applayering.revision_os_final"
  ]

  provisioner "citrix-layerartifact" {
    # handle_run_once = true  # Set to true to auto-process RunOnce registry entries.
    # When false (default), the provisioner exits successfully if RunOnce entries
    # are detected, leaving the user to handle them manually via the CE VM.
  }
}
