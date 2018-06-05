PACKAGE   = http
DATE     ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell git describe --tags --always --dirty --match = v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

GO        = go
GODOC     = godoc
GOFMT     = gofmt
GOLINT    = gometalinter

MONITOR   = monitor

V         = 0
Q         = $(if $(filter 1,$V),,@)
M         = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all

all: monitor

# Executables
monitor:
	$(info $(M) building executable monitor…) @ ## Build program binary
	$Q cd cmd/$(MONITOR) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(MONITOR)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(MONITOR)_$(VERSION) bin/$(PACKAGE)_$(MONITOR)

# Dependencies
.PHONY: dep
dep:
	$(info $(M) building vendor…) @
	$Q dep ensure

# Check
.PHONY: check
check: lint test

# Tests
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Tools
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…) @
	$Q GO_VENDOR=1 $(GOLINT) "--vendor" \
					"--disable=gotype" \
					"--disable=vetshadow" \
					"--disable=gocyclo" \
					"--fast" \
					"--json" \
					"./..." \
			| grep -v schema.gen.go

.PHONY: fmt
fmt:
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc:
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: clean
clean:
	$(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin/$(PACKAGE)_*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
