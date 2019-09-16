// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

/*
#include "libvirtblocks_c_go.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesDisk*
virtblocks_devices_disk_new()
{
    return devices_disk_wrap(devices_disk_new());
}

void
virtblocks_devices_disk_free(VirtBlocksDevicesDisk *disk)
{
    if (disk == NULL) return;
    devices_disk_free(disk->goPtr);
    free(disk);
}

void
virtblocks_devices_disk_set_filename(VirtBlocksDevicesDisk *disk,
                                     const char *filename)
{
    assert(disk != NULL);
    devices_disk_set_filename(disk->goPtr, (char *) filename);
}

VirtBlocksArray*
virtblocks_devices_disk_get_qemu_commandline(VirtBlocksDevicesDisk *disk,
                                             VirtBlocksError **error)
{
    int goPtr;
    int ret;

    assert(disk != NULL);
    assert(error != NULL);

    ret = devices_disk_get_qemu_commandline(disk->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return array_wrap(ret);
}
*/
import "C"
