all: build

build:
	@true

clean:
	rm -f tests/functest.bin

test:
	@true

run-examples:
	for example in $$(ls examples/); do \
	( \
		cd examples/$$example || exit 1; \
		$(MAKE) run-examples || exit 1; \
	) || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../scripts/verify-gofmt.sh

vet:
	go vet ./...

functest:
	go test ./tests/... -c -o ./tests/functest.bin
	./tests/functest.bin

.PHONY: build clean test run-examples fmt verify-fmt vet functest
