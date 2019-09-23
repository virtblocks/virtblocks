// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

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
