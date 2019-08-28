all: build

build: libvirtblocks_c_go.la

# We need to build twice: first just the glue code, which also generates
# the header file; then also the wrappers, which themselves call to the
# glue code and as such can't be compiled withoug the generated header file
gobuild:
	go build -buildmode c-archive -o libvirtblocks_c_go.a glue.go
	go build -buildmode c-archive -o libvirtblocks_c_go.a

libvirtblocks_c_go.la: ../libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../$(@:.la=.a) .libs/$(@:.la=.a)
	sed \
	  -e 's/@LIBNAME@/$(@:.la=)/g' \
	  -e 's/@LDFLAGS@//g' \
	  <$< >$@

clean:
	rm -rf *.a *.la .libs/ libvirtblocks_c_go.h

test: gobuild
	go test

run-examples:
	@true

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet: gobuild
	go vet ./...

.PHONY: build gobuild clean test run-examples fmt verify-fmt vet
