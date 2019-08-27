/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

int virtblocks_util_build_file_name(char **file_name,
                                    const char *base,
                                    const char *ext);

typedef enum {
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_NONE,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL,
} VirtBlocksDevicesMemballoonModel;

typedef int VirtBlocksDevicesMemballoon;

VirtBlocksDevicesMemballoon virtblocks_devices_memballoon_new();
void virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon memballoon);

void virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon memballoon,
                                             VirtBlocksDevicesMemballoonModel model);
VirtBlocksDevicesMemballoonModel virtblocks_devices_memballoon_get_model(VirtBlocksDevicesMemballoon memballoon);

char* virtblocks_devices_memballoon_to_string(VirtBlocksDevicesMemballoon memballoon);
