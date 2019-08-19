build: libvirtblocks.la capi.lo

gobuild:
	go build -buildmode c-archive -o libvirtblocks.a capi.go

capi.lo: capi.c gobuild
	libtool --mode=compile --tag=CC gcc -c capi.c

libvirtblocks.la: libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../libvirtblocks.a .libs/libvirtblocks.a
	sed 's/@LIBTOOL_LA_NAME@/$@/g' <$< >$@

clean:
	rm -rf *.a *.la *.lo *.o .libs/ libvirtblocks.h

.PHONY: build gobuild clean
