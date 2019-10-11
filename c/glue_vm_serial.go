// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
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
