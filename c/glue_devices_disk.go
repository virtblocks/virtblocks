// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"github.com/virtblocks/virtblocks/c/pkg/types"
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"unsafe"
)

//export devices_disk_new
func devices_disk_new() C.uint {
	var goDisk = devices.NewDisk()
	return C.uint(objects.DevicesDiskAdd(goDisk))
}

//export devices_disk_free
func devices_disk_free(cDisk C.uint) {
	objects.DevicesDiskDel(uint(cDisk))
}

//export devices_disk_set_filename
func devices_disk_set_filename(cDisk C.uint, cFilename *C.char) {
	var goDisk = objects.DevicesDiskGet(uint(cDisk))
	var goFilename = C.GoString(cFilename)
	goDisk.SetFilename(goFilename)
}

//export devices_disk_get_qemu_commandline
func devices_disk_get_qemu_commandline(cDisk C.uint, cError *C.uint) C.uint {
	var goDisk = objects.DevicesDiskGet(uint(cDisk))

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
