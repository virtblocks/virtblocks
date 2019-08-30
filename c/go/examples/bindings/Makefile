all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) $(LDFLAGS) -o bindings main.o ../../.libs/$(LIBNAME).a

clean:
	rm -f *.o bindings

.PHONY: build clean
