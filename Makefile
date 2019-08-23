# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=swapi-cli
DOCKER_NAME=swapi-cli
DOCKER_VERSION=1.0.0
MAIN_PATH=cmd/swapicli
    
all: test build docker
build: 
	cd $(MAIN_PATH);$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	cd $(MAIN_PATH);rm -f $(BINARY_NAME)

docker:
	docker build -t $(DOCKER_NAME):$(DOCKER_VERSION) .
docker-run: 
	docker run $(DOCKER_NAME):$(DOCKER_VERSION)