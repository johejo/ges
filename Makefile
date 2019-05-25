GO=go
TARGET_PATH=github.com/johejo/ges/cmd/ges
TARGET_NAME=ges
OUTPUT_DIR=./out

IMAGE_NAME=$(TARGET_NAME)
IMAGE_TAG=latest

OPENAPI_GENERATOR=openapi-generator generate
OPENAPI_SPEC=./docs/spec/openapi.yaml
GENERATED_DIR=./generated

.PHONY: default all build binary clean test e2e check prepare

.PHONY: wire mockgen tidy

.PHONY: docker-image docker-build

.PHONY: openapi-go openapi-go-server openapi-html

.PHONY: doc

default: binary

all: prepare binary docker-image doc

clean:
	make openapi-clean
	rm -rf $(OUTPUT_DIR)

prepare: tidy wire mockgen test

binary: tidy wire
	make build

build:
	$(GO) build -o $(OUTPUT_DIR)/$(TARGET_NAME) $(TARGET_PATH)

test:
	$(GO) test ./... -v -cover

e2e:
	echo 'e2e'

check:
	./scripts/check_tools.sh


tidy:
	$(GO) mod tidy

wire:
	wire $(TARGET_PATH)

mockgen:
	echo 'mockgen'


docker-image: prepare
	make docker-build

docker-build:
	docker build -t $(IMAGE_NAME):$(TAG) .


openapi-clean:
	rm -rf $(GENERATED_DIR)

openapi-go:
	$(OPENAPI_GENERATOR) -i $(OPENAPI_SPEC) -g go -o $(GENERATED_DIR)/go

openapi-go-server:
	$(OPENAPI_GENERATOR) -i $(OPENAPI_SPEC) -g go-server -o $(GENERATED_DIR)/go-server

openapi-html:
	$(OPENAPI_GENERATOR) -i $(OPENAPI_SPEC) -g html -o $(GENERATED_DIR)/html

doc: openapi-html
