/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"

typedef struct _VirtBlocksVmDisk VirtBlocksVmDisk;

VirtBlocksVmDisk *virtblocks_vm_disk_new(void);
void virtblocks_vm_disk_free(VirtBlocksVmDisk *disk);

void virtblocks_vm_disk_set_filename(VirtBlocksVmDisk *disk,
                                     const char *filename);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmDisk,
                               virtblocks_vm_disk_free);
