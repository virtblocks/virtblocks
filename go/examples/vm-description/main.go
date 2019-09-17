// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
)

func main() {
	var disk = devices.NewDisk().SetFilename("test.qcow2")
	var desc = vm.NewDescription().SetMemory(512).SetDisk(disk)

	var args, err = desc.QemuCommandLine()
	if err != nil {
		panic(err)
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}
