.PHONY: install
install:
	go mod download
	go install golang.org/x/tools/cmd/stringer@latest
	go get golang.org/x/tools/cmd/stringer@latest

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	mkdir -p coverage
	go test -coverpkg=./... -coverprofile=./coverage/profile.cov ./...

.PHONY: test-in-docker
test-in-docker:
	docker build -t golox:test --target=builder .
	docker run -it golox:test make test

.PHONY: cover
cover: test
	go tool cover -func ./coverage/profile.cov

.PHONY: cover-html
cover-html: test
	go tool cover -html ./coverage/profile.cov

.PHONY: build
build: generate build-golox

build-%: generate
	go build -o ./build/$* .

.PHONY: all
all: build test
