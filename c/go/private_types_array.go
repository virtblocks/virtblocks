// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksArray*
array_wrap(int goPtr)
{
    VirtBlocksArray *self = malloc(sizeof(VirtBlocksArray));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
