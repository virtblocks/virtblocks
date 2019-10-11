// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"sync"
)

var vmMemballoonObjectsLock sync.RWMutex
var vmMemballoonObjects = make([]*vm.Memballoon, 1)

func VmMemballoonAdd(memballoon *vm.Memballoon) uint {
	vmMemballoonObjectsLock.Lock()
	defer vmMemballoonObjectsLock.Unlock()

	vmMemballoonObjects = append(vmMemballoonObjects, memballoon)
	return uint(len(vmMemballoonObjects) - 1)
}

func VmMemballoonGet(ref uint) *vm.Memballoon {
	vmMemballoonObjectsLock.RLock()
	defer vmMemballoonObjectsLock.RUnlock()

	return vmMemballoonObjects[ref]
}

func VmMemballoonDel(ref uint) {
	vmMemballoonObjectsLock.Lock()
	defer vmMemballoonObjectsLock.Unlock()

	vmMemballoonObjects[ref] = nil
}
