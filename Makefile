# Variables
IMAGE_NAME := avatarui
DOCKERFILE := ./Dockerfile
GO_BINARY := go-avatar
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

.PHONY: re
re: build docker-build up
build:
	go mod tidy && go build -o $(GO_BINARY)

docker-build: build
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
up: 
	docker-compose up -d $(IMAGE_NAME)