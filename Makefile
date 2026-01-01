.PHONY: test clean all

NAME = mac
OS_ARCH ?= darwin_arm64
TESTARGS ?= -tags=asdf,brew,service

build:
	goreleaser build --snapshot --clean

generate:
	go generate

install: build
	rm -rf /tmp/tfproviders/
	mkdir -p /tmp/tfproviders/
	mv dist/terraform-provider-${NAME}_${OS_ARCH}_*/* /tmp/tfproviders/

test:
	go test $(TESTARGS) -race -parallel=4 ./...

testacc:
	TF_ACC=1 go test $(TESTARGS) -race -parallel=4 ./...
