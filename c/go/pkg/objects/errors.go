// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/c/go/pkg/errors"
	"sync"
)

var errorObjectsLock sync.RWMutex
var errorObjects = make([]*errors.Error, 1)

func ErrorAdd(err *errors.Error) int {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects = append(errorObjects, err)
	return len(errorObjects) - 1
}

func ErrorGet(ref int) *errors.Error {
	errorObjectsLock.RLock()
	defer errorObjectsLock.RUnlock()

	return errorObjects[ref]
}

func ErrorDel(ref int) {
	errorObjectsLock.Lock()
	defer errorObjectsLock.Unlock()

	errorObjects[ref] = nil
}
