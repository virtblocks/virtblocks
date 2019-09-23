// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesMemballoon*
devices_memballoon_wrap(unsigned int goPtr)
{
    VirtBlocksDevicesMemballoon *self = malloc(sizeof(VirtBlocksDevicesMemballoon));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
