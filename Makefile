#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=golang-echo-realworld-example-app
export LDFLAGS="-w -s"

all: build test

init:
	go build && sudo cp .env.example .env && sudo echo "DB_PASSWORD=asd@asd@" >> .env && sudo chmod 777 ./reports && ./reports &
build:
	go build -race  .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

############################################################
# Test
############################################################

test:
	go test -v -race ./...

container:
	docker build -t echo-realworld .

run-container:
	docker run --rm -it echo-realworld

.PHONY: build run build-static test container
