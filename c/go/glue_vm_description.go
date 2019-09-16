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
	"github.com/virtblocks/virtblocks/go/native/pkg/vm"
	"unsafe"
)

//export vm_description_new
func vm_description_new() C.int {
	var goDescription = vm.NewDescription()
	return C.int(objects.VmDescriptionAdd(goDescription))
}

//export vm_description_free
func vm_description_free(cDescription C.int) {
	objects.VmDescriptionDel(int(cDescription))
}

//export vm_description_set_emulator
func vm_description_set_emulator(cDescription C.int, cEmulator *C.char) {
	if cEmulator == nil {
		panic("cEmulator == nil")
	}
	var goDescription = objects.VmDescriptionGet(int(cDescription))
	var goEmulator = C.GoString(cEmulator)
	goDescription.SetEmulator(goEmulator)
}

//export vm_description_set_disk
func vm_description_set_disk(cDescription C.int, cDisk C.int) {
	var goDescription = objects.VmDescriptionGet(int(cDescription))
	var goDisk = objects.DevicesDiskGet(int(cDisk))
	goDescription.SetDisk(goDisk)
}

//export vm_description_set_memory
func vm_description_set_memory(cDescription C.int, cMemory C.int) {
	var goDescription = objects.VmDescriptionGet(int(cDescription))
	var goMemory = int(cMemory)
	goDescription.SetMemory(goMemory)
}

//export vm_description_get_qemu_commandline
func vm_description_get_qemu_commandline(cDescription C.int, cError *C.int) C.int {
	var goDescription = objects.VmDescriptionGet(int(cDescription))

	var goRet, goErr = goDescription.QemuCommandLine()

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
