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

VirtBlocksVmDescription*
virtblocks_vm_description_new()
{
    return vm_description_wrap(vm_description_new());
}

void
virtblocks_vm_description_free(VirtBlocksVmDescription *desc)
{
    if (desc == NULL) return;
    vm_description_free(desc->goPtr);
    free(desc);
}

void
virtblocks_vm_description_set_emulator(VirtBlocksVmDescription *desc,
                                       const char *emulator)
{
    assert(desc != NULL);
    vm_description_set_emulator(desc->goPtr, (char *) emulator);
}

void
virtblocks_vm_description_set_memory(VirtBlocksVmDescription *desc,
                                     int memory)
{
    assert(desc != NULL);
    vm_description_set_memory(desc->goPtr, memory);
}

void
virtblocks_vm_description_set_disk(VirtBlocksVmDescription *desc,
                                   VirtBlocksDevicesDisk *disk)
{
    assert(desc != NULL);
    assert(disk != NULL);
    vm_description_set_disk(desc->goPtr, disk->goPtr);
}

VirtBlocksArray*
virtblocks_vm_description_get_qemu_commandline(VirtBlocksVmDescription *desc,
                                               VirtBlocksError **error)
{
    int goPtr;
    int ret;

    assert(desc != NULL);
    assert(error != NULL);

    ret = vm_description_get_qemu_commandline(desc->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return array_wrap(ret);
}
*/
import "C"
