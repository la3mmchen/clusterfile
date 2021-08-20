#!make
include properties.env
export $(shell sed 's/=.*//' properties.env)
GIT_COMMIT := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$$PWD")

.DEFAULT_GOAL := default

test:
	@go test -v

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
	@echo "\n____________________________"
