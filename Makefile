# Variables
IMAGE_NAME := avatarui
DOCKERFILE := ./Dockerfile
GO_BINARY := go-avatar
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
 
.PHONY: build
build: go

go:
	go mod tidy && go build -o output/$(GO_BINARY)
	chmod +x  output/$(GO_BINARY)
	cp index.html output
docker-build: build
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
up: 
	docker-compose up -d $(IMAGE_NAME)
clean: 
	rm -rf output && mkdir output