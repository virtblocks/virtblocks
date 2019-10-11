// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"sync"
)

var vmDiskObjectsLock sync.RWMutex
var vmDiskObjects = make([]*vm.Disk, 1)

func VmDiskAdd(disk *vm.Disk) uint {
	vmDiskObjectsLock.Lock()
	defer vmDiskObjectsLock.Unlock()

	vmDiskObjects = append(vmDiskObjects, disk)
	return uint(len(vmDiskObjects) - 1)
}

func VmDiskGet(ref uint) *vm.Disk {
	vmDiskObjectsLock.RLock()
	defer vmDiskObjectsLock.RUnlock()

	return vmDiskObjects[ref]
}

func VmDiskDel(ref uint) {
	vmDiskObjectsLock.Lock()
	defer vmDiskObjectsLock.Unlock()

	vmDiskObjects[ref] = nil
}
