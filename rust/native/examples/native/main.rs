// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

use virtblocks_rust_native::devices;

fn main() {
    let mut memballoon = devices::Memballoon::new();

    println!("{}", memballoon);

    memballoon.set_model(devices::MemballoonModel::Virtio);
    println!("{}", memballoon);
}
