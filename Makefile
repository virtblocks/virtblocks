# Top-level targets meant to be invoked by users

all: build

build: build-subdirs-$(VIRTBLOCKS_LANGUAGE)
clean: clean-subdirs-$(VIRTBLOCKS_LANGUAGE)
test: test-subdirs-$(VIRTBLOCKS_LANGUAGE)
run-examples: run-examples-subdirs-$(VIRTBLOCKS_LANGUAGE)
fmt: fmt-subdirs-$(VIRTBLOCKS_LANGUAGE)
verify-fmt: verify-fmt-subdirs-$(VIRTBLOCKS_LANGUAGE)
vet: vet-subdirs-$(VIRTBLOCKS_LANGUAGE)

.PHONY: all build clean test run-examples fmt verify-fmt vet

# These are the regular targets, the ones which operate on all code

SUBDIRS = \
	rust/native \
	go/native \
	c/rust \
	c/go \
	go/rust \
	$(NULL)

%-subdirs-:
	for d in $(SUBDIRS); do \
		$(MAKE) -C $$d $(subst -subdirs-,,$@) || exit 1; \
	done

# These only operate on the Go code that can be compiled without having
# a Rust compiler installed, and are kept around for compatibility with
# the existing CI setup. Once we update that, we'll be able to get rid
# of them.

SUBDIRS_GOLANG = \
	go/native \
	c/go \
	$(NULL)

%-subdirs-golang:
	for d in $(SUBDIRS_GOLANG); do \
		$(MAKE) -C $$d $(subst -subdirs-golang,,$@) || exit 1; \
	done
