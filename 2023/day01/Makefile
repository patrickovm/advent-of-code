.PHONY: clean build test run

build:
	go build -o bin/day01 cmd/main.go

clean:
	rm -rf bin/*

test:
	go test -v ./part_one/

run:
	./bin/day01

default: build 
	
