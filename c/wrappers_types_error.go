// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

/*
#include "libvirtblocks.h"
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

unsigned int
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
*/
import "C"
