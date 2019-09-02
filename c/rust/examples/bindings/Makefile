all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o bindings main.o -L../../.libs/ -l$(LIBNAME) $(LDFLAGS)

clean:
	rm -f *.o bindings

.PHONY: build clean
