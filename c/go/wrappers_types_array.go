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

void
virtblocks_array_free(VirtBlocksArray *array)
{
    if (array == NULL) return;
    array_free(array->goPtr);
    free(array);
}

unsigned int
virtblocks_array_get_length(const VirtBlocksArray *array)
{
    assert(array != NULL);
    return array_get_length(array->goPtr);
}

const void*
virtblocks_array_get_item(const VirtBlocksArray *array,
                          unsigned int index)
{
    assert(array != NULL);
    return array_get_item(array->goPtr, index);
}
*/
import "C"
