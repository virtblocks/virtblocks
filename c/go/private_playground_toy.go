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
*/
import "C"
