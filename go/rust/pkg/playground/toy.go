// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package playground

// #cgo CPPFLAGS: -I../../../../c/rust/
// #cgo LDFLAGS: -L../../../../c/rust/.libs/ -lvirtblocks_c_rust -pthread -lm -ldl
// #include <virtblocks.h>
// #include "toy_private.h"
import "C"

import (
	"unsafe"
)

type ToyErrorCode int

const (
	CallbackFailed = ToyErrorCode(C.VIRTBLOCKS_PLAYGROUND_TOY_ERROR_CALLBACK_FAILED)
)

type ToyError struct {
	code    ToyErrorCode
	message string
}

// Convert a C error to a Go error. For other long-lived objects we expect
// the callers to call Free() once they no longer need them, but requiring
// the same for short-lived objects such as errors would be too much.
// Unfortunately that means we have to make a copy of all information
// contained in the error, notably the message
func newToyError(err *C.VirtBlocksError) ToyError {
	defer C.virtblocks_error_free(err)

	var cCode = C.virtblocks_error_get_code(err)
	var cMessage = C.virtblocks_error_get_message(err)
	defer C.free(unsafe.Pointer(cMessage))

	return ToyError{code: ToyErrorCode(cCode), message: C.GoString(cMessage)}
}

func (self ToyError) Code() ToyErrorCode {
	return self.code
}

func (self ToyError) Error() string {
	return self.message
}

type ToyCallback func(*Toy, string) bool

// We need to keep a reference to some extra data, in this case the filter
// callback, around so that we can release resources at the appropriate
// time
type toyExtra struct {
	filterPtr int
}

type Toy struct {
	ptr   *C.VirtBlocksPlaygroundToy
	extra *toyExtra
}

//export playground_toy_callback_trampoline_go
func playground_toy_callback_trampoline_go(cToy *C.VirtBlocksPlaygroundToy, cExt *C.char, cFilter C.int) C.int {
	var goToy = &Toy{ptr: cToy}
	var goExt = C.GoString(cExt)
	var goFilter = ToyCallbackGet(int(cFilter))

	var goRet = (*goFilter)(goToy, goExt)

	// Returning a C bool from a Go function callable from C is annoying,
	// so we just return either 0 or 1 and let the C trampoline deal with
	// converting that into a a C bool
	if goRet {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

func NewToy(goBase string, goFilter ToyCallback) *Toy {
	var cBase = C.CString(goBase)
	var cFilter = ToyCallbackAdd(&goFilter)

	// We can't invoke a Go function from C directly, so we have to go
	// through a double trampoline: the C trampoline will call back into
	// a trampoline implemented in Go which will receive a reference to
	// the user-provided Go callback and finally invoke it
	var cToy = C.virtblocks_playground_toy_new(cBase, C.VirtBlocksPlaygroundToyCallback(C.playground_toy_callback_trampoline_c), unsafe.Pointer(&cFilter), nil)

	return &Toy{ptr: cToy, extra: &toyExtra{filterPtr: cFilter}}
}

func (self *Toy) Free() {
	C.virtblocks_playground_toy_free(self.ptr)
	ToyCallbackDel(self.extra.filterPtr)
}

func (self Toy) Base() string {
	var cRet = C.virtblocks_playground_toy_get_base(self.ptr)
	defer C.free(unsafe.Pointer(cRet))
	return C.GoString(cRet)
}

func (self *Toy) Run(goExt string) (string, error) {
	var cExt = C.CString(goExt)
	var cError *C.VirtBlocksError = nil

	var cRet = C.virtblocks_playground_toy_run(self.ptr, cExt, &cError)

	if cRet != nil {
		return C.GoString(cRet), nil
	} else {
		return "", newToyError(cError)
	}
}
