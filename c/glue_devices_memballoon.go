// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"github.com/virtblocks/virtblocks/go/pkg/devices"
)

//export devices_memballoon_new
func devices_memballoon_new() C.uint {
	var goMemballoon = devices.NewMemballoon()
	return C.uint(objects.DevicesMemballoonAdd(goMemballoon))
}

//export devices_memballoon_free
func devices_memballoon_free(cMemballoon C.uint) {
	objects.DevicesMemballoonDel(uint(cMemballoon))
}

//export devices_memballoon_set_model
func devices_memballoon_set_model(cMemballoon C.uint, cModel C.uint) {
	var goMemballoon = objects.DevicesMemballoonGet(uint(cMemballoon))
	goMemballoon.SetModel(devices.MemballoonModel(cModel))
}

//export devices_memballoon_get_model
func devices_memballoon_get_model(cMemballoon C.uint) C.uint {
	var goMemballoon = objects.DevicesMemballoonGet(uint(cMemballoon))
	return C.uint(goMemballoon.Model())
}

//export devices_memballoon_to_string
func devices_memballoon_to_string(cMemballoon C.uint) *C.char {
	var goMemballoon = objects.DevicesMemballoonGet(uint(cMemballoon))
	return C.CString(goMemballoon.String())
}
