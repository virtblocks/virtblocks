// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"github.com/virtblocks/virtblocks/go/pkg/command"
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
