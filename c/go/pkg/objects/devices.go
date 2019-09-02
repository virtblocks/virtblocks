// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/devices"
)

var devicesMemballoonObjects = make([]*devices.Memballoon, 1)

func DevicesMemballoonAdd(memballoon *devices.Memballoon) int {
	devicesMemballoonObjects = append(devicesMemballoonObjects, memballoon)
	return len(devicesMemballoonObjects) - 1
}

func DevicesMemballoonGet(ref int) *devices.Memballoon {
	return devicesMemballoonObjects[ref]
}

func DevicesMemballoonDel(ref int) {
	devicesMemballoonObjects[ref] = nil
}
