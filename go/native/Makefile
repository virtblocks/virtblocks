all: build

build:
	@true

clean:
	@true

test:
	@true

run:
	for example in $$(ls examples/); do \
		go run examples/$$example/main.go || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet:
	go vet ./...

.PHONY: build clean test run fmt verify-fmt vet
