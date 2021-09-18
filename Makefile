GOVERSION=1.17
GOCMD=go${GOVERSION}

compile: clean build

.PHONY: clean
clean:
	${GOCMD} clean

.PHONY: build
build:
	${GOCMD} build -race ./...

.PHONY: lint
lint:
	@golangci-lint run --timeout 5m0s

.PHONY: test
test:
	${GOCMD} test -race -cover ./...

.PHONY: tidy
tidy:
	${GOCMD} mod tidy

.PHONY: vendor
vendor:
	${GOCMD} mod vendor

all: clean lint build test
