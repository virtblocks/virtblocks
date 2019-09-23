// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksError*
error_wrap(unsigned int goPtr)
{
    VirtBlocksError *self = malloc(sizeof(VirtBlocksError));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
