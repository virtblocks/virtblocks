// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/command"
	"time"
)

func main() {
	var command = command.NewCommand("./test.sh")
	command.Spawn()
	id, err := command.Id()
	if err == nil {
		fmt.Printf("PID: %d\n", id)
	}
	time.Sleep(1 * time.Second)
	err = command.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
	}
}
