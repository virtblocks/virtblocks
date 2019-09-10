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
	"github.com/virtblocks/virtblocks/go/native/pkg/subprocess"
)

//export subprocess_new
func subprocess_new(cCmd *C.char) C.int {
	if cCmd == nil {
		panic("cCmd == nil")
	}
	var goCmd = C.GoString(cCmd)

	var goSubprocess = subprocess.NewSubprocess(goCmd)
	return C.int(objects.SubprocessAdd(goSubprocess))
}

//export subprocess_free
func subprocess_free(cSubprocess C.int) {
	objects.SubprocessDel(int(cSubprocess))
}

//export subprocess_spawn
func subprocess_spawn(cSubprocess C.int) C.int {
	var goSubprocess = objects.SubprocessGet(int(cSubprocess))
	err := goSubprocess.Spawn()
	if err != nil {
		panic("error handling not implemented")
	}
	return 0
}

//export subprocess_wait
func subprocess_wait(cSubprocess C.int) C.int {
	var goSubprocess = objects.SubprocessGet(int(cSubprocess))
	err := goSubprocess.Wait()
	if err != nil {
		panic("error handling not implemented")
	}
	return 0
}
