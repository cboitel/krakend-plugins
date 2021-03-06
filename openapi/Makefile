# This Makefile is a simple example that demonstrates usual steps to build a binary that can be run in the same
# architecture that was compiled in. The "ldflags" in the build assure that any needed dependency is included in the
# binary and no external dependencies are needed to run the service.
BIN_NAME := krakend_openapi.so
VERSION := 0.1.0
PKGNAME := krakend
LICENSE := Apache 2.0
VENDOR=
URL := http://github.com/cboitel/krakend-plugins/openapi
RELEASE := 0
DESC := Open API Client Plugin for krakend API gateway
MAINTAINER := Christian Boitel <christian.boitel@gmail.com>
GOLANG_VERSION := 1.13
GO_SRC=$(wildcard *.go)
NOW=$(shell date +%Y%m%d.%H%M%S%:::z) 

.PHONY: clean build all test

all: test

build: ${BIN_NAME}

${BIN_NAME}: ${GO_SRC}
	@echo "Building the plugin..."
	go get .
	go build -buildmode=plugin -ldflags="-X main.PluginVersion=${VERSION} -X main.PluginBuildDate=${NOW}" -o ${BIN_NAME} .
	@echo "You can now use ./${BIN_NAME}"

test: build
	go test -v ./tests

clean:
	rm -rf ${BIN_NAME}
	go clean
