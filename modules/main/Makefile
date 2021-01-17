SHELL := /usr/bin/env bash

MODULE_NAME = main
PUBLISH_PROJECT = hscasn
VERSION = 1.0.0
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
