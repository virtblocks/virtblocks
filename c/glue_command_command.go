// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

//#include "virtblocks.h"
//#include "private.h"
import "C"

import (
	"github.com/virtblocks/virtblocks/c/pkg/objects"
	"github.com/virtblocks/virtblocks/c/pkg/types"
	"github.com/virtblocks/virtblocks/go/pkg/command"
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

//export command_add_arg
func command_add_arg(cCommand C.uint, cArg *C.char) {
	if cArg == nil {
		panic("cArg == nil")
	}
	var goCommand = objects.CommandGet(uint(cCommand))
	var goArg = C.GoString(cArg)

	goCommand.AddArg(goArg)
}

//export command_spawn
func command_spawn(cCommand C.uint, cError *C.uint) {
	var goCommand = objects.CommandGet(uint(cCommand))

	var goErr = goCommand.Spawn()

	if goErr != nil {
		var tmp = types.NewError(goErr)
		*cError = C.uint(objects.ErrorAdd(tmp))
	} else {
		*cError = 0
	}
}

//export command_get_id
func command_get_id(cCommand C.uint, cError *C.uint) C.uint {
	var goCommand = objects.CommandGet(uint(cCommand))

	var goRet, goErr = goCommand.Id()

	if goErr != nil {
		var tmp = types.NewError(goErr)
		*cError = C.uint(objects.ErrorAdd(tmp))
	} else {
		*cError = 0
	}

	return C.uint(goRet)
}

//export command_kill
func command_kill(cCommand C.uint, cError *C.uint) {
	var goCommand = objects.CommandGet(uint(cCommand))

	var goErr = goCommand.Kill()

	if goErr != nil {
		var tmp = types.NewError(goErr)
		*cError = C.uint(objects.ErrorAdd(tmp))
	} else {
		*cError = 0
	}
}

//export command_wait
func command_wait(cCommand C.uint, cError *C.uint) C.int {
	var goCommand = objects.CommandGet(uint(cCommand))

	var goRet, goErr = goCommand.Wait()

	if goErr != nil {
		var tmp = types.NewError(goErr)
		*cError = C.uint(objects.ErrorAdd(tmp))
	} else {
		*cError = 0
	}

	return C.int(goRet)
}
