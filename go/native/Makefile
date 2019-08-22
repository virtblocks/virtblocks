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

.PHONY: build clean test run
