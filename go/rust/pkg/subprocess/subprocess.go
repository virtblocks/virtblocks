// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package subprocess

// #cgo CPPFLAGS: -I../../../../c/rust/
// #cgo LDFLAGS: -L../../../../c/rust/.libs/ -lvirtblocks_c_rust -pthread -lm -ldl
// #include <virtblocks.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Error struct {
	code    int
	message string
}

// Convert a C error to a Go error. For other long-lived objects we expect
// the callers to call Free() once they no longer need them, but requiring
// the same for short-lived objects such as errors would be too much.
// Unfortunately that means we have to make a copy of all information
// contained in the error, notably the message
func newError(err *C.VirtBlocksError) Error {
	defer C.virtblocks_error_free(err)

	var cCode = C.virtblocks_error_get_code(err)
	var cMessage = C.virtblocks_error_get_message(err)
	defer C.free(unsafe.Pointer(cMessage))

	return Error{code: int(cCode), message: C.GoString(cMessage)}
}

func (self Error) Code() int {
	return self.code
}

func (self Error) Error() string {
	return self.message
}

type SubProcess struct {
	ptr *C.VirtBlocksSubprocess
}

func NewSubProcess(prog string) *SubProcess {
	cs := C.CString(prog)
	defer C.free(unsafe.Pointer(cs))
	s := &SubProcess{ptr: C.virtblocks_subprocess_new(cs)}
	runtime.SetFinalizer(s, (*SubProcess).Free)
	return s
}

func (self *SubProcess) Spawn() error {
	var cError *C.VirtBlocksError = nil
	var cRet = C.virtblocks_subprocess_spawn(self.ptr, &cError)

	if !cRet {
		return newError(cError)
	}
	return nil
}

func (self *SubProcess) Id() uint {
	return uint(C.virtblocks_subprocess_id(self.ptr))
}

func (self *SubProcess) Free() {
	C.virtblocks_subprocess_free(self.ptr)
	self.ptr = nil
}
