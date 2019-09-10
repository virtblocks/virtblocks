/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksSubprocess VirtBlocksSubprocess;

VirtBlocksSubprocess *virtblocks_subprocess_new(const char *cmd);
void virtblocks_subprocess_free(VirtBlocksSubprocess *sub);
int virtblocks_subprocess_spawn(VirtBlocksSubprocess *sub);
int virtblocks_subprocess_wait(VirtBlocksSubprocess *sub);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksSubprocess,
                               virtblocks_subprocess_free);
