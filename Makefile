.PHONY: install
install:
	go mod download
	go install golang.org/x/tools/cmd/stringer@latest
	go get golang.org/x/tools/cmd/stringer@latest

.PHONY: format
format:
	go fmt ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: unittest
unittest:
	mkdir -p coverage
	go test -coverpkg=./... -coverprofile=./coverage/profile.cov $(shell go list ./... | grep -v /test/)

.PHONY: e2etest
e2etest:
	go test ./test/...

.PHONY: test
test:
	mkdir -p coverage
	go test -coverpkg=./... -coverprofile=./coverage/profile.cov ./...

.PHONY: %-in-docker
%-in-docker: DOCKER_TARGET=builder
%-in-docker:
	docker build -t golox:test --target=$(DOCKER_TARGET) .
	docker run -it golox:test make $*

.PHONY: cover
cover: test
	go tool cover -func ./coverage/profile.cov

.PHONY: cover-html
cover-html: test
	go tool cover -html ./coverage/profile.cov

.PHONY: build
build: generate build-golox

.PHONY: build-%
build-%: generate
	go build -o ./build/$* .

.PHONY: run
run: ARGS = interpreter -i
run:
	go run . $(ARGS)

.PHONY: all
all: build test
