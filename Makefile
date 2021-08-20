#!make
include properties.env
export $(shell sed 's/=.*//' properties.env)
GIT_COMMIT := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$$PWD")

.DEFAULT_GOAL := default

tests: unit-tests app-tests

.PHONY: default
default: build run-help run-func

.PHONY: build
build: 
	@rm -f ${EXECUTABLE}
	@go build -o ${EXECUTABLE} -ldflags "-X main.AppVersion=${GIT_COMMIT}" .

run:
	@./${EXECUTABLE}

run-help:
	@./${EXECUTABLE} --help
	@echo "\n____________________________ \n"

run-func:
	@./${EXECUTABLE} preflight --offline
	@echo "\n____________________________"
	@./${EXECUTABLE} status --offline
	@echo "\n____________________________\n"

.PHONY: unit-tests
unit-tests:
	@go test -cover -failfast -short "./.../types"
	@echo "\n____________________________"
	@go test -cover -failfast -short "./.../app"
	@echo "\n____________________________"

.PHONY: app-tests
app-tests:
	@go test -cover -failfast -short "."
	@echo "\n____________________________"
