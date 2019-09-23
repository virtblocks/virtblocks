// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "virtblocks.h"
#include "private.h"

VirtBlocksPlaygroundToy*
playground_toy_wrap(unsigned int goPtr)
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
