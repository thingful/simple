# Makefile for building our simple Go server, and deploying to Docker hub

# Include env.mk file if it exists
-include env.mk

BUILDFLAGS = -a -installsuffix cgo
VERSION = $(shell git rev-parse HEAD)

SHELL := /bin/bash

.PHONY: help
help: ## Show this help message
	@echo 'usage: make [target] ...'
	@echo
	@echo 'targets:'
	@echo
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s : | sed -e 's/^/  /')"

.PHONY: compile
compile: ## compiles a static executable inside the container
	export CGO_ENABLED=0
	export GOOS=linux
	go build ${BUILDFLAGS} -o ${EXECUTABLE} .

.PHONY: test
test: ## runs tests inside the container
	go test -v .

.PHONY: build
build: ## builds our final container using docker-compose
	docker-compose build

.PHONY: push
push: build ## pushes the container to Docker hub
	docker tag ${USERNAME}/${EXECUTABLE}:latest thingful/simple:${VERSION}
	docker push ${USERNAME}/${EXECUTABLE}:latest
	docker push ${USERNAME}/${EXECUTABLE}:${VERSION}
