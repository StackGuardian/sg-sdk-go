# First format the generated SDK code
gofmt -w . && goimports -w .

# Apply all the git patches in order
# Include any new git patches at the end of this section
git apply gitPatches/inital.patch



