# Contributing to packer-plugin-citrix

Thank you for your interest in contributing. This document covers how to set up
your development environment, run tests, and submit changes.

## Prerequisites

- [Go 1.26.0+](https://golang.org/doc/install)
- [HashiCorp Packer 1.7.0+](https://developer.hashicorp.com/packer/downloads)
- `packer-sdc` — required only if modifying config structs:
  ```shell
  go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@latest
  ```

## Building

```shell
go build -o packer-plugin-citrix.exe
```

## Running Tests

```shell
# Unit tests
go test ./...

# Verbose output for a specific package
go test ./builder/citrix/applayering/... -v
```

Acceptance tests (`*_acc_test.go`) require a live ELM environment and are not
run in CI. To run them locally, set the required environment variables and use
the `acc` build tag:

```shell
go test -tags acc ./...
```

## Regenerating HCL2 Specs

After modifying any config struct in `builder/citrix/applayering/config/config.go`
or `provisioner/citrix/applayering/provisioner.go`, regenerate the HCL2 spec files:

```shell
go generate ./...
```

> `elm-client/elm_client.go` is auto-generated from the ELM WSDL via `gowsdl`.
> Do not edit it manually — all helper logic belongs in `elm-client/soap_helpers.go`.

## Submitting Changes

1. Fork the repository and create a feature branch.
2. Make your changes, including tests where applicable.
3. Ensure `go test ./...` passes and `go generate ./...` is up to date.
4. Open a pull request against `master` with a clear description of the change.

## Code Style

- Follow standard Go formatting (`gofmt`).
- Keep the strategy pattern in place when adding new operation types — add a new
  `*Strategy` struct in `builder/citrix/applayering/strategy.go`; do not modify
  `Builder.Run()`.
- All new config fields must have a mapstructure tag and a doc comment for
  `packer-sdc struct-markdown` to pick up.
