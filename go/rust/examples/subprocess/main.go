// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/rust/pkg/subprocess"
)

func main() {
	var sub = subprocess.NewSubProcess("ls")
	sub.Spawn()
	fmt.Println("PID:", sub.Id())
}
