BINARY=infra-bot

GO=go
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOLINT=$(GO) vet

IMAGE_NAME=infra-bot
TAG=latest

LDFLAGS=-X github.com/i-stanko/infra-bot/internal/version.Version=dev \
        -X github.com/i-stanko/infra-bot/internal/version.Commit=dev \
        -X github.com/i-stanko/infra-bot/internal/version.BuildDate=unknown

.PHONY: all build run test lint clean docker-build docker-run

all: build

build:
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BINARY) ./main.go

run: build
	./$(BINARY) start

test:
	$(GOTEST) ./...

lint:
	$(GOLINT) ./...

clean:
	rm -f $(BINARY)

docker-build:
	docker build -t $(IMAGE_NAME):$(TAG) .

docker-run:
	docker run --rm -e TELE_TOKEN=$$TELE_TOKEN $(IMAGE_NAME):$(TAG) start
