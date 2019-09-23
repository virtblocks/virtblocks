// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksCommand*
command_wrap(unsigned int goPtr)
{
    VirtBlocksCommand *self = malloc(sizeof(VirtBlocksCommand));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
