PACKAGES := $(shell go list ./...)

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## build: build a binary
.PHONY: build
build: test
	go build -o ./app -v

## autobuild: auto build when source files change
.PHONY: autobuild
autobuild:
	# curl -sf https://gobinaries.com/cespare/reflex | sh
	reflex -g '*.go' -- sh -c 'echo "\n\n\n\n\n\n" && make build'

## start: build and run local project
.PHONY: start
start: build
	clear
	@echo ""
	./app

## dockerbuild: build project into a container image
.PHONY: dockerbuild
dockerbuild: test
	docker-compose build

## login: login to ecr public
.PHONY: login
login:
	aws ecr-public get-login-password --region us-east-1 | \
		docker login --username AWS --password-stdin public.ecr.aws

## push: push container image to registry
.PHONY: push 
push: dockerbuild
	docker-compose push
