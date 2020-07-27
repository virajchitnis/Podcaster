GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_DIR=bin
BINARY_NAME=$(BUILD_DIR)/Podcaster

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run: build
	./$(BINARY_NAME) -debug

deps:
	$(GOGET) -u github.com/gin-gonic/gin