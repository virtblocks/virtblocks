// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"os"
)

func main() {
	var disk = vm.NewDisk().SetFilename("test.qcow2")
	var serial = vm.NewSerial().SetPath("test.socket")

	var desc = vm.NewDescription()
	desc.SetCpus(1).SetMemory(512)
	desc.SetDisk(disk).SetSerial(serial)

	var args, err = desc.QemuCommandLine()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}
