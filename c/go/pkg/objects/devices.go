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

// Go objects are not allow to cross the language boundary, but we still
// need the C code to manipulate them. The way we achieve this is by
// having an array of objects on the Go side and passing "references" to
// them back and forth from the C side: each "reference" is simply the
// index of the object in the array, with 0 being a special "reference"
// that always points to nil

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
