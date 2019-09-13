// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import "C"

import (
	"github.com/virtblocks/virtblocks/c/go/pkg/objects"
	"github.com/virtblocks/virtblocks/c/go/pkg/types"
	"github.com/virtblocks/virtblocks/go/native/pkg/devices"
	"unsafe"
)

//export devices_memballoon_new
func devices_memballoon_new() C.int {
	var goMemballoon = devices.NewMemballoon()
	return C.int(objects.DevicesMemballoonAdd(goMemballoon))
}

//export devices_memballoon_free
func devices_memballoon_free(cMemballoon C.int) {
	objects.DevicesMemballoonDel(int(cMemballoon))
}

//export devices_memballoon_set_model
func devices_memballoon_set_model(cMemballoon C.int, cModel C.int) {
	var goMemballoon = objects.DevicesMemballoonGet(int(cMemballoon))
	goMemballoon.SetModel(devices.MemballoonModel(cModel))
}

//export devices_memballoon_get_model
func devices_memballoon_get_model(cMemballoon C.int) C.int {
	var goMemballoon = objects.DevicesMemballoonGet(int(cMemballoon))
	return C.int(goMemballoon.Model())
}

//export devices_memballoon_to_string
func devices_memballoon_to_string(cMemballoon C.int) *C.char {
	var goMemballoon = objects.DevicesMemballoonGet(int(cMemballoon))
	return C.CString(goMemballoon.String())
}

//export devices_disk_new
func devices_disk_new() C.int {
	var goDisk = devices.NewDisk()
	return C.int(objects.DevicesDiskAdd(goDisk))
}

//export devices_disk_free
func devices_disk_free(cDisk C.int) {
	objects.DevicesDiskDel(int(cDisk))
}

//export devices_disk_set_filename
func devices_disk_set_filename(cDisk C.int, cFilename *C.char) {
	var goDisk = objects.DevicesDiskGet(int(cDisk))
	var goFilename = C.GoString(cFilename)
	goDisk.SetFilename(goFilename)
}

//export devices_disk_get_qemu_commandline
func devices_disk_get_qemu_commandline(cDisk C.int, cError *C.int) C.int {
	var goDisk = objects.DevicesDiskGet(int(cDisk))

	var goRet, goErr = goDisk.QemuCommandLine()

	if goErr == nil {
		*cError = 0

		var goTmp = make([]unsafe.Pointer, len(goRet))
		for i, v := range goRet {
			goTmp[i] = unsafe.Pointer(C.CString(v))
		}
		var tmp = types.NewArray(goTmp)

		return C.int(objects.ArrayAdd(tmp))
	} else {
		var tmp = types.NewError(goErr)
		*cError = C.int(objects.ErrorAdd(tmp))

		return 0
	}
}
