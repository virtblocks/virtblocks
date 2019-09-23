// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
)

//export error_free
func error_free(cError C.uint) {
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
	return C.CString(goError.Message())
}
