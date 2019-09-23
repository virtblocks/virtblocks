// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"unsafe"
)

//export array_free
func array_free(cArray C.uint) {
	objects.ArrayDel(uint(cArray))
}

//export array_get_length
func array_get_length(cArray C.uint) C.uint {
	var goArray = objects.ArrayGet(uint(cArray))
	return C.uint(goArray.Length())
}

//export array_get_item
func array_get_item(cArray C.uint, cIndex C.uint) unsafe.Pointer {
	var goArray = objects.ArrayGet(uint(cArray))
	return goArray.Item(uint(cIndex))
}
