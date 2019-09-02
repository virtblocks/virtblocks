all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) -o playground main.o ../../.libs/$(LIBNAME).a $(LDFLAGS)

clean:
	rm -f *.o playground

.PHONY: build clean
