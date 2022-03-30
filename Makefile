GOCMD=go
GOOS=linux
CGO_ENABLED=1
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=dataserver
VERSION=1.0.1

all: build docker

build:
	$(GOBUILD) -o $(BINARY_NAME)
	chmod 777 $(BINARY_NAME)

docker:
	docker build -t $(BINARY_NAME):$(VERSION) .
