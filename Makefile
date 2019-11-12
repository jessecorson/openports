.phony: before_install rich_test

before_install: 
	go get -u github.com/kyoh86/richgo

rich_test:
	richgo test -v ./...

test: before_install rich_test

testf:
	go test -v ./...

install:
	go install

.phone: build-linux build-mac build-windows

build-linux:
	env GOOS=linux GOARCH=arm go build -o bin/linux/openports
build-mac:
	env GOOS=darwin go build -o bin/darwin/openports
build-windows:
	env GOOS=windows GOARCH=386 go build -o bin/windows/openports.exe

build-binaries: build-linux build-mac build-windows
