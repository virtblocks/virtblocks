all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o command main.o -L../../.libs/ -l$(LIBNAME) $(LDFLAGS)

clean:
	rm -f *.o command

.PHONY: build clean
