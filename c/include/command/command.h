/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

typedef struct _VirtBlocksCommand VirtBlocksCommand;

VirtBlocksCommand *virtblocks_command_new(const char *prog);
void virtblocks_command_free(VirtBlocksCommand *command);

void virtblocks_command_add_arg(VirtBlocksCommand *command,
                                const char *arg);
void virtblocks_command_spawn(VirtBlocksCommand *command,
                              VirtBlocksError **error);
unsigned int virtblocks_command_get_id(VirtBlocksCommand *command,
                                       VirtBlocksError **error);
void virtblocks_command_kill(VirtBlocksCommand *command,
                             VirtBlocksError **error);
int virtblocks_command_wait(VirtBlocksCommand *command,
                            VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksCommand,
                               virtblocks_command_free);
