// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package playground

/*
#cgo CPPFLAGS: -I../../../../c/rust/
#cgo LDFLAGS: -L../../../../c/rust/.libs/ -lvirtblocks_c_rust -pthread -lm -ldl
#include <virtblocks.h>
#include "toy_private.h"

bool
playground_toy_callback_trampoline_c(const VirtBlocksPlaygroundToy *toy,
                                     const char *ext,
                                     void *data)
{
    return !!playground_toy_callback_trampoline_go((VirtBlocksPlaygroundToy *) toy,
                                                   (char *) ext,
                                                   *((int *) data));
}
*/
import "C"

var toyCallbackObjects = make([]*ToyCallback, 1)

func ToyCallbackAdd(cb *ToyCallback) int {
	toyCallbackObjects = append(toyCallbackObjects, cb)
	return len(toyCallbackObjects) - 1
}

func ToyCallbackGet(ref int) *ToyCallback {
	return toyCallbackObjects[ref]
}

func ToyCallbackDel(ref int) {
	toyCallbackObjects[ref] = nil
}
