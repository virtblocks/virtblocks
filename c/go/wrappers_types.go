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
virtblocks_error_free(VirtBlocksError *error)
{
    if (error == NULL) return;
    error_free(error->goPtr);
    free(error);
}

VirtBlocksErrorDomain
virtblocks_error_get_domain(const VirtBlocksError *error)
{
    assert(error != NULL);
    return error_get_domain(error->goPtr);
}

uint32_t
virtblocks_error_get_code(const VirtBlocksError *error)
{
    assert(error != NULL);
    return error_get_code(error->goPtr);
}

char *
virtblocks_error_get_message(const VirtBlocksError *error)
{
    assert(error != NULL);
    return error_get_message(error->goPtr);
}

void
virtblocks_array_free(VirtBlocksArray *array)
{
    if (array == NULL) return;
    array_free(array->goPtr);
    free(array);
}

uint32_t
virtblocks_array_get_length(const VirtBlocksArray *array)
{
    assert(array != NULL);
    return array_get_length(array->goPtr);
}

const void*
virtblocks_array_get_item(const VirtBlocksArray *array,
                          uint32_t index)
{
    assert(array != NULL);
    return array_get_item(array->goPtr, index);
}
*/
import "C"
