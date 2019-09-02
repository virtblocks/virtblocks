all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o bindings main.o ../../.libs/lib$(LIBNAME).a $(LDFLAGS)

clean:
	rm -f *.o bindings

.PHONY: build clean
