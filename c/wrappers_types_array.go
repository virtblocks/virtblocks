// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
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
