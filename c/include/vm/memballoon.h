/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"

typedef enum {
  VIRTBLOCKS_VM_MEMBALLOON_MODEL_NONE = 0,
  VIRTBLOCKS_VM_MEMBALLOON_MODEL_VIRTIO,
  VIRTBLOCKS_VM_MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL,
  VIRTBLOCKS_VM_MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL,
} VirtBlocksVmMemballoonModel;

typedef struct _VirtBlocksVmMemballoon VirtBlocksVmMemballoon;

VirtBlocksVmMemballoon *virtblocks_vm_memballoon_new(void);
void virtblocks_vm_memballoon_free(VirtBlocksVmMemballoon *memballoon);

void virtblocks_vm_memballoon_set_model(VirtBlocksVmMemballoon *memballoon,
                                        VirtBlocksVmMemballoonModel model);
VirtBlocksVmMemballoonModel virtblocks_vm_memballoon_get_model(const VirtBlocksVmMemballoon *memballoon);

char *virtblocks_vm_memballoon_to_string(const VirtBlocksVmMemballoon *memballoon);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmMemballoon,
                               virtblocks_vm_memballoon_free);
