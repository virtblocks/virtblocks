all: build

build: libvirtblocks_c_go.la

gobuild:
	go build -buildmode c-archive -o libvirtblocks_c_go.a capi.go

capi.o: capi.c gobuild
	gcc -fpic -DPIC -c capi.c -o capi.o

# If something weird happens later, just try running `ranlib` after `ar`
libvirtblocks_c_go.la: ../libtool.la.in gobuild capi.o
	mkdir -p .libs/
	ln -sf ../$(@:.la=.a) .libs/$(@:.la=.a)
	ar r $(@:.la=.a) capi.o
	sed \
	  -e 's/@LIBNAME@/$(@:.la=)/g' \
	  -e 's/@LDFLAGS@//g' \
	  <$< >$@

clean:
	rm -rf *.a *.la *.lo *.o .libs/ libvirtblocks_c_go.h

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
