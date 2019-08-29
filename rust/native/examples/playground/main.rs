// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

use virtblocks_rust_native::playground;

// Create an object that already has a *local* closure attached
fn create_baz() -> playground::Toy {
    let ext_is_baz = |_toy: &playground::Toy, ext: &str| -> bool { ext == "baz" };
    playground::Toy::new("baz", ext_is_baz)
}

// Function to be used as a filter
fn base_and_ext_match(toy: &playground::Toy, ext: &str) -> bool {
    toy.base() == ext
}

fn main() {
    // Closure to be used as a filter. In this case we're moving a local
    // variable inside the closure
    let base = "foo";
    let base_is_foo = move |toy: &playground::Toy, _ext: &str| -> bool { toy.base() == base };

    let foo_is_foo_toy = playground::Toy::new("foo", base_is_foo);
    let foo_match_toy = playground::Toy::new("foo", base_and_ext_match);
    let bar_is_foo_toy = playground::Toy::new("bar", base_is_foo);
    let bar_match_toy = playground::Toy::new("bar", base_and_ext_match);
    let baz_toy = create_baz();

    let results = [
        foo_is_foo_toy.run("exe"),
        foo_match_toy.run("exe"),
        bar_is_foo_toy.run("bar"),
        bar_match_toy.run("bar"),
        baz_toy.run("baz"),
        baz_toy.run("quux"),
    ];

    for r in &results {
        match r {
            Ok(o) => println!("Result: {}", o),
            Err(e) => println!("Error: {}", e),
        }
    }
}
