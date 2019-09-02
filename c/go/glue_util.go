// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import "C"

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/util"
)

//export util_build_file_name
func util_build_file_name(cFileName **C.char, cBase *C.char, cExt *C.char) C.int {
	if cBase == nil || cExt == nil || cFileName == nil {
		return -1
	}
	var goBase = C.GoString(cBase)
	var goExt = C.GoString(cExt)

	var goRes = util.BuildFileName(goBase, goExt)

	*cFileName = C.CString(goRes)
	return 0
}
