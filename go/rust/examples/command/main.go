// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/rust/pkg/command"
)

func main() {
	var command = command.NewCommand("ls")
	command.Spawn()
	fmt.Println("PID:", command.Id())
}
