SHELL := /usr/bin/env bash

.PHONY: test
.PHONY: build
.PHONY: publish 

test:
	@(cd modules; ../scripts/for_all.sh "make test")
	@(cd lib; ../scripts/for_all.sh "../../scripts/for_all.sh \"make test\"")

build:
	@(cd modules; ../scripts/for_all.sh "make build")
	@(cd lib; ../scripts/for_all.sh "../../scripts/for_all.sh \"make build\"")

publish:
	@(cd modules; ../scripts/for_all.sh "make publish")
	@(cd lib; ../scripts/for_all.sh "../../scripts/for_all.sh \"make publish\"")
