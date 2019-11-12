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
	env GOOS=linux GOARCH=arm go build -o bin/linux/openports/openports && pushd bin/linux && tar -cvzf openports-linux.tar.gz openports ; popd
build-mac:
	env GOOS=darwin go build -o bin/darwin/openports/openports && pushd bin/darwin && tar -cvzf openports-darwin.tar.gz openports ; popd
build-windows:
	env GOOS=windows GOARCH=386 go build -o bin/windows/openports/openports.exe && pushd bin/windows && tar -cvzf openports-windows.tar.gz openports ; popd

build-binaries: build-linux build-mac build-windows
