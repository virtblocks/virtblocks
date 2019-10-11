// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmMemballoon*
vm_memballoon_wrap(unsigned int goPtr)
{
    VirtBlocksVmMemballoon *self = malloc(sizeof(VirtBlocksVmMemballoon));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
