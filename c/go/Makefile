export LIBNAME = virtblocks_c_go
export LDFLAGS = -pthread

all: build

build: lib$(LIBNAME).la

# We need to build twice: first just the glue code, which also generates
# the header file; then also the wrappers, which themselves call to the
# glue code and as such can't be compiled withoug the generated header file
gobuild:
	go build -buildmode c-archive -o lib$(LIBNAME).a glue.go
	go build -buildmode c-archive -o lib$(LIBNAME).a

lib$(LIBNAME).la: ../libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../lib$(LIBNAME).a .libs/lib$(LIBNAME).a
	sed \
	  -e 's/@LIBNAME@/$(LIBNAME)/g' \
	  -e 's/@LDFLAGS@/$(LDFLAGS)/g' \
	  <$< >$@

clean:
	rm -rf *.a *.la .libs/ lib$(LIBNAME).h

test: gobuild
	go test

run-examples: build
	for example in $$(ls examples/); do \
		$(MAKE) -C examples/$$example || exit 1; \
		examples/$$example/$$example || exit 1; \
		$(MAKE) -C examples/$$example clean || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet: gobuild
	go vet ./...

.PHONY: build gobuild clean test run-examples fmt verify-fmt vet
