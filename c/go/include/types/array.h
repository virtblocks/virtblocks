/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksArray VirtBlocksArray;

void virtblocks_array_free(VirtBlocksArray *array);

unsigned int virtblocks_array_get_length(const VirtBlocksArray *array);
const void* virtblocks_array_get_item(const VirtBlocksArray *array,
                                      unsigned int index);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksArray,
                               virtblocks_array_free);
