echo "Deleting previous binaries..."
rm -rf dist/
echo "Building the project..."

GOOS=windows GOARCH=amd64 go build -o dist/windows/amd64/cman.exe
GOOS=darwin GOARCH=arm64 go build -o dist/darwin/arm64/cman
GOOS=linux GOARCH=amd64 go build -o dist/linux/amd64/cman