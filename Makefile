BINARY_NAME := get-secret-value
VENDOR_DIR := ./vendor
BIN_PATH := bin
BUILD_PATH := main.go

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "${VENDOR_DIR}/*" | \
	xargs -I {} dirname {}  | \
	uniq)

all: build

ensure:
	dep ensure

build: ensure
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_PATH}/${BINARY_NAME}_linux_amd64 ${BUILD_PATH}
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_PATH}/${BINARY_NAME}_darwin_amd64 ${BUILD_PATH}

fmt: ${GO_SRC_DIRS}
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go fmt ; \
		popd > /dev/null ; \
	done;

.PHONY: all ensure build fmt
