/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksDevicesDisk VirtBlocksDevicesDisk;

VirtBlocksDevicesDisk *virtblocks_devices_disk_new(void);
void virtblocks_devices_disk_free(VirtBlocksDevicesDisk *disk);

void virtblocks_devices_disk_set_filename(VirtBlocksDevicesDisk *disk,
                                          const char *filename);
VirtBlocksArray* virtblocks_devices_disk_get_qemu_commandline(VirtBlocksDevicesDisk *disk,
                                                              VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksDevicesDisk,
                               virtblocks_devices_disk_free);
