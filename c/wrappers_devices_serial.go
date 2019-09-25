// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesSerial*
virtblocks_devices_serial_new()
{
    return devices_serial_wrap(devices_serial_new());
}

void
virtblocks_devices_serial_free(VirtBlocksDevicesSerial *serial)
{
    if (serial == NULL) return;
    devices_serial_free(serial->goPtr);
    free(serial);
}

void
virtblocks_devices_serial_set_path(VirtBlocksDevicesSerial *serial,
                                   const char *path)
{
    assert(serial != NULL);
    devices_serial_set_path(serial->goPtr, (char *) path);
}

VirtBlocksArray*
virtblocks_devices_serial_get_qemu_commandline(VirtBlocksDevicesSerial *serial,
                                               VirtBlocksError **error)
{
    unsigned int goPtr;
    unsigned int ret;

    assert(serial != NULL);
    assert(error != NULL);

    ret = devices_serial_get_qemu_commandline(serial->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return array_wrap(ret);
}
*/
import "C"
