// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/c/go/pkg/types"
	"sync"
)

var errorObjectsLock sync.RWMutex
var errorObjects = make([]*types.Error, 1)

func ErrorAdd(err *types.Error) int {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects = append(errorObjects, err)
	return len(errorObjects) - 1
}

func ErrorGet(ref int) *types.Error {
	errorObjectsLock.RLock()
	defer errorObjectsLock.RUnlock()

	return errorObjects[ref]
}

func ErrorDel(ref int) {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects[ref] = nil
}

var arrayObjectsLock sync.RWMutex
var arrayObjects = make([]*types.Array, 1)

func ArrayAdd(err *types.Array) int {
	arrayObjectsLock.Lock()
	defer arrayObjectsLock.Unlock()

	arrayObjects = append(arrayObjects, err)
	return len(arrayObjects) - 1
}

func ArrayGet(ref int) *types.Array {
	arrayObjectsLock.RLock()
	defer arrayObjectsLock.RUnlock()

	return arrayObjects[ref]
}

func ArrayDel(ref int) {
	arrayObjectsLock.Lock()
	defer arrayObjectsLock.Unlock()

	arrayObjects[ref] = nil
}
