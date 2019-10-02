linux:
	GOOS=linux GOARCH=amd64 go build -o bin/export cmd/export/main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/export cmd/export/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/export.exe cmd/export/main.go

build: linux
