all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o subprocess main.o -L../../.libs/ -l$(LIBNAME) $(LDFLAGS)

clean:
	rm -f *.o subprocess

.PHONY: build clean
