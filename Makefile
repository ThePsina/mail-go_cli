GOLINTFLAGS ?=
GOBIN        = $(shell go env GOPATH)/bin

.PHONY: all
all: build

.PHONY: build
build:
ifndef TARGET
	@echo 'build target is not defined'
else
	go build $(GOTAGS) \
		-ldflags '$(LDFLAGS)' \
		-o bin/$(TARGET) \
		./cmd/$(TARGET)
endif

.PHONY: test
test:
	sh test.sh

.PHONY: fmt
fmt:
	gofmt -s -w ./cmd

.PHONY: lint
lint:
	golangci-lint --color=always run ./... $(GOLINTFLAGS)