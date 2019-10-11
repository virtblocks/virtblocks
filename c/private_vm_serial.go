// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmSerial*
vm_serial_wrap(unsigned int goPtr)
{
    VirtBlocksVmSerial *self = malloc(sizeof(VirtBlocksVmSerial));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
