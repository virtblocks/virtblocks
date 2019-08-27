// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/rust/pkg/devices"
	"github.com/virtblocks/virtblocks/go/rust/pkg/util"
)

func main() {
	var file_name = util.BuildFileName("guest", ".xml")
	fmt.Println(file_name)

	var memballoon = devices.NewMemballoon()
	defer memballoon.Free()

	fmt.Println(memballoon)

	memballoon.SetModel(devices.MemballoonModelVirtio)
	fmt.Println(memballoon)
}
