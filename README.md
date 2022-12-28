# golox

An interpreter for the `Lox` programming language, written in go.

## Project Root Structure

- `cmd/` - driver program for the binary
- `configs/` - configuration files
- `pkg/` - implementation and unit tests of core features
- `test/` - contains test data and end-to-end tests

## Build Tasks

The following tasks are defined in the `Makefile` in the project root.

- `make install` - install all dependencies
- `make format` - format the code
- `make generate` - run codegen
- `make build` - build the golox binary to `build/golox`
- `make build-<binary-name>` - build the golox binary to `build/<binary-name>`
- `make unittest` - run unit tests only
- `make e2etest` - run end to end tests only
- `make test` - run all tests
- `make <task>-in-docker` - run the `task` in docker (e.g., `make test-in-docker`)
- `make cover` - run tests and display code coverage
- `make cover-html` - run tests and display code coverage visually in a web browser
- `make run` - run the interpreter

## Usage

Run `golox --help` to see usage.
