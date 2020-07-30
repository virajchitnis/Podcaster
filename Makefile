CGO=CGO_ENABLED=0
GOCMD=$(CGO) go
GOBUILD=$(GOCMD) build -ldflags '-s'
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u -a -ldflags '-s'
BUILD_DIR=bin
BINARY_NAME=$(BUILD_DIR)/Podcaster
BINARY_LINUX=$(BINARY_NAME)_linux
DOCKER_TAG=podcaster
DOCKER_CONTAINER_NAME=podcaster

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v

clean: 
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run: build
	./$(BINARY_NAME) -debug -config "examples/etc/podcaster/config.yaml"

deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) gopkg.in/yaml.v2

docker: build-linux-amd64
	docker build -t $(DOCKER_TAG) .

run-docker: docker
	docker run --rm -v "$(PWD)/examples/docker_config":/etc/podcaster -v "$(PWD)/examples/var/podcaster":/var/podcaster -p 8080:8080 --name $(DOCKER_CONTAINER_NAME) $(DOCKER_TAG)