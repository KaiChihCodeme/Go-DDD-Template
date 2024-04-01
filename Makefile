.PHONY: build

build:
	go build -o bin/go-template ./cmd

run:
	./bin/go-template

br:
	go build -o bin/go-template ./cmd
	./bin/go-template	