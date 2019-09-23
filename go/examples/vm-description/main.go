// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

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
