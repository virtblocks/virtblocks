all: build

build:
	cargo build

clean:
	cargo clean

test:
	cargo test

run-examples:
	for example in $$(ls examples/); do \
		cargo run --example $$example || exit 1; \
	done

fmt:
	cargo fmt

verify-fmt:
	cargo fmt -- --check

vet:
	cargo clippy

.PHONY: build clean test run-examples fmt verify-fmt vet
