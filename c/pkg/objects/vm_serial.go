// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"sync"
)

var vmSerialObjectsLock sync.RWMutex
var vmSerialObjects = make([]*vm.Serial, 1)

func VmSerialAdd(serial *vm.Serial) uint {
	vmSerialObjectsLock.Lock()
	defer vmSerialObjectsLock.Unlock()

	vmSerialObjects = append(vmSerialObjects, serial)
	return uint(len(vmSerialObjects) - 1)
}

func VmSerialGet(ref uint) *vm.Serial {
	vmSerialObjectsLock.RLock()
	defer vmSerialObjectsLock.RUnlock()

	return vmSerialObjects[ref]
}

func VmSerialDel(ref uint) {
	vmSerialObjectsLock.Lock()
	defer vmSerialObjectsLock.Unlock()

	vmSerialObjects[ref] = nil
}
