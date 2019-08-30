all: build

build:
	$(CC) -I../../ -o main.o -c main.c
	$(CC) $(LDFLAGS) -o playground main.o ../../.libs/$(LIBNAME).a

clean:
	rm -f *.o playground

.PHONY: build clean
