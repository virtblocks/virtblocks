/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksSubprocess VirtBlocksSubprocess;

VirtBlocksSubprocess *virtblocks_subprocess_new(const char *prog);

void virtblocks_subprocess_free(VirtBlocksSubprocess *self);

void virtblocks_subprocess_add_arg(VirtBlocksSubprocess *self,
                                   const char *arg);

void virtblocks_subprocess_take_fd(VirtBlocksSubprocess *self,
                                   int sourcefd, int targetfd);

bool virtblocks_subprocess_spawn(VirtBlocksSubprocess *self,
                                 VirtBlocksError **error);

bool virtblocks_subprocess_kill(VirtBlocksSubprocess *self,
                                VirtBlocksError **error);

int virtblocks_subprocess_wait(VirtBlocksSubprocess *self,
                               VirtBlocksError **error);

uint32_t virtblocks_subprocess_id(VirtBlocksSubprocess *self);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksSubprocess,
                               virtblocks_subprocess_free);
