# Variables
IMAGE_NAME := mrvasquez96/avatar-ui:latest
DOCKERFILE := ./Dockerfile
GO_BINARY := go-avatar
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
 
.PHONY: build
build clean vue go: 

# Output
clean: 
	rm -rf output && mkdir output
# Golang
go:
	go mod tidy && go build -o output/$(GO_BINARY)
	chmod +x  output/$(GO_BINARY)

vue:
	cd static && npm run build
# cp index.html output

# Docker
docker-build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
docker-up: 
	docker-compose up -d $(IMAGE_NAME)


docker: build

re: clean go 
	cd output && ./go-avatar