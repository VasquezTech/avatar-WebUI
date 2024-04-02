# Variables
IMAGE_NAME := avatar-ui:v0.0.2
DOCKERFILE := ./Dockerfile
DOCKER_BUILD := docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

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
	go mod tidy && go build -o $(GO_PWD)
	chmod +x $(GO_PWD)
# cp index.html output

# Docker
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) . 
	docker-compose up -d $(IMAGE_NAME)

re: clean vue-install go 
	cd $(OUTPUT_DIR) && ./$(GO_BINARY)
run: go
	cd $(OUTPUT_DIR) && ./$(GO_BINARY)