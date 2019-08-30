LIBNAME = libvirtblocks_c_go
LDFLAGS = -pthread

all: build

build: $(LIBNAME).la

# We need to build twice: first just the glue code, which also generates
# the header file; then also the wrappers, which themselves call to the
# glue code and as such can't be compiled withoug the generated header file
gobuild:
	go build -buildmode c-archive -o $(LIBNAME).a glue.go
	go build -buildmode c-archive -o $(LIBNAME).a

$(LIBNAME).la: ../libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../$(LIBNAME).a .libs/$(LIBNAME).a
	sed \
	  -e 's/@LIBNAME@/$(LIBNAME)/g' \
	  -e 's/@LDFLAGS@/$(LDFLAGS)/g' \
	  <$< >$@

clean:
	rm -rf *.a *.la .libs/ $(LIBNAME).h

test: gobuild
	go test

run-examples:
	@true

fmt:
	go fmt ./...

verify-fmt:
	../../scripts/verify-gofmt.sh

vet: gobuild
	go vet ./...

.PHONY: build gobuild clean test run-examples fmt verify-fmt vet
