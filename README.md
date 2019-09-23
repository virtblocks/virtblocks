# Virt Blocks

Reusable components for managing hypervisors, written in safe
languages.

## Status

Kinda just figuring out whether the idea can work at all.

Accordingly, the focus right now is not on writing code that does
something useful, but rather in exploring the challenges and cost of
integration, and more specifically in seeking out problematic
scenarios in advance so that we don't get too far down a path before
finding out it won't work for us.

## Requirements

The code needs to be callable from both C and Go.

## Repository structure

It's basically a bunch of mini-projects stored in a single repository
for convenience; each one can be consumed on its own, although there
are dependencies between some.

Specifically:

* `go/`: Virt Blocks, native implementation;

* `c/`: C bindings for Virt Blocks, calling to the native Go
  implementation.

## Build system

Each sub-project has its own `Makefile` which exposes a number of
standard targets:

* `build`, `clean`, `test`: pretty self-explanatory;

* `fmt`, `verify-fmt`, `vet`: automatically format code, verify it
  is formatted properly, and try to spot issues;

* `run-examples`: run all example programs.

Some targets are no-op for certain sub-projects because they don't
make sense in that context.

The top-level `Makefile` provides a convenient way to build and test
all sub-projects at once.

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## License

See [`LICENSE.md`](LICENSE.md) for details.
