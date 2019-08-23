all: build

build: rustbuild
	@true

rustbuild:
	$(MAKE) -C ../../c/rust build

clean:
	go clean

test: build
	@true

run: build
	for example in $$(ls examples/); do \
		go run examples/$$example/main.go || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet: build
	go vet ./...

.PHONY: build rustbuild clean test run fmt verify-fmt vet
