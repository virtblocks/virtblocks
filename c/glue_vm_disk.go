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
