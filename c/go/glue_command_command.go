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
	"github.com/virtblocks/virtblocks/go/native/pkg/command"
)

//export command_new
func command_new(cProg *C.char) C.uint {
	if cProg == nil {
		panic("cProg == nil")
	}
	var goProg = C.GoString(cProg)

	var goCommand = command.NewCommand(goProg)
	return C.uint(objects.CommandAdd(goCommand))
}

//export command_free
func command_free(cCommand C.uint) {
	objects.CommandDel(uint(cCommand))
}

//export command_spawn
func command_spawn(cCommand C.uint) C.int {
	var goCommand = objects.CommandGet(uint(cCommand))
	err := goCommand.Spawn()
	if err != nil {
		panic("error handling not implemented")
	}
	return 0
}

//export command_wait
func command_wait(cCommand C.uint) C.int {
	var goCommand = objects.CommandGet(uint(cCommand))
	err := goCommand.Wait()
	if err != nil {
		panic("error handling not implemented")
	}
	return 0
}
