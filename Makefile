BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}{{end}}' ./...)

VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null)
REVISION ?= $(shell git rev-parse --short HEAD 2>/dev/null)

GO_LDFLAGS := -X main.version=$(VERSION)
GO_LDFLAGS += -X main.revision=$(REVISION)

bin/dotenveditor: $(BUILD_FILES)
	@go build -trimpath -ldflags "$(GO_LDFLAGS)" -o "$@" ./cmd/dotenveditor

build: bin/dotenveditor
