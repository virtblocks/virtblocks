// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"sync"
)

var devicesDiskObjectsLock sync.RWMutex
var devicesDiskObjects = make([]*devices.Disk, 1)

func DevicesDiskAdd(disk *devices.Disk) uint {
	devicesDiskObjectsLock.Lock()
	defer devicesDiskObjectsLock.Unlock()

	devicesDiskObjects = append(devicesDiskObjects, disk)
	return uint(len(devicesDiskObjects) - 1)
}

func DevicesDiskGet(ref uint) *devices.Disk {
	devicesDiskObjectsLock.RLock()
	defer devicesDiskObjectsLock.RUnlock()

	return devicesDiskObjects[ref]
}

func DevicesDiskDel(ref uint) {
	devicesDiskObjectsLock.Lock()
	defer devicesDiskObjectsLock.Unlock()

	devicesDiskObjects[ref] = nil
}
