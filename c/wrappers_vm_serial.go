// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmSerial*
virtblocks_vm_serial_new()
{
    return vm_serial_wrap(vm_serial_new());
}

void
virtblocks_vm_serial_free(VirtBlocksVmSerial *serial)
{
    if (serial == NULL) return;
    vm_serial_free(serial->goPtr);
    free(serial);
}

void
virtblocks_vm_serial_set_path(VirtBlocksVmSerial *serial,
                              const char *path)
{
    assert(serial != NULL);
    vm_serial_set_path(serial->goPtr, (char *) path);
}
*/
import "C"
