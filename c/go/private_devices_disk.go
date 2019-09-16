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

VirtBlocksDevicesDisk*
devices_disk_wrap(int goPtr)
{
    VirtBlocksDevicesDisk *self = malloc(sizeof(VirtBlocksDevicesDisk));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
