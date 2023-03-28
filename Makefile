.PHONY: all
all: lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./main.go

.PHONY: build
run:
	go build .

.PHONY: test
run:
	go test -v ./...