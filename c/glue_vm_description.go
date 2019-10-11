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

//export vm_description_new
func vm_description_new() C.uint {
	var goDescription = vm.NewDescription()
	return C.uint(objects.VmDescriptionAdd(goDescription))
}

//export vm_description_free
func vm_description_free(cDescription C.uint) {
	objects.VmDescriptionDel(uint(cDescription))
}

//export vm_description_set_emulator
func vm_description_set_emulator(cDescription C.uint, cEmulator *C.char) {
	if cEmulator == nil {
		panic("cEmulator == nil")
	}
	var goDescription = objects.VmDescriptionGet(uint(cDescription))
	var goEmulator = C.GoString(cEmulator)
	goDescription.SetEmulator(goEmulator)
}

//export vm_description_set_disk
func vm_description_set_disk(cDescription C.uint, cDisk C.uint) {
	var goDescription = objects.VmDescriptionGet(uint(cDescription))
	var goDisk = objects.VmDiskGet(uint(cDisk))
	goDescription.SetDisk(goDisk)
}

//export vm_description_set_serial
func vm_description_set_serial(cDescription C.uint, cSerial C.uint) {
	var goDescription = objects.VmDescriptionGet(uint(cDescription))
	var goSerial = objects.VmSerialGet(uint(cSerial))
	goDescription.SetSerial(goSerial)
}

//export vm_description_set_cpus
func vm_description_set_cpus(cDescription C.uint, cCpus C.uint) {
	var goDescription = objects.VmDescriptionGet(uint(cDescription))
	var goCpus = uint(cCpus)
	goDescription.SetCpus(goCpus)
}

//export vm_description_set_memory
func vm_description_set_memory(cDescription C.uint, cMemory C.uint) {
	var goDescription = objects.VmDescriptionGet(uint(cDescription))
	var goMemory = uint(cMemory)
	goDescription.SetMemory(goMemory)
}

//export vm_description_get_qemu_commandline
func vm_description_get_qemu_commandline(cDescription C.uint, cError *C.uint) C.uint {
	var goDescription = objects.VmDescriptionGet(uint(cDescription))

	var goRet, goErr = goDescription.QemuCommandLine()

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
