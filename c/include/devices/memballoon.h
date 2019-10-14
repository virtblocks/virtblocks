/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"

typedef enum {
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_NONE = 0,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL,
} VirtBlocksDevicesMemballoonModel;

typedef struct _VirtBlocksDevicesMemballoon VirtBlocksDevicesMemballoon;

VirtBlocksDevicesMemballoon *virtblocks_devices_memballoon_new(void);
void virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon *memballoon);

void virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon *memballoon,
                                             VirtBlocksDevicesMemballoonModel model);
VirtBlocksDevicesMemballoonModel virtblocks_devices_memballoon_get_model(const VirtBlocksDevicesMemballoon *memballoon);

char *virtblocks_devices_memballoon_to_string(const VirtBlocksDevicesMemballoon *memballoon);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksDevicesMemballoon,
                               virtblocks_devices_memballoon_free);
