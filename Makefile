all: build

fmt:
	go fmt ./...

build:
	cd capi && \
	go build -buildmode c-archive -o libvirtblocks.a capi.go && \
	libtool --mode=compile gcc -c capi.c && \
	mkdir -p .libs/ && \
	ln -sf ../libvirtblocks.a .libs/libvirtblocks.a && \
	( \
	  echo "# libvirtblocks.la - a libtool library file" && \
	  echo "# Generated by libtool (except actually not) 0.0.1" && \
	  echo "dlname=''" && \
	  echo "library_names=''" && \
	  echo "old_library='libvirtblocks.a'" && \
	  echo "inherited_linker_flags=''" && \
	  echo "installed=no" && \
	  echo "shouldnotlink=no" \
	) >libvirtblocks.la

clean:
	cd capi && \
	rm -rf *.a *.la *.lo *.o .libs/ libvirtblocks.h

.PHONY: all fmt clean build