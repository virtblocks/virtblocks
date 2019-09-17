// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/c/pkg/types"
	"sync"
)

var arrayObjectsLock sync.RWMutex
var arrayObjects = make([]*types.Array, 1)

func ArrayAdd(err *types.Array) uint {
	arrayObjectsLock.Lock()
	defer arrayObjectsLock.Unlock()

	arrayObjects = append(arrayObjects, err)
	return uint(len(arrayObjects) - 1)
}

func ArrayGet(ref uint) *types.Array {
	arrayObjectsLock.RLock()
	defer arrayObjectsLock.RUnlock()

	return arrayObjects[ref]
}

func ArrayDel(ref uint) {
	arrayObjectsLock.Lock()
	defer arrayObjectsLock.Unlock()

	arrayObjects[ref] = nil
}
