// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/devices"
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
