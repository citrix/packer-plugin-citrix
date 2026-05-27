# Copyright (c) Citrix, Inc.
# SPDX-License-Identifier: Apache-2.0

PLUGIN_NAME=packer-plugin-citrix
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)
PACKER_SDC?=$(shell go env GOPATH)/bin/packer-sdc

.PHONY: build test generate install-packer-sdc plugin-check

build:
	@go build -o $(PLUGIN_NAME).exe

test:
	@go test ./...

install-packer-sdc: ## Install packer-sdc
	@go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@$(HASHICORP_PACKER_PLUGIN_SDK_VERSION)

plugin-check: install-packer-sdc build
	@$(PACKER_SDC) plugin-check $(PLUGIN_NAME).exe

generate: install-packer-sdc
	@go generate ./...
	@rm -rf .docs
	@$(PACKER_SDC) renderdocs -src docs -partials docs-partials/ -dst .docs/
	@./.web-docs/scripts/compile-to-webdocs.sh "." ".docs" ".web-docs" "citrix"
	@rm -rf ".docs"
