.PHONY: test

generate:
	go generate ./...

test:
	mkdir -p coverage
	go test -coverpkg=./... -coverprofile=./coverage/profile.cov ./...

cover: test
	go tool cover -func ./coverage/profile.cov

cover-html: test
	go tool cover -html ./coverage/profile.cov

build: generate
	go build -o ./build/golox .

all: build test
