// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/c/pkg/types"
	"sync"
)

var errorObjectsLock sync.RWMutex
var errorObjects = make([]*types.Error, 1)

func ErrorAdd(err *types.Error) uint {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects = append(errorObjects, err)
	return uint(len(errorObjects) - 1)
}

func ErrorGet(ref uint) *types.Error {
	errorObjectsLock.RLock()
	defer errorObjectsLock.RUnlock()

	return errorObjects[ref]
}

func ErrorDel(ref uint) {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects[ref] = nil
}
