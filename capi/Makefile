all: stamp

stamp:
	go build -buildmode c-archive -o libvirtblocks.a capi.go
	cc -c -o capi.o capi.c
	touch stamp

clean:
	rm -f *.o libvirtblocks.* stamp

.PHONY: all clean
