SHELL := /usr/bin/env bash

MODULE_NAME = users
PUBLISH_PROJECT = hscasn
VERSION = $(shell date +"%M%H%m%d%y")
IMAGE_TAG = gcr.io/${PUBLISH_PROJECT}/${MODULE_NAME}:${VERSION}

.PHONY: test
.PHONY: build
.PHONY: publish 

test:
	@go test ./...

build:
	@docker build \
                -f ../../shared_dockerfiles/GoAlpine.Dockerfile \
                -t "${IMAGE_TAG}" \
                .

publish:
	@make build
	@docker push "${IMAGE_TAG}"
