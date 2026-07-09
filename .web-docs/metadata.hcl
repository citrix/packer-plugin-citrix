# Copyright (c) Citrix, Inc.
# SPDX-License-Identifier: Apache-2.0

# For details on using this integration template, see:
# https://github.com/hashicorp/integration-template
#
# This metadata.hcl file and the adjacent `components` docs directory should
# be kept in a `.web-docs` directory at the root of your plugin repository.
integration {
  name        = "Citrix App Layering"
  description = "The Citrix App Layering plugin automates layer creation and revision workflows on Citrix App Layering (ELM)."
  identifier  = "packer/citrix/citrix"
  flags       = []
  docs {
    process_docs    = true
    readme_location = "./README.md"
    external_url    = "https://github.com/citrix/packer-plugin-citrix"
  }
  license {
    type = "Apache-2.0"
    url  = "https://github.com/citrix/packer-plugin-citrix/blob/main/LICENSE"
  }
  component {
    type = "builder"
    name = "Citrix App Layering"
    slug = "applayering"
  }
  component {
    type = "provisioner"
    name = "Citrix Layer Artifact"
    slug = "layerartifact"
  }
}
