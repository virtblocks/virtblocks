all: build

# Some versions of cargo are known to produce library archives that
# can't be linked properly. Luckily there's a relatively straightforward
# workaround available.
#
# https://github.com/rust-lang/rust/issues/58277
build:
	cargo build
	case $$(cargo version) in \
	  *1.34*|*1.35*) \
	    ar d target/debug/libvirtblocks.a clzsi2.o; \
	    ar d target/debug/.libs/libvirtblocks.a clzsi2.o; \
	    ;; \
	esac

clean:
	cargo clean
	rm -rf target/

.PHONY: build clean
