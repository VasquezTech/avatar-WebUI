# Variables
IMAGE_NAME := docker.io/mrvasquez96/avatar-ui
SERVICE_NAME := avatar-ui
DOCKERFILE := ./Dockerfile
GIT_COMMIT := $(shell git rev-parse HEAD)
VUE_FILES := ./static/public/
FRONTEND := index.html

OUTPUT_DIR := output/
GO_PWD := $(OUTPUT_DIR)go-avatar # Build
GO_BINARY:=go-avatar # Startup

# Output
clean:  
	rm -rf $(OUTPUT_DIR) && mkdir output

vue:
	cp ./$(FRONTEND) $(VUE_FILES)
	cd static && npm run build

vue-install:
	cd static && npm install && npm run build
	cp ./$(FRONTEND) $(VUE_FILES)

# Golang
go:
	export GIT_COMMIT=$(shell git rev-parse HEAD) && \
	go mod tidy && \
	go build -ldflags "-X main.gitCommit=$(shell git rev-parse HEAD)" -o $(GO_PWD)
	chmod +x $(GO_PWD)

# Docker
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) . 
	docker-compose up -d $(SERVICE_NAME)

re: clean vue-install go 
	cd $(OUTPUT_DIR) && ./$(GO_BINARY)
run: go
	cd $(OUTPUT_DIR) && ./$(GO_BINARY)
