all: build

build: rustbuild
	@true

rustbuild:
	$(MAKE) -C ../../c/rust build

clean:
	@true

test: build
	@true

run-examples: build
	for example in $$(ls examples/); do \
		go run examples/$$example/main.go || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet: build
	go vet ./...

.PHONY: build rustbuild clean test run-examples fmt verify-fmt vet
