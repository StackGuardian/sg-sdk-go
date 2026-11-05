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

# Check the code and its dependencies for known Go vulnerabilities
# (installs govulncheck pinned to v1.9.0)
govulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@v1.9.0
	govulncheck ./...

# Run the offline unit tests (internal/ and core/). The integration tests in
# tests/ and the per-resource client_test.go files hit the live StackGuardian
# API and need SG_API_TOKEN / SG_BASE_URL, so they are not part of this target.
test:
	go test ./internal/... ./core/...