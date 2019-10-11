// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmDisk*
virtblocks_vm_disk_new()
{
    return vm_disk_wrap(vm_disk_new());
}

void
virtblocks_vm_disk_free(VirtBlocksVmDisk *disk)
{
    if (disk == NULL) return;
    vm_disk_free(disk->goPtr);
    free(disk);
}

void
virtblocks_vm_disk_set_filename(VirtBlocksVmDisk *disk,
                                const char *filename)
{
    assert(disk != NULL);
    vm_disk_set_filename(disk->goPtr, (char *) filename);
}

VirtBlocksArray*
virtblocks_vm_disk_get_qemu_commandline(VirtBlocksVmDisk *disk,
                                        VirtBlocksError **error)
{
    unsigned int goPtr;
    unsigned int ret;

    assert(disk != NULL);
    assert(error != NULL);

    ret = vm_disk_get_qemu_commandline(disk->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return array_wrap(ret);
}
*/
import "C"
