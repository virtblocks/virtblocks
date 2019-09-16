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
	"unsafe"
)

//export array_free
func array_free(cArray C.int) {
	objects.ArrayDel(int(cArray))
}

//export array_get_length
func array_get_length(cArray C.int) C.uint {
	var goArray = objects.ArrayGet(int(cArray))
	return C.uint(goArray.Length())
}

//export array_get_item
func array_get_item(cArray C.int, cIndex C.uint) unsafe.Pointer {
	var goArray = objects.ArrayGet(int(cArray))
	return goArray.Item(uint(cIndex))
}
