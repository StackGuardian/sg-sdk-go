.PHONY: format apply-patch build cli test

# Makefile

# Format the Go SDK code
format:
	gofmt -w .
	goimports -w .

# Apply git patches in order
# Any new patches are to be added at the end of this block
apply-patch:
	git apply gitPatches/basePatch-workflowGroups.patch
	git apply gitPatches/basePatch-UnmashalJSON-Optional.patch
	git apply gitPatches/basePatch-optional-for-patched-policies.patch
	git apply gitPatches/basePatch-optional-for-patched-integration.patch
# Build target to format and apply patches in sequence
build: format apply-patch
# Build the sg command-line interface into bin/sg (VERSION defaults to the git describe output)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
cli:
	go build -ldflags "-X github.com/StackGuardian/sg-sdk-go/cli.Version=$(VERSION)" -o bin/sg ./cmd/sg

# Offline unit tests (SDK runtime + CLI)
test:
	go test ./internal/ ./core/ ./cli/
