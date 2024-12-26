GOOS=windows GOARCH=amd64 go build -o dist/windows/amd64/cakeman.exe
GOOS=darwin GOARCH=arm64 go build -o dist/darwin/arm64/cakeman
GOOS=linux GOARCH=amd64 go build -o dist/linux/amd64/cakeman