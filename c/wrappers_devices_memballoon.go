// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksDevicesMemballoon*
virtblocks_devices_memballoon_new()
{
    return devices_memballoon_wrap(devices_memballoon_new());
}

void
virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon *memballoon)
{
    if (memballoon == NULL) return;
    devices_memballoon_free(memballoon->goPtr);
    free(memballoon);
}

void
virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon *memballoon,
                                        VirtBlocksDevicesMemballoonModel model)
{
    assert(memballoon != NULL);
    devices_memballoon_set_model(memballoon->goPtr, model);
}

VirtBlocksDevicesMemballoonModel
virtblocks_devices_memballoon_get_model(const VirtBlocksDevicesMemballoon *memballoon)
{
    assert(memballoon != NULL);
    return devices_memballoon_get_model(memballoon->goPtr);
}

char*
virtblocks_devices_memballoon_to_string(const VirtBlocksDevicesMemballoon *memballoon)
{
    assert(memballoon != NULL);
    return devices_memballoon_to_string(memballoon->goPtr);
}
*/
import "C"
