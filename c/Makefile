LIBNAME = virtblocks
LDFLAGS = -pthread

all: build

build: lib$(LIBNAME).la

# We need two steps: the first one generates an header file that can be
# used to access the exported glue functions, and the second one builds
# an archive we can link from C. The first step is necessary because the
# wrappers call to the glue code and as such can't be compiled without
# the generated header file
gobuild:
	go tool cgo -exportheader lib$(LIBNAME).h glue_*.go
	go build -buildmode c-archive -o lib$(LIBNAME).a

lib$(LIBNAME).la: libtool.la.in gobuild
	mkdir -p .libs/
	ln -sf ../lib$(LIBNAME).a .libs/lib$(LIBNAME).a
	sed \
	  -e 's/@LIBNAME@/$(LIBNAME)/g' \
	  -e 's/@LDFLAGS@/$(LDFLAGS)/g' \
	  <$< >$@

clean:
	rm -rf *.a *.la .libs/ _obj/ lib$(LIBNAME).h

test: gobuild
	go test

run-examples: build
	for example in $$(ls examples/); do \
	( \
		cd examples/$$example || exit 1; \
		$(MAKE) run-examples || exit 1; \
	) || exit 1; \
	done

fmt:
	go fmt ./...

verify-fmt:
	../scripts/verify-gofmt.sh

vet: gobuild
	go vet ./...

functest:
	@true

.PHONY: build gobuild clean test run-examples fmt verify-fmt vet functest
