// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/playground"
)

type Input struct {
	toy *playground.Toy
	ext string
}

// Create an object that already has a *local* closure attached
func create_baz() *playground.Toy {
	var ext_is_baz = func(_toy *playground.Toy, ext string) bool {
		return ext == "baz"
	}
	return playground.NewToy("baz", ext_is_baz)
}

// Function to be used as a filter
func base_and_ext_match(toy *playground.Toy, ext string) bool {
	return toy.Base() == ext
}

func main() {
	// Closure to be used as a filter. In this case we're using a local
	// variable inside the closure
	var base = "foo"
	var base_is_foo = func(toy *playground.Toy, _ext string) bool {
		return toy.Base() == base
	}

	var foo_is_foo_toy = playground.NewToy("foo", base_is_foo)
	var foo_match_toy = playground.NewToy("foo", base_and_ext_match)
	var bar_is_foo_toy = playground.NewToy("bar", base_is_foo)
	var bar_match_toy = playground.NewToy("bar", base_and_ext_match)
	var baz_toy = create_baz()

	var inputs = [...]Input{
		Input{toy: foo_is_foo_toy, ext: "exe"},
		Input{toy: foo_match_toy, ext: "exe"},
		Input{toy: bar_is_foo_toy, ext: "bar"},
		Input{toy: bar_match_toy, ext: "bar"},
		Input{toy: baz_toy, ext: "baz"},
		Input{toy: baz_toy, ext: "quux"},
	}

	for _, input := range inputs {
		res, err := input.toy.Run(input.ext)
		if err == nil {
			fmt.Printf("Result: %s\n", res)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}
