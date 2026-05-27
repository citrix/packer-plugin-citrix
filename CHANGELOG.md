# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-04-22

### Changed

- When ELM returns `WorkTicketId == 0` (e.g. a conflicting operation on the same
  layer), the plugin now fails immediately and surfaces the ELM error message to
  the admin. Previously the plugin retried with exponential backoff, which masked
  the actual error. The `conflict_retry_timeout` and `conflict_retry_interval`
  config fields have been removed.

### Fixed

- `base_version_name` in `revision_os_layer` was ignored — the SOAP request was not
  passing `BaseLayerRevisionId`, causing ELM to always use the latest version as the
  base. Platform and app revision steps were not affected.
- All five create/revision steps now check for `WorkTicketId == 0` and report the
  ELM error message (previously the plugin would continue with ticket 0 and
  eventually fail with a confusing "work ticket 0 not found" error).

## [1.0.0] - 2026-03-23

### Added

- `citrix-applayering` builder supporting the following ELM operations:
  - `create_app_layer` — create a new app layer
  - `create_platform_layer` — create a new platform layer
  - `revision_os_layer` — add a new version to an existing OS layer
  - `revision_platform_layer` — add a new version to an existing platform layer
  - `revision_app_layer` — add a new version to an existing app layer
  - `connect_only` — debug: connect to an existing CE VM without calling ELM
- `citrix-layerartifact` provisioner that polls BlockFinalize bits and calls
  `ShutdownForFinalize` to publish the completed layer version to ELM
- WinRM communicator support for connecting to CE VMs
- Preflight validation for all operation types (layer existence, name conflicts)
- Configurable IP-wait timeout (`wait_for_ip_timeout`)
- `insecure_connection` option for self-signed certificate environments
- HCL2-native configuration with full `packer-sdc` generated specs
- Example templates for all six operation types
