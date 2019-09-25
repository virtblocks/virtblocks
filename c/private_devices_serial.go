// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesSerial*
devices_serial_wrap(unsigned int goPtr)
{
    VirtBlocksDevicesSerial *self = malloc(sizeof(VirtBlocksDevicesSerial));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
