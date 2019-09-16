// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/go/pkg/objects"
	"github.com/virtblocks/virtblocks/c/go/pkg/types"
	"github.com/virtblocks/virtblocks/go/native/pkg/playground"
	"unsafe"
)

//export playground_toy_new
func playground_toy_new(cBase *C.char, cFilter C.VirtBlocksPlaygroundToyCallback, data unsafe.Pointer, dataFree C.VirtBlocksPlaygroundToyCallbackDataFree) C.uint {
	if cBase == nil {
		panic("cBase == nil")
	}
	if cFilter == nil {
		panic("cFilter == nil")
	}
	var goBase = C.GoString(cBase)

	// We're going to need this information later, so we store it into an
	// extra data object that will be associated with the playground.Toy
	// that we're about to create
	var goToyExtra = objects.NewPlaygroundToyExtra(unsafe.Pointer(cFilter), data, unsafe.Pointer(dataFree))

	var goFilter = func(goToy *playground.Toy, goExt string) bool {
		// Look up the extra data for the object
		var goPtr = objects.PlaygroundToyRef(goToy)
		var _, goExtra = objects.PlaygroundToyGet(goPtr)

		// Create a new C object wrapping the object reference obtained
		// above
		var cToy = C.playground_toy_wrap(C.uint(goPtr))

		// Extract all information about the C callback we're about to
		// call from the extra data
		var cFilter = C.VirtBlocksPlaygroundToyCallback(goExtra.Callback())
		var cExt = C.CString(goExt)
		var cData = goExtra.Data()

		// Invoke the C callback through a special trampoline
		return bool(C.playground_toy_callback_call(cFilter, cToy, cExt, cData))
	}

	var goToy = playground.NewToy(goBase, goFilter)

	// Register the object and return a reference to it
	return C.uint(objects.PlaygroundToyAdd(goToy, goToyExtra))
}

//export playground_toy_free
func playground_toy_free(cToy C.uint) {
	// Look up the extra data for the object
	var _, goExtra = objects.PlaygroundToyGet(uint(cToy))
	var cData = goExtra.Data()
	var cDataFree = C.VirtBlocksPlaygroundToyCallbackDataFree(goExtra.DataFree())

	// If we have been provided both some user data and a function that
	// can be used to release it, this is a perfect time to do so. Call
	// the dataFree function using a special trampoline
	if cData != nil && cDataFree != nil {
		C.playground_toy_callback_data_free_call(cDataFree, cData)
	}

	// Make the object inaccessible from C
	objects.PlaygroundToyDel(uint(cToy))
}

//export playground_toy_get_base
func playground_toy_get_base(cToy C.uint) *C.char {
	var goToy, _ = objects.PlaygroundToyGet(uint(cToy))

	var goRet = goToy.Base()

	return C.CString(goRet)
}

//export playground_toy_run
func playground_toy_run(cToy C.uint, cExt *C.char, cError *C.uint) *C.char {
	if cExt == nil {
		panic("cExt == nil")
	}
	var goToy, _ = objects.PlaygroundToyGet(uint(cToy))
	var goExt = C.GoString(cExt)

	var goRet, goErr = goToy.Run(goExt)

	if goErr == nil {
		// The Go method completed successfully: set the error to NULL
		// and return the computed value
		*cError = 0

		return C.CString(goRet)
	} else {
		// The Go method reported a failure: convert the native error
		// to a C-accessible types.Error and return it along with a
		// NULL return value
		var tmp = types.NewError(goErr.(playground.ToyError))
		*cError = C.uint(objects.ErrorAdd(tmp))

		return nil
	}
}
