// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmDescription*
vm_description_wrap(unsigned int goPtr)
{
    VirtBlocksVmDescription *self = malloc(sizeof(VirtBlocksVmDescription));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
