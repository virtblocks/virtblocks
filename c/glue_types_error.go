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

//export error_free
func error_free(cError C.uint) {
	// We need to release the error message manually because C types are
	// unmanaged and thus exempt from garbage collection
	var goError = objects.ErrorGet(uint(cError))
	C.free(unsafe.Pointer(goError.Message()))
	// Everything else will be garbage collected once the object is removed
	// from the corresponding table and thus has zero references to it left
	objects.ErrorDel(uint(cError))
}

//export error_get_domain
func error_get_domain(cError C.uint) C.uint {
	var goError = objects.ErrorGet(uint(cError))
	return C.uint(goError.Domain())
}

//export error_get_code
func error_get_code(cError C.uint) C.uint {
	var goError = objects.ErrorGet(uint(cError))
	return C.uint(goError.Code())
}

//export error_get_message
func error_get_message(cError C.uint) *C.char {
	var goError = objects.ErrorGet(uint(cError))
	return (*C.char)(goError.Message())
}
