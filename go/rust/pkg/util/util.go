// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package util

// #cgo CPPFLAGS: -I../../../../c/rust/
// #cgo LDFLAGS: -L../../../../c/rust/target/debug/ -lvirtblocks_c_rust -ldl
// #include <stdlib.h>
// #include <virtblocks.h>
import "C"

import (
	"unsafe"
)

func BuildFileName(base string, ext string) string {
	var cBase = C.CString(base)
	var cExt = C.CString(ext)
	var cFileName *C.char
	defer C.free(unsafe.Pointer(cFileName))

	C.virtblocks_util_build_file_name(&cFileName, cBase, cExt)

	return C.GoString(cFileName)
}
