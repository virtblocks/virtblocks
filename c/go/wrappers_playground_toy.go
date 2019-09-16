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

VirtBlocksPlaygroundToy*
virtblocks_playground_toy_new(const char *base,
                              VirtBlocksPlaygroundToyCallback filter,
                              void *data,
                              VirtBlocksPlaygroundToyCallbackDataFree dataFree)
{
    return playground_toy_wrap(playground_toy_new((char *) base, filter, data, dataFree));
}

void
virtblocks_playground_toy_free(VirtBlocksPlaygroundToy *toy)
{
    if (toy == NULL) return;
    playground_toy_free(toy->goPtr);
    free(toy);
}

char*
virtblocks_playground_toy_get_base(const VirtBlocksPlaygroundToy *toy)
{
    assert(toy != NULL);
    return playground_toy_get_base(toy->goPtr);
}

char*
virtblocks_playground_toy_run(const VirtBlocksPlaygroundToy *toy,
                              const char *ext,
                              VirtBlocksError **error)
{
    int goPtr;
    char *ret;

    assert(toy != NULL);
    assert(error != NULL);

    ret = playground_toy_run(toy->goPtr, (char *) ext, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return ret;
}
*/
import "C"
