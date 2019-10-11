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

VirtBlocksArray*
virtblocks_vm_serial_get_qemu_commandline(VirtBlocksVmSerial *serial,
                                          VirtBlocksError **error)
{
    unsigned int goPtr;
    unsigned int ret;

    assert(serial != NULL);
    assert(error != NULL);

    ret = vm_serial_get_qemu_commandline(serial->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return array_wrap(ret);
}
*/
import "C"
