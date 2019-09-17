// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"sync"
)

var vmDescriptionObjectsLock sync.RWMutex
var vmDescriptionObjects = make([]*vm.Description, 1)

func VmDescriptionAdd(desc *vm.Description) uint {
	vmDescriptionObjectsLock.Lock()
	defer vmDescriptionObjectsLock.Unlock()

	vmDescriptionObjects = append(vmDescriptionObjects, desc)
	return uint(len(vmDescriptionObjects) - 1)
}

func VmDescriptionGet(ref uint) *vm.Description {
	vmDescriptionObjectsLock.RLock()
	defer vmDescriptionObjectsLock.RUnlock()

	return vmDescriptionObjects[ref]
}

func VmDescriptionDel(ref uint) {
	vmDescriptionObjectsLock.Lock()
	defer vmDescriptionObjectsLock.Unlock()

	vmDescriptionObjects[ref] = nil
}
