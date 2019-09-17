EXAMPLE = devices
LIBNAME = virtblocks_c_go
LDFLAGS = -pthread

all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o $(EXAMPLE) main.o -L../../.libs/ -l$(LIBNAME) $(LDFLAGS)

clean:
	rm -f *.o $(EXAMPLE)

run-examples:
	$(MAKE) build
	./$(EXAMPLE)
	$(MAKE) clean

.PHONY: build clean run-examples
