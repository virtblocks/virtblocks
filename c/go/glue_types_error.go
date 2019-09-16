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
