EXAMPLE = playground

all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o $(EXAMPLE) main.o -L../../.libs/ -l$(LIBNAME) $(LDFLAGS)

clean:
	rm -f *.o $(EXAMPLE)

.PHONY: build clean
