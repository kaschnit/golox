# golox

A lox interpreter written in go.

## Project Root Structure
* `cmd` - driver program for the binary
* `configs` - configuration files
* `pkg` - implementation and unit tests of core features

## Build tasks
* `make generate` - run codegen
* `make build` - build the golox binary to `/build/golox`
* `make test` - run tests
* `make cover` - run tests and display code coverage
* `make cover-html` - run tests and display code coverage visually in a web browser


## Usage
Run `golox --help` to see usage.