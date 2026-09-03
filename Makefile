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