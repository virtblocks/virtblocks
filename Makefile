all: build

virtblocks-language:
	@if ! test "$(VIRTBLOCKS_LANGUAGE)"; then \
	  echo "error: No language selected" 2>&1; \
	  echo "Use 'make VIRTBLOCKS_LANGUAGE=(rust|golang)' or" 2>&1; \
	  echo "run 'export VIRTBLOCKS_LANGUAGE=(rust|golang)' beforehand" 2>&1; \
	  exit 1; \
	fi

fmt: virtblocks-language fmt-$(VIRTBLOCKS_LANGUAGE)
verify-fmt: virtblocks-language verify-fmt-$(VIRTBLOCKS_LANGUAGE)
vet: virtblocks-language vet-$(VIRTBLOCKS_LANGUAGE)
build: virtblocks-language build-$(VIRTBLOCKS_LANGUAGE)
clean: virtblocks-language clean-$(VIRTBLOCKS_LANGUAGE)

.PHONY: all virtblocks-language fmt verify-fmt vet build clean

fmt-golang:
	go fmt ./...

verify-fmt-golang:
	hack/verify-gofmt.sh

vet-golang:
	go vet ./pkg/... ./cmd/...

build-golang:
	$(MAKE) -C capi build

clean-golang:
	$(MAKE) -C capi clean

.PHONY: fmt-golang verify-fmt-golang vet-golang build-golang clean-golang

fmt-rust:
	cargo fmt

verify-fmt-rust:
	cargo fmt -- --check

vet-rust:
	cargo clippy

# Some versions of cargo are known to produce library archives that
# can't be linked properly. Luckily there's a relatively straightforward
# workaround available.
#
# https://github.com/rust-lang/rust/issues/58277
build-rust:
	cargo build
	case $$(cargo version) in \
	  *1.34*|*1.35*) \
	    ar d target/debug/libvirtblocks.a clzsi2.o; \
	    ar d target/debug/.libs/libvirtblocks.a clzsi2.o; \
	    ;; \
	esac

clean-rust:
	cargo clean

.PHONY: fmt-rust verify-fmt-rust vet-rust build-rust clean-rust
