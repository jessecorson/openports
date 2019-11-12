.phony: before_install rich_test

before_install: 
	go get -u github.com/kyoh86/richgo

rich_test:
	richgo test -v ./...

test: before_install rich_test

testf:
	go test -v ./...
