/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef struct _VirtBlocksVmDescription VirtBlocksVmDescription;

VirtBlocksVmDescription *virtblocks_vm_description_new(void);
void virtblocks_vm_description_free(VirtBlocksVmDescription *description);

void virtblocks_vm_description_set_emulator(VirtBlocksVmDescription *description,
                                            const char *emulator);
void virtblocks_vm_description_set_memory(VirtBlocksVmDescription *description,
                                          int memory);
void virtblocks_vm_description_set_disk(VirtBlocksVmDescription *description,
                                        VirtBlocksDevicesDisk *disk);
VirtBlocksArray* virtblocks_vm_description_get_qemu_commandline(VirtBlocksVmDescription *description,
                                                                VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmDescription,
                               virtblocks_vm_description_free);
