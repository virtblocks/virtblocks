// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"github.com/virtblocks/virtblocks/c/pkg/types"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"unsafe"
)

//export vm_serial_new
func vm_serial_new() C.uint {
	var goSerial = vm.NewSerial()
	return C.uint(objects.VmSerialAdd(goSerial))
}

//export vm_serial_free
func vm_serial_free(cSerial C.uint) {
	objects.VmSerialDel(uint(cSerial))
}

//export vm_serial_set_path
func vm_serial_set_path(cSerial C.uint, cPath *C.char) {
	var goSerial = objects.VmSerialGet(uint(cSerial))
	var goPath = C.GoString(cPath)
	goSerial.SetPath(goPath)
}

//export vm_serial_get_qemu_commandline
func vm_serial_get_qemu_commandline(cSerial C.uint, cError *C.uint) C.uint {
	var goSerial = objects.VmSerialGet(uint(cSerial))

	var goRet, goErr = goSerial.QemuCommandLine()

	if goErr == nil {
		*cError = 0

		var goTmp = make([]unsafe.Pointer, len(goRet))
		for i, v := range goRet {
			goTmp[i] = unsafe.Pointer(C.CString(v))
		}
		var tmp = types.NewArray(goTmp)

		return C.uint(objects.ArrayAdd(tmp))
	} else {
		var tmp = types.NewError(goErr)
		*cError = C.uint(objects.ErrorAdd(tmp))

		return 0
	}
}
