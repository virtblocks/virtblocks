// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package command

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

type Command struct {
	ptr *C.VirtBlocksCommand
}

func NewCommand(goProg string) *Command {
	cProg := C.CString(goProg)
	defer C.free(unsafe.Pointer(cProg))
	s := &Command{ptr: C.virtblocks_command_new(cProg)}
	runtime.SetFinalizer(s, (*Command).Free)
	return s
}

func (self *Command) Spawn() error {
	var cError *C.VirtBlocksError = nil
	var cRet = C.virtblocks_command_spawn(self.ptr, &cError)

	if !cRet {
		return newError(cError)
	}
	return nil
}

func (self *Command) Id() uint {
	return uint(C.virtblocks_command_id(self.ptr))
}

func (self *Command) Free() {
	C.virtblocks_command_free(self.ptr)
	self.ptr = nil
}
