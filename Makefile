# Go CLI — Makefile

BINARY := bin/go-cli-template

.PHONY: run build test vet lint docker clean install

run: build
	./$(BINARY) example

build:
	go build -o $(BINARY) .

test:
	go test -race -cover ./...

vet:
	go vet ./...

lint: vet
	golangci-lint run ./...

install:
	go install .

docker:
	docker build -t go-cli-template .

clean:
	rm -rf bin/