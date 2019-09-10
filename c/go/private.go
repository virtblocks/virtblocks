// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksError*
error_wrap(int goPtr)
{
    VirtBlocksError *self = malloc(sizeof(VirtBlocksError));
    self->goPtr = goPtr;
    return self;
}

VirtBlocksPlaygroundToy*
playground_toy_wrap(int goPtr)
{
    VirtBlocksPlaygroundToy *self = malloc(sizeof(VirtBlocksPlaygroundToy));
    self->goPtr = goPtr;
    return self;
}

bool
playground_toy_callback_call(VirtBlocksPlaygroundToyCallback callback,
                             const VirtBlocksPlaygroundToy *toy,
                             const char *ext,
                             void *data)
{
    return callback(toy, ext, data);
}

void
playground_toy_callback_data_free_call(VirtBlocksPlaygroundToyCallbackDataFree dataFree,
                                       void *data)
{
    dataFree(data);
}

VirtBlocksDevicesMemballoon*
devices_memballoon_wrap(int goPtr)
{
    VirtBlocksDevicesMemballoon *self = malloc(sizeof(VirtBlocksDevicesMemballoon));
    self->goPtr = goPtr;
    return self;
}

VirtBlocksSubprocess*
subprocess_wrap(int goPtr)
{
    VirtBlocksSubprocess *self = malloc(sizeof(VirtBlocksSubprocess));
    self->goPtr = goPtr;
    return self;
}
*/
import "C"
