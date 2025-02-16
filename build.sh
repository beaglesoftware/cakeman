#!/bin/sh
set -e

echo "Deleting previous binaries..."
rm -rf dist/
echo "Building the project..."

GOOS=windows GOARCH=amd64 go build -o dist/windows/amd64/cman.exe
GOOS=darwin GOARCH=arm64 go build -o dist/macos/arm64/cman
GOOS=linux GOARCH=amd64 go build -o dist/linux/amd64/cman
echo "Cakeman built \033[0;92msuccessfully\033[0;0m!"
