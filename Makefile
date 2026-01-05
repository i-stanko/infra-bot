APP_NAME := infra-bot
REGISTRY := quay.io
NAMESPACE := istanko

VERSION ?= dev
GIT_SHA := $(shell git rev-parse --short HEAD)

IMAGE_TAG := $(REGISTRY)/$(NAMESPACE)/$(APP_NAME):$(VERSION)-$(GIT_SHA)

TARGETOS ?= linux
TARGETARCH ?= amd64

.PHONY: format build image push clean linux windows mac arm

format:
	gofmt -s -w ./

build: format
	CGO_ENABLED=0 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) \
	go build -o $(APP_NAME) \
	-ldflags "-X github.com/i-stanko/infra-bot/cmd.appVersion=$(VERSION)" .

image:
	docker build \
	--build-arg TARGETOS=$(TARGETOS) \
	--build-arg TARGETARCH=$(TARGETARCH) \
	-t $(IMAGE_TAG) .

push:
	docker push $(IMAGE_TAG)

linux:
	make image TARGETOS=linux TARGETARCH=amd64

arm:
	make image TARGETOS=linux TARGETARCH=arm64

mac:
	make image TARGETOS=darwin TARGETARCH=arm64

windows:
	make image TARGETOS=windows TARGETARCH=amd64

clean:
	rm -f $(APP_NAME)
	docker rmi -f $(IMAGE_TAG) || true
