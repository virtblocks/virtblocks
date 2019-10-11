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

//export vm_disk_new
func vm_disk_new() C.uint {
	var goDisk = vm.NewDisk()
	return C.uint(objects.VmDiskAdd(goDisk))
}

//export vm_disk_free
func vm_disk_free(cDisk C.uint) {
	objects.VmDiskDel(uint(cDisk))
}

//export vm_disk_set_filename
func vm_disk_set_filename(cDisk C.uint, cFilename *C.char) {
	var goDisk = objects.VmDiskGet(uint(cDisk))
	var goFilename = C.GoString(cFilename)
	goDisk.SetFilename(goFilename)
}

//export vm_disk_get_qemu_commandline
func vm_disk_get_qemu_commandline(cDisk C.uint, cError *C.uint) C.uint {
	var goDisk = objects.VmDiskGet(uint(cDisk))

	var goRet, goErr = goDisk.QemuCommandLine()

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
