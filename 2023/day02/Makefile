.PHONY: clean build test run

build:
	go build -o bin/day02 cmd/main.go

clean:
	rm -rf bin/*

test:
	go test -v ./part_one/

run:
	./bin/day02

default: build 
	
