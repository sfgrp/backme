PROJ_NAME = backme

VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell TZ=UTC date +'%Y-%m-%d_%H:%M:%ST%Z')

NO_C = CGO_ENABLED=0
FLAGS_LD = -ldflags "-X github.com/dimus/$(PROJ_NAME)/pkg.Build=$(DATE) \
                     -X github.com/dimus/$(PROJ_NAME)/pkg.Version=$(VERSION)"
FLAGS_REL = -trimpath -ldflags "-s -w -X github.com/dimus/$(PROJ_NAME)/pkg.Build=$(DATE)"
FLAGS_SHARED = $(NO_C) GOARCH=amd64

RELEASE_DIR = /tmp
TEST_OPTS = -v -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic

GOCMD = go
GOTEST = $(GOCMD) test
GOVET = $(GOCMD) vet
GOBUILD = $(GOCMD) build $(FLAGS_LD)
GORELEASE = $(GOCMD) build $(FLAGS_REL)
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install $(FLAGS_LD)
all: install

## Dependencies
deps: ## Download dependencies
	$(GOCMD) mod download;

## Tools
tools: deps ## Install tools
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

## Test:
test: ## Run the tests of the project
	$(GOTEST) $(TEST_OPTS) ./...

## Build:
build: ## Build binary
	$(NO_C) $(GOCMD) build \
		-o $(PROJ_NAME) \
		$(FLAGS_LD) \
		.

## Build Release
buildrel: ## Build binary without debug info and with hardcoded version
	$(NO_C) $(GOCMD) build \
		-o $(PROJ_NAME) \
		$(FLAGS_REL) \
		.

## Install:
install: ## Build and install binary
	$(NO_C) $(GOINSTALL)

## Release
release: ## Build and package binaries for a release
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=linux $(GORELEASE); \
	tar zcvf /tmp/$(PROJ_NAME)-$(VER)-linux.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=darwin $(GORELEASE); \
	tar zcvf /tmp/$(PROJ_NAME)-$(VER)-mac.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=windows $(GORELEASE); \
	zip -9 /tmp/$(PROJ_NAME)-$(VER)-win-64.zip $(PROJ_NAME).exe; \
	$(GOCLEAN);


## Version
version: ## Display current version
	@echo $(VERSION)
