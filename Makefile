# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=swapi-cli
MAIN_PATH=cmd/swapicli
    
all: test build
build: 
	cd $(MAIN_PATH);$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	cd $(MAIN_PATH);rm -f $(BINARY_NAME)