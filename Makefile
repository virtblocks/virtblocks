all: build

fmt:
	go fmt ./...

build:
	cd capi && \
	go build -buildmode c-archive -o libvirtblocks.a capi.go && \
	cc -c -o capi.o capi.c

clean:
	cd capi && \
	rm -f *.o libvirtblocks.* stamp

.PHONY: all clean build
