// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/command"
	"os"
	"time"
)

func main() {
	var command = command.NewCommand("./test.sh").AddArgs("foo", "bar")
	command.SetStdout(os.Stdout).SetStderr(os.Stderr)

	err := command.Spawn()
	fmt.Printf("Command started: err=%v\n", err)

	id, err := command.Id()
	fmt.Printf("Command ID: id=%d, err=%v\n", id, err)

	time.Sleep(1 * time.Second)

	err = command.Kill()
	fmt.Printf("Command killed: err=%v\n", err)

	rc, err := command.Wait()
	fmt.Printf("Command returned: rc=%d, err=%v\n", rc, err)
}
