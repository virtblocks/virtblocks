/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"

typedef struct _VirtBlocksArray VirtBlocksArray;

void virtblocks_array_free(VirtBlocksArray *array);

unsigned int virtblocks_array_get_length(const VirtBlocksArray *array);
const void* virtblocks_array_get_item(const VirtBlocksArray *array,
                                      unsigned int index);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksArray,
                               virtblocks_array_free);
