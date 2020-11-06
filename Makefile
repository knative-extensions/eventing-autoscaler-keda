##################################################
# Variables                                      #
##################################################
VERSION		   ?= master
IMAGE_REGISTRY ?= docker.io
IMAGE_REPO     ?= avarghese23

IMAGE = $(IMAGE_REGISTRY)/$(IMAGE_REPO)/eventing-autoscaler-keda:$(VERSION)

ARCH       ?=amd64
CGO        ?=0
TARGET_OS  ?=linux

GIT_VERSION = $(shell git describe --always --abbrev=7)
GIT_COMMIT  = $(shell git rev-list -1 HEAD)
DATE        = $(shell date -u +"%Y.%m.%d.%H.%M.%S")


##################################################
# Build                                          #
##################################################
GO_BUILD_VARS= GO111MODULE=on CGO_ENABLED=$(CGO) GOOS=$(TARGET_OS) GOARCH=$(ARCH)

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: build
build: gofmt
	rm -rf ./build/*
	mkdir -p ./build
	$(GO_BUILD_VARS) go build -mod vendor -o build/eventing-autoscaler-keda cmd/controller/main.go 

.PHONY: build-image
build-image: build
	docker build -t $(IMAGE) ./
