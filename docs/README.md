The Citrix App Layering plugin automates layer creation and revision workflows on
[Citrix App Layering](https://docs.citrix.com/en-us/citrix-app-layering) (ELM — Enterprise Layer Manager)
using [HashiCorp Packer](https://www.packer.io).

The plugin connects to the ELM via its SOAP API to create a Compositing Engine (CE) VM,
waits for the VM to obtain an IP address, and connects via WinRM so that Packer provisioners
can install software inside the CE VM. After provisioning, the layer version is finalized
in the ELM and made available for use in image templates.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run
[`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    citrix = {
      source  = "github.com/citrix/citrix"
      version = ">= 1.0.0"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/citrix/citrix
```

### Components

#### Builders

- [citrix-applayering](/packer/integrations/citrix/citrix/latest/components/builder/applayering) -
  The citrix-applayering builder connects to the ELM via its SOAP API to create a
  Compositing Engine (CE) VM for layer packaging.

#### Provisioners

- [citrix-layerartifact](/packer/integrations/citrix/citrix/latest/components/provisioner/layerartifact) -
  The citrix-layerartifact provisioner waits for the CE VM to complete
  pre-finalization tasks and then shuts it down, triggering the ELM to finalize
  and save the completed layer version.
