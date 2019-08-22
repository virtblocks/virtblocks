all: build

build: libvirtblocks_c_go.la capi.lo

gobuild:
	go build -buildmode c-archive -o libvirtblocks_c_go.a capi.go

capi.lo: capi.c gobuild
	libtool --mode=compile --tag=CC gcc -c capi.c

libvirtblocks_c_go.la: libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../$(@:.la=.a) .libs/$(@:.la=.a)
	sed 's/@LIBNAME@/$(@:.la=)/g' <$< >$@

clean:
	rm -rf *.a *.la *.lo *.o .libs/ libvirtblocks_c_go.h

test: gobuild
	go test

run:
	@true

.PHONY: build gobuild clean test run
