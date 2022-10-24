GOPATH ?= $(shell go env GOPATH)

# List of effective go files
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" | egrep -v "^\./\.go" | grep -v _test.go)

# List of packages except testsutils
PACKAGES ?= $(shell go list ./... | grep -v "mock" )

# Build folder
BUILD_FOLDER = build

# Test coverage variables
COVERAGE_BUILD_FOLDER = $(BUILD_FOLDER)/coverage

UNIT_COVERAGE_OUT = $(COVERAGE_BUILD_FOLDER)/ut_cov.out
UNIT_COVERAGE_HTML =$(COVERAGE_BUILD_FOLDER)/ut_index.html

# Test lint variables
GOLANGCI_VERSION = v1.44.0

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OPEN = xdg-open
endif
ifeq ($(UNAME_S),Darwin)
	OPEN = open
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build/coverage:
	@mkdir -p build/coverage

unit-test: build/coverage
	@go test -covermode=count -coverprofile $(UNIT_COVERAGE_OUT) $(PACKAGES) -v -timeout 1s

unit-test-cov: unit-test
	@go tool cover -html=$(UNIT_COVERAGE_OUT) -o $(UNIT_COVERAGE_HTML)

fix-lint: ## Run linter to fix issues
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_VERSION) golangci-lint run --fix

# @misspell -error $(GOFILES)
test-lint: ## Check linting
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_VERSION) golangci-lint run -v

mockgen-install:
	@type mockgen >/dev/null 2>&1 || {   \
		echo "Installing mockgen..."; \
		go install github.com/golang/mock/mockgen@v1.6.0;  \
	}

mockgen: mockgen-install
	$(GOPATH)/bin/mockgen -source pkg/defender/client/client.go -destination pkg/defender/client/mock/client.go -package mock Client