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

VirtBlocksSubprocess*
virtblocks_subprocess_new(const char *cmd)
{
    return subprocess_wrap(subprocess_new((char *)cmd));
}

void
virtblocks_subprocess_free(VirtBlocksSubprocess *sub)
{
    if (sub == NULL) return;
    subprocess_free(sub->goPtr);
    free(sub);
}

int
virtblocks_subprocess_spawn(VirtBlocksSubprocess *sub)
{
    assert(sub != NULL);
    return subprocess_spawn(sub->goPtr);
}

int
virtblocks_subprocess_wait(VirtBlocksSubprocess *sub)
{
    assert(sub != NULL);
    return subprocess_wait(sub->goPtr);
}
*/
import "C"
