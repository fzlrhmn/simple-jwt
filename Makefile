.DEFAULT_GOAL	:= build
SHELL 		  	:= /bin/bash
BUILD_DIR   	:= build
BINARY 	      := user
CGO_ENABLED   := 0
GOTEST				:= go test
IS_OK         := ok true y 1

ifndef GOOS
  GOOS := $(shell go env GOHOSTOS)
endif

ifndef GOARCH
	GOARCH := $(shell go env GOHOSTARCH)
endif

ifneq ($(filter $(TEST_VERBOSE),$(IS_OK)),)
  GOTEST += -v
endif

.PHONY: build
build:
	dep ensure -v
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY) ./userd

.PHONY: quickbuild
quickbuild:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY) ./userd