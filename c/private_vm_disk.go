// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmDisk*
vm_disk_wrap(unsigned int goPtr)
{
    VirtBlocksVmDisk *self = malloc(sizeof(VirtBlocksVmDisk));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
