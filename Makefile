.PHONY: all
all: lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./main.go

.PHONY: build
build:
	go build .

.PHONY: test
test:
	go test -v ./...

.PHONY: gci
gci:
	gci write --skip-generated -s standard -s default .