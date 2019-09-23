// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/playground"
	"sync"
	"unsafe"
)

// This struct is used to store additional information about a
// playground.Toy that is not part of the native Go object, but is
// necessary for the C bindings to work
type PlaygroundToyExtra struct {
	callback unsafe.Pointer
	data     unsafe.Pointer
	dataFree unsafe.Pointer
}

func NewPlaygroundToyExtra(callback unsafe.Pointer, data unsafe.Pointer, dataFree unsafe.Pointer) *PlaygroundToyExtra {
	return &PlaygroundToyExtra{callback: callback, data: data, dataFree: dataFree}
}

func (self *PlaygroundToyExtra) Callback() unsafe.Pointer {
	return self.callback
}

func (self *PlaygroundToyExtra) Data() unsafe.Pointer {
	return self.data
}

func (self *PlaygroundToyExtra) DataFree() unsafe.Pointer {
	return self.dataFree
}

// Ties each playground.Toy with the corresponding extra data
type playgroundToyObject struct {
	ptr   *playground.Toy
	extra *PlaygroundToyExtra
}

// Go objects are not allow to cross the language boundary, but we still
// need the C code to manipulate them. The way we achieve this is by
// having an array of objects on the Go side and passing "references" to
// them back and forth from the C side: each "reference" is simply the
// index of the object in the array, with 0 being a special "reference"
// that always points to nil
var playgroundToyObjectsLock sync.RWMutex
var playgroundToyObjects = make([]*playgroundToyObject, 1)

// Register a playground.Toy (and the corresponding extra data) so that
// it's later accessible from C
func PlaygroundToyAdd(toy *playground.Toy, extra *PlaygroundToyExtra) uint {
	playgroundToyObjectsLock.Lock()
	defer playgroundToyObjectsLock.Unlock()

	var object = &playgroundToyObject{ptr: toy, extra: extra}
	playgroundToyObjects = append(playgroundToyObjects, object)
	return uint(len(playgroundToyObjects) - 1)
}

// Look up a playground.Toy (and the corresponding extra data) given
// a reference. This basically converts a C object to a Go object
func PlaygroundToyGet(ref uint) (*playground.Toy, *PlaygroundToyExtra) {
	playgroundToyObjectsLock.RLock()
	defer playgroundToyObjectsLock.RUnlock()

	var object = playgroundToyObjects[ref]
	return object.ptr, object.extra
}

// Return the reference for a specific playground.Toy. This is usually only
// necessary when callbacks crossing the language boundary are involved
func PlaygroundToyRef(toy *playground.Toy) uint {
	playgroundToyObjectsLock.RLock()
	defer playgroundToyObjectsLock.RUnlock()

	for ref, object := range playgroundToyObjects {
		if object != nil && object.ptr == toy {
			return uint(ref)
		}
	}
	panic("Pointer not found")
}

// Make a playground.Toy no longer accessible from C, and cause it to
// be garbage collected
func PlaygroundToyDel(ref uint) {
	playgroundToyObjectsLock.Lock()
	defer playgroundToyObjectsLock.Unlock()

	playgroundToyObjects[ref] = nil
}
