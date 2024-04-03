# Variables
IMAGE_NAME := docker.io/mrvasquez96/avatar-ui
SERVICE_NAME := avatar-ui
DOCKERFILE := ./Dockerfile
GIT_COMMIT := $(shell git rev-parse HEAD)

OUTPUT_DIR := output/
GO_PWD := $(OUTPUT_DIR)serve/ # Build
GO_BINARY:=go-avatar # Startup

# Output
clean:  
	rm -rf $(OUTPUT_DIR) 

vue:
	cd static && \
	npm run build

vue-install:
	cd static && \
	npm install && \
	npm run build

# Golang
go-build:
	mkdir -p $(GO_PWD)
	go mod tidy
#	go build -ldflags "-X main.gitCommit=$(shell git rev-parse HEAD)" -o $(GO_PWD)$(GO_BINARY)
	go build -o ./$(GO_PWD)$(GO_BINARY) && \
	chmod +x $(GO_PWD)$(GO_BINARY)

# Docker
docker-build: clean
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) . 
	docker-compose up -d $(SERVICE_NAME)

build-d: clean
	docker build -t $(IMAGE_NAME):dev -f $(DOCKERFILE) . 

docker-run:
	docker run -p 127.0.0.1:8055:8055 mrvasquez96/avatar-ui:latest

re: clean vue-install go-build
	cd $(Go) && ./serve/$(GO_BINARY)

dev: clean go-build
	chown -R $(USER) .
	cd $(OUTPUT_DIR) && \
	chmod +x ./serve/$(GO_BINARY) && \
	./serve/$(GO_BINARY) dev
	
