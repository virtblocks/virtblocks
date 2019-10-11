/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"
#include "../types/array.h"
#include "../types/error.h"

typedef struct _VirtBlocksVmDisk VirtBlocksVmDisk;

VirtBlocksVmDisk *virtblocks_vm_disk_new(void);
void virtblocks_vm_disk_free(VirtBlocksVmDisk *disk);

void virtblocks_vm_disk_set_filename(VirtBlocksVmDisk *disk,
                                     const char *filename);
VirtBlocksArray* virtblocks_vm_disk_get_qemu_commandline(VirtBlocksVmDisk *disk,
                                                         VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmDisk,
                               virtblocks_vm_disk_free);
