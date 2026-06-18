# Packer Plugin for Citrix App Layering

The Packer Plugin for Citrix App Layering automates layer creation and revision workflows on
[Citrix App Layering][docs-citrix] (ELM — Enterprise Layer Manager) using [HashiCorp Packer][packer].

The plugin includes one builder and one provisioner:

**Builder**

- `citrix-applayering` — Connects to ELM via its SOAP API to create a Compositing Engine (CE) VM,
  waits for the VM to obtain an IP address, and connects via WinRM so that Packer provisioners
  can install software inside the CE VM.

**Provisioner**

- `citrix-layerartifact` — Waits for the CE VM to finish background tasks (BlockFinalize) and then
  calls `ShutdownForFinalize` to publish the completed layer version back to ELM.

## Requirements

- [Go 1.26.4][golang-install] — required only if building the plugin from source.
- Citrix App Layering (ELM) appliance reachable from the machine running Packer.
- ELM TCP port 443 open from the Packer host to the ELM appliance.
- CE VM TCP port 5985 (WinRM HTTP) open from the Packer host to the hypervisor network.

> **Note**: This plugin targets Windows packaging machines only. The CE VM created by ELM is a
> WinPE-based Windows VM; only the WinRM communicator is supported. SSH is not supported.

## Quick Start

```hcl
packer {
  required_version = ">= 1.7.0"
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
    packaging_disk_file_name       = "MyAppLayer_v1.0"
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

Run with:

```shell
export PKR_VAR_elm_password="<your-elm-password>"
export PKR_VAR_vm_password="<your-ce-vm-password>"
packer init .
packer build my_app.pkr.hcl
```

See the [example](./example/) folder for all supported operation types.

The builder uses Packer's built-in WinRM communicator to connect to the CE VM. For
WinRM connection details and the full list of options, see
[Establish a WinRM Connection][docs-winrm].

## Supported Operations

| Operation block | Description |
|-----------------|-------------|
| `create_app_layer` | Create a new app layer |
| `create_platform_layer` | Create a new platform layer |
| `revision_os_layer` | Add a new version to an existing OS layer |
| `revision_platform_layer` | Add a new version to an existing platform layer |
| `revision_app_layer` | Add a new version to an existing app layer |
| `connect_only` | Debug: connect to an existing CE VM without calling ELM |

Exactly one operation block must be present in a source block.

## Provisioner Reference

### `citrix-layerartifact`

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `handle_run_once` | bool | `false` | Controls automatic processing of HKLM `RunOnce` registry entries when `IC_RUNONCE` is set in `BlockFinalize`. When `false` (default), encountering `IC_RUNONCE` logs an info message and the provisioner exits successfully, leaving the user to handle RunOnce entries manually via the CE VM. When `true`, entries are processed one at a time automatically; note that many RunOnce entries expect an interactive Windows session, which the headless CE VM does not provide. |

## Installation

### Using Pre-built Releases

Packer v1.7.0 and later supports the `packer init` command for automatic plugin installation.
Copy the `required_plugins` block from the Quick Start section above into your configuration and run
`packer init`.

You can also download [pre-built binary releases][releases-citrix-plugin] from GitHub and install
them manually following the [Packer plugin installation documentation][docs-packer-plugin-install].

### Building from Source

```shell
git clone https://github.com/citrix/packer-plugin-citrix.git
cd packer-plugin-citrix
go build -o packer-plugin-citrix.exe
```

Copy the resulting `packer-plugin-citrix.exe` to your Packer plugin directory following the
[Packer plugin installation documentation][docs-packer-plugin-install].

## Known Limitations

- **Windows CE VM only**: The CE VM created by ELM is WinPE-based. Only the WinRM communicator
  is supported; SSH is not available.
- **ELM SOAP API**: The plugin communicates with ELM over SOAP on port 443. Self-signed
  certificates are common in lab environments; use `insecure_connection = true` in those cases.

## Contributing

See [AGENTS.md](./AGENTS.md) for architecture and development guidance.

Run unit tests:

```shell
go test ./...
```

Regenerate HCL2 spec after modifying config structs:

```shell
go generate ./...
```

[packer]: https://www.packer.io
[docs-citrix]: https://docs.citrix.com/en-us/citrix-app-layering
[golang-install]: https://golang.org/doc/install
[releases-citrix-plugin]: https://github.com/citrix/packer-plugin-citrix/releases
[docs-packer-plugin-install]: https://developer.hashicorp.com/packer/docs/plugins/install-plugins
[docs-winrm]: https://developer.hashicorp.com/packer/docs/communicators/winrm
