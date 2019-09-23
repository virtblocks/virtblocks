/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

typedef struct _VirtBlocksCommand VirtBlocksCommand;

VirtBlocksCommand *virtblocks_command_new(const char *prog);
void virtblocks_command_free(VirtBlocksCommand *command);
int virtblocks_command_spawn(VirtBlocksCommand *command);
int virtblocks_command_wait(VirtBlocksCommand *command);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksCommand,
                               virtblocks_command_free);
