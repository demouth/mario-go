all: build-win build-mac build-linux

clean:
	rm -f release/mario-go
	rm -f release/mario-go.exe 
	rm -f release/windows_386_mario-go.zip
	rm -f release/darwin_amd64_mario-go.tar.gz
	rm -f release/linux_amd64_mario-go.tar.gz

build-win: clean
	GOOS=windows GOARCH=386 go build -o release/mario-go.exe cmd/main.go
	zip -r -q -T release/windows_386_mario-go.zip release/mario-go.exe

build-mac: clean
	GOOS=darwin GOARCH=amd64 go build -o release/mario-go cmd/main.go
	tar -czf release/darwin_amd64_mario-go.tar.gz release/mario-go

build-linux: clean
	GOOS=linux GOARCH=amd64 go build -o release/mario-go cmd/main.go
	tar -czf release/linux_amd64_mario-go.tar.gz release/mario-go

run-linux: build-linux
	docker build -t demouth/mario-go-centos7 .
	docker run --rm -v $(shell pwd):/mario-go/ -ti demouth/mario-go-centos7 sh -c 'cd /mario-go/release; /bin/bash'


