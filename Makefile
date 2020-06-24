IMAGE ?= wjwinner/gitlab-wechat-hook
TAG ?= v0.2.0
OS   := linux
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(ARCH) go build -a -ldflags '-extldflags "-static"' -o bin/hook  cmd/hook/hook.go
