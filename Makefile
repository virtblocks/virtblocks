# Top-level targets meant to be invoked by users

all: build

build: build-subdirs
clean: clean-subdirs
test: test-subdirs
run-examples: run-examples-subdirs
fmt: fmt-subdirs
verify-fmt: verify-fmt-subdirs
vet: vet-subdirs
functest: functest-subdirs

.PHONY: all build clean test run-examples fmt verify-fmt vet functest

# These are the regular targets, the ones which operate on all code

SUBDIRS = \
	go \
	c \
	$(NULL)

%-subdirs:
	for d in $(SUBDIRS); do \
		$(MAKE) -C $$d $(subst -subdirs,,$@) || exit 1; \
	done
