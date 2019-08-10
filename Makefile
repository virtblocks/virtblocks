all: build

fmt:
	go fmt ./...

verify-fmt:
	hack/verify-gofmt.sh

vet:
	go vet ./pkg/... ./cmd/...

build:
	$(MAKE) -C capi build

clean:
	$(MAKE) -C capi clean

.PHONY: all fmt clean build vet verify-fmt
