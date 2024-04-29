.PHONY: build

build:
	chmod +x ./scripts/swagger.sh
	./scripts/swagger.sh
	go build -o bin/go-template ./cmd

run:
	./bin/go-template

br:
	./scripts/swagger.sh
	go build -o bin/go-template ./cmd
	./bin/go-template	