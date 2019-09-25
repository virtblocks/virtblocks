// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"sync"
)

var devicesSerialObjectsLock sync.RWMutex
var devicesSerialObjects = make([]*devices.Serial, 1)

func DevicesSerialAdd(serial *devices.Serial) uint {
	devicesSerialObjectsLock.Lock()
	defer devicesSerialObjectsLock.Unlock()

	devicesSerialObjects = append(devicesSerialObjects, serial)
	return uint(len(devicesSerialObjects) - 1)
}

func DevicesSerialGet(ref uint) *devices.Serial {
	devicesSerialObjectsLock.RLock()
	defer devicesSerialObjectsLock.RUnlock()

	return devicesSerialObjects[ref]
}

func DevicesSerialDel(ref uint) {
	devicesSerialObjectsLock.Lock()
	defer devicesSerialObjectsLock.Unlock()

	devicesSerialObjects[ref] = nil
}
