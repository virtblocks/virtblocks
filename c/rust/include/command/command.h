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

void virtblocks_command_add_arg(VirtBlocksCommand *command,
                                const char *arg);

void virtblocks_command_take_fd(VirtBlocksCommand *command,
                                int sourcefd, int targetfd);

bool virtblocks_command_spawn(VirtBlocksCommand *command,
                              VirtBlocksError **error);

bool virtblocks_command_kill(VirtBlocksCommand *command,
                             VirtBlocksError **error);

int virtblocks_command_wait(VirtBlocksCommand *command,
                            VirtBlocksError **error);

uint32_t virtblocks_command_id(VirtBlocksCommand *command);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksCommand,
                               virtblocks_command_free);
