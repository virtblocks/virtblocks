// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksArray*
array_wrap(unsigned int goPtr)
{
    VirtBlocksArray *self = malloc(sizeof(VirtBlocksArray));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
