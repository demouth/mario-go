all: build-win build-mac build-linux



build-win:
	GOOS=windows GOARCH=386 go build -o release/windows_386_mario-go.exe cmd/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o release/mac_amd64_mario-go cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o release/linux_amd64_mario-go cmd/main.go



