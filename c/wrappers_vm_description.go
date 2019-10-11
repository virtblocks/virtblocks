// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmDescription*
virtblocks_vm_description_new(VirtBlocksVmModel model)
{
    return vm_description_wrap(vm_description_new(model));
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
virtblocks_vm_description_set_cpus(VirtBlocksVmDescription *desc,
                                   unsigned int cpus)
{
    assert(desc != NULL);
    vm_description_set_cpus(desc->goPtr, cpus);
}

void
virtblocks_vm_description_set_memory(VirtBlocksVmDescription *desc,
                                     unsigned int memory)
{
    assert(desc != NULL);
    vm_description_set_memory(desc->goPtr, memory);
}

void
virtblocks_vm_description_set_disk(VirtBlocksVmDescription *desc,
                                   VirtBlocksVmDisk *disk)
{
    assert(desc != NULL);
    assert(disk != NULL);
    vm_description_set_disk(desc->goPtr, disk->goPtr);
}

void
virtblocks_vm_description_set_serial(VirtBlocksVmDescription *desc,
                                     VirtBlocksVmSerial *serial)
{
    assert(desc != NULL);
    assert(serial != NULL);
    vm_description_set_serial(desc->goPtr, serial->goPtr);
}

VirtBlocksArray*
virtblocks_vm_description_get_qemu_commandline(VirtBlocksVmDescription *desc,
                                               VirtBlocksError **error)
{
    unsigned int goPtr;
    unsigned int ret;

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
