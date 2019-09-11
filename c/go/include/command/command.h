/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksCommand VirtBlocksCommand;

VirtBlocksCommand *virtblocks_command_new(const char *prog);
void virtblocks_command_free(VirtBlocksCommand *command);
int virtblocks_command_spawn(VirtBlocksCommand *command);
int virtblocks_command_wait(VirtBlocksCommand *command);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksCommand,
                               virtblocks_command_free);
