// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"sync"
)

var devicesMemballoonObjectsLock sync.RWMutex
var devicesMemballoonObjects = make([]*devices.Memballoon, 1)

func DevicesMemballoonAdd(memballoon *devices.Memballoon) uint {
	devicesMemballoonObjectsLock.Lock()
	defer devicesMemballoonObjectsLock.Unlock()

	devicesMemballoonObjects = append(devicesMemballoonObjects, memballoon)
	return uint(len(devicesMemballoonObjects) - 1)
}

func DevicesMemballoonGet(ref uint) *devices.Memballoon {
	devicesMemballoonObjectsLock.RLock()
	defer devicesMemballoonObjectsLock.RUnlock()

	return devicesMemballoonObjects[ref]
}

func DevicesMemballoonDel(ref uint) {
	devicesMemballoonObjectsLock.Lock()
	defer devicesMemballoonObjectsLock.Unlock()

	devicesMemballoonObjects[ref] = nil
}