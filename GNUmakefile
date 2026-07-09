# Copyright (c) Citrix, Inc.
# SPDX-License-Identifier: Apache-2.0

PLUGIN_NAME=packer-plugin-citrix
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)
PACKER_SDC?=$(shell go env GOPATH)/bin/packer-sdc

.PHONY: build test test-coverage test-coverage-check generate install-packer-sdc plugin-check

# Auto-generated files excluded from coverage:
#   - elm-client/elm_client.go         : gowsdl-generated SOAP stub
#   - **/*.hcl2spec.go                 : packer-sdc mapstructure-to-hcl2 output
#   - main.go / version/version.go     : entrypoint and version constant
COVERAGE_EXCLUDE := elm_client\.go|hcl2spec\.go|/main\.go|version/version\.go
COVERAGE_MIN     := 80.0

build:
	@go build -o $(PLUGIN_NAME).exe

test:
	@go test ./...

test-coverage:
	@go test -coverpkg=./... -coverprofile=coverage.out -covermode=atomic ./...
	@grep -v -E "$(COVERAGE_EXCLUDE)" coverage.out > coverage-handwritten.out
	@echo ""
	@echo "=== Hand-written code coverage (auto-gen excluded) ==="
	@go tool cover -func=coverage-handwritten.out | tail -1

test-coverage-check: test-coverage
	@total=$$(go tool cover -func=coverage-handwritten.out | tail -1 | awk '{print $$3}' | tr -d '%'); \
	awk -v t=$$total -v m=$(COVERAGE_MIN) 'BEGIN{ if (t+0 < m+0) { printf "FAIL: coverage %.1f%% < %.1f%%\n", t, m; exit 1 } else { printf "OK: coverage %.1f%% >= %.1f%%\n", t, m } }'

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
