.PHONY: test

install:
	go mod download
	go install golang.org/x/tools/cmd/stringer@latest

generate:
	go generate ./...

test:
	mkdir -p coverage
	go test -coverpkg=./... -coverprofile=./coverage/profile.cov ./...

cover: test
	go tool cover -func ./coverage/profile.cov

cover-html: test
	go tool cover -html ./coverage/profile.cov

build: generate build-golox

build-%: generate
	go build -o ./build/$* .

all: build test
