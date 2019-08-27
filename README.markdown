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

We're evaluating both Rust and Go at the moment, but in the not too
distant future we'll have to pick one or the other and start writing
non-toy code.

## Requirements

The code needs to be callable from both C and Go.

## Repository structure

It's basically a bunch of mini-projects stored in a single repository
for convenience; each one can be consumed on its own, although there
are dependencies between some.

Specifically:

* `rust/native` and `go/native`: Virt Blocks, implemented in Rust and
  Go respectively;

* `c/rust` and `c/go`: C bindings for Virt Blocks, calling to the
  Rust and Go implementation respectively;

* `go/rust`: Go bindings for the Rust implementation of Virt Blocks.

There is no `rust/go` because that's not currently a requirement.

## Build system

Each sub-project has its own `Makefile` which exposes a number of
standard targets:

* `build`, `clean`, `test`: pretty self-explanatory;

* `fmt`, `verify-fmt`, `vet`: automatically format code, verify it
  is formatted properly, and try to spot issues;

* `run`: run all example programs.

Some targets are no-op for certain sub-projects because they don't
make sense in that context.

The top-level `Makefile` provides a convenient way to build and test
all sub-projects at once.

## License

Virt Blocks is distributed under the terms of the MIT License. See
the `LICENSE` file for details.

## Contributing

GitHub pull requests and GitLab merge requests are both accepted, but
the former is preferred because it's hooked up with CI.
