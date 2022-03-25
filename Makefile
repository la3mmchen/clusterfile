#!make
include properties.env
export $(shell sed 's/=.*//' properties.env)
GIT_COMMIT := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$$PWD")

.DEFAULT_GOAL := default

.PHONY: default
default: build run tests

#
# *** build steps *** 
# 
.PHONY: build
build: go-mod build-executable

go-mod:
	@go mod vendor
	@go mod verify

.PHONY: build-executable
build-executable: 
	@rm -f ${EXECUTABLE}
	@go build -o ${EXECUTABLE} -ldflags "-X main.AppVersion=${GIT_COMMIT}" .

#
# *** example runs ****
#
.PHONY: run
run: run-help run-func

run-help:
	@./${EXECUTABLE} --help
	@echo "\n____________________________ \n"

run-func:
	@./${EXECUTABLE} preflight --offline
	@echo "\n____________________________"
	@./${EXECUTABLE} status --offline
	@echo "\n____________________________\n"

#
# *** tests ****
#
.PHONY: tests
tests: unit-tests app-tests

.PHONY: unit-tests
unit-tests:
	@go test -cover -failfast -short "./.../types"
	@echo "\n____________________________"
	@go test -cover -failfast -short "./.../app"
	@echo "\n____________________________"

.PHONY: app-tests
app-tests:
	@go test -cover -parallel 1 -failfast -short "."
	@echo "\n____________________________"

.PHONY: add-to-path
add-to-path: build-executable
	@cp ./${EXECUTABLE} /usr/local/bin/${BINARY_NAME}