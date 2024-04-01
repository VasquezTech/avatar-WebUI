# Variables
IMAGE_NAME := mrvasquez96/avatar-ui
DOCKERFILE := ./Dockerfile
GO_BINARY := go-avatar
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
 
.PHONY: clean vue go
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
	docker build -t mrvasquez96/avatar-ui -f $(DOCKERFILE) .
docker-up: 
	docker-compose up -d $(IMAGE_NAME)



docker: clean vue go
# cp -r /tmp/build/output /app

re: clean go 
	cd output && ./go-avatar