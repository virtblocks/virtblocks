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

//export vm_memballoon_new
func vm_memballoon_new() C.uint {
	var goMemballoon = vm.NewMemballoon()
	return C.uint(objects.VmMemballoonAdd(goMemballoon))
}

//export vm_memballoon_free
func vm_memballoon_free(cMemballoon C.uint) {
	objects.VmMemballoonDel(uint(cMemballoon))
}

//export vm_memballoon_set_model
func vm_memballoon_set_model(cMemballoon C.uint, cModel C.uint) {
	var goMemballoon = objects.VmMemballoonGet(uint(cMemballoon))
	goMemballoon.SetModel(vm.MemballoonModel(cModel))
}

//export vm_memballoon_get_model
func vm_memballoon_get_model(cMemballoon C.uint) C.uint {
	var goMemballoon = objects.VmMemballoonGet(uint(cMemballoon))
	return C.uint(goMemballoon.Model())
}

//export vm_memballoon_to_string
func vm_memballoon_to_string(cMemballoon C.uint) *C.char {
	var goMemballoon = objects.VmMemballoonGet(uint(cMemballoon))
	return C.CString(goMemballoon.String())
}
