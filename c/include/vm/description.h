/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

typedef struct _VirtBlocksVmDescription VirtBlocksVmDescription;

VirtBlocksVmDescription *virtblocks_vm_description_new(void);
void virtblocks_vm_description_free(VirtBlocksVmDescription *description);

void virtblocks_vm_description_set_emulator(VirtBlocksVmDescription *description,
                                            const char *emulator);
void virtblocks_vm_description_set_memory(VirtBlocksVmDescription *description,
                                          unsigned int memory);
void virtblocks_vm_description_set_disk(VirtBlocksVmDescription *description,
                                        VirtBlocksDevicesDisk *disk);
VirtBlocksArray* virtblocks_vm_description_get_qemu_commandline(VirtBlocksVmDescription *description,
                                                                VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmDescription,
                               virtblocks_vm_description_free);
