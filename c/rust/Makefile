all: build

build: libvirtblocks_c_rust.la

# Some versions of cargo are known to produce library archives that
# can't be linked properly. Luckily there's a relatively straightforward
# workaround available.
#
# https://github.com/rust-lang/rust/issues/58277
rustbuild:
	cargo build
	case $$(cargo version) in \
	  *1.34*|*1.35*) \
	    ar d target/debug/libvirtblocks_c_rust.a clzsi2.o; \
	    ;; \
	esac

libvirtblocks_c_rust.la: ../libtool.la.in rustbuild
	mkdir -p .libs/
	ln -sf ../target/debug/$(@:.la=.a) .libs/$(@:.la=.a)
	sed \
	  -e 's/@LIBNAME@/$(@:.la=)/g' \
	  -e 's/@LDFLAGS@/-pthread -lm -ldl/g' \
	  <$< >$@

clean:
	cargo clean

test:
	cargo test

run-examples:
	@true

fmt:
	cargo fmt

verify-fmt:
	cargo fmt -- --check

vet:
	cargo clippy

.PHONY: build rustbuild clean test run-examples fmt verify-fmt vet
