// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksVmMemballoon*
virtblocks_vm_memballoon_new()
{
    return vm_memballoon_wrap(vm_memballoon_new());
}

void
virtblocks_vm_memballoon_free(VirtBlocksVmMemballoon *memballoon)
{
    if (memballoon == NULL) return;
    vm_memballoon_free(memballoon->goPtr);
    free(memballoon);
}

void
virtblocks_vm_memballoon_set_model(VirtBlocksVmMemballoon *memballoon,
                                   VirtBlocksVmMemballoonModel model)
{
    assert(memballoon != NULL);
    vm_memballoon_set_model(memballoon->goPtr, model);
}

VirtBlocksVmMemballoonModel
virtblocks_vm_memballoon_get_model(const VirtBlocksVmMemballoon *memballoon)
{
    assert(memballoon != NULL);
    return vm_memballoon_get_model(memballoon->goPtr);
}

char*
virtblocks_vm_memballoon_to_string(const VirtBlocksVmMemballoon *memballoon)
{
    assert(memballoon != NULL);
    return vm_memballoon_to_string(memballoon->goPtr);
}
*/
import "C"
