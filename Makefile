CGO=CGO_ENABLED=0
GOCMD=$(CGO) go
GOBUILD=$(GOCMD) build -ldflags '-s'
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u -a -ldflags '-s'
BUILD_DIR=bin
BINARY_NAME=$(BUILD_DIR)/Podcaster

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run: build
	./$(BINARY_NAME) -debug -config "example_config.yaml"

deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) gopkg.in/yaml.v2