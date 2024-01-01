GO		  ?= go
PWD 	  := $(shell pwd)
GOPATH	:= $(shell $(GO) env GOPATH)
SHELL 	:= /bin/bash

INSTALL_PATH := $(PWD)/bin
LIBRARY_PATH := $(PWD)/lib

# golangci-lint
GOLANGCI_LINT_VERSION := 1.53.1
GOLANGCI_LINT_OUTPUT := $(shell $(INSTALL_PATH)/golangci-lint --version 2>/dev/null)
INSTALL_GOLANGCI_LINT := $(findstring $(GOLANGCI_LINT_VERSION), $(GOLANGCI_LINT_OUTPUT))

fmt:
ifdef GO_DIFF_FILES
	@echo "Running $@ check"
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh $(GO_DIFF_FILES)
else
	@echo "Running $@ check"
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh core/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh global/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh sdk/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh tests/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh searcher/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh web/
endif

#TODO: Check code specifications by golangci-lint
static-check: getdeps
	@echo "Running $@ check"
	@source $(PWD)/scripts/setenv.sh && GO111MODULE=on $(INSTALL_PATH)/golangci-lint run --timeout=30m --config $(PWD)/.golangci.yml; cd pkg && GO111MODULE=on $(INSTALL_PATH)/golangci-lint run --timeout=30m --config $(PWD)/.golangci.yml
