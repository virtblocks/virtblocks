// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/command"
	"log"
)

func main() {
	var command = command.NewCommand("./test.sh")
	command.Spawn()
	id, err := command.Id()
	if err == nil {
		log.Printf("PID: %d", id)
	}
	err = command.Wait()
	log.Printf("Command finished with error: %v", err)
}
