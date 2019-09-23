// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesDisk*
devices_disk_wrap(unsigned int goPtr)
{
    VirtBlocksDevicesDisk *self = malloc(sizeof(VirtBlocksDevicesDisk));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
