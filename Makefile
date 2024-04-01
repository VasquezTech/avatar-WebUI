# Variables
IMAGE_NAME := avatarui
DOCKERFILE := ./Dockerfile
GO_BINARY := go-avatar
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
re: clean go 
	cd output && ./$(GO_BINARY)
clean: 
	rm -rf output && mkdir output 
.PHONY: build
build: go docker-build up
go:
	go mod tidy && go build -o output/$(GO_BINARY)
	cp index.html output

docker-build: build
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
up: 
	docker-compose up -d $(IMAGE_NAME)