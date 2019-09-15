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

//export error_free
func error_free(cError C.int) {
	objects.ErrorDel(int(cError))
}

//export error_get_domain
func error_get_domain(cError C.int) C.int {
	var goError = objects.ErrorGet(int(cError))
	return C.int(goError.Domain())
}

//export error_get_code
func error_get_code(cError C.int) C.int {
	var goError = objects.ErrorGet(int(cError))
	return C.int(goError.Code())
}

//export error_get_message
func error_get_message(cError C.int) *C.char {
	var goError = objects.ErrorGet(int(cError))
	return C.CString(goError.Message())
}

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
