SHELL := /bin/bash

setup_ci: install-stringer

setup: install-stringer install-pre-commit-hooks
	go mod vendor

install-stringer:
	@type stringer >/dev/null 2>&1 || go install golang.org/x/tools/cmd/stringer@latest

install-pre-commit-hooks:
	pre-commit install --install-hooks

run-pre-commit:
	pre-commit run -a

test:
	go test -v -race ./...

coverage:
	go test -v -race -cover -covermode=atomic ./...

test-with-update-snapshots:
	UPDATE_SNAPSHOTS=true go test -v -race ./...
