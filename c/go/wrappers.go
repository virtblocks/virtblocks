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

int
virtblocks_util_build_file_name(char **file_name,
                                const char *base,
                                const char *ext)
{
    return util_build_file_name(file_name, (char *) base, (char *) ext);
}

struct _VirtBlocksDevicesMemballoon {
    int goPtr;
};

VirtBlocksDevicesMemballoon*
virtblocks_devices_memballoon_new()
{
    VirtBlocksDevicesMemballoon *self = malloc(sizeof(VirtBlocksDevicesMemballoon));
    self->goPtr = devices_memballoon_new();
    return self;
}

void
virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon *memballoon)
{
    devices_memballoon_free(memballoon->goPtr);
    free(memballoon);
}

void
virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon *memballoon,
                                        VirtBlocksDevicesMemballoonModel model)
{
    devices_memballoon_set_model(memballoon->goPtr, model);
}

VirtBlocksDevicesMemballoonModel
virtblocks_devices_memballoon_get_model(const VirtBlocksDevicesMemballoon *memballoon)
{
    return devices_memballoon_get_model(memballoon->goPtr);
}

char*
virtblocks_devices_memballoon_to_string(const VirtBlocksDevicesMemballoon *memballoon)
{
    return devices_memballoon_to_string(memballoon->goPtr);
}
*/
import "C"
