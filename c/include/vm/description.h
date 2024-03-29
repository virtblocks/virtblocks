/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"
#include "../types/array.h"
#include "../types/error.h"
#include "disk.h"
#include "model.h"
#include "serial.h"

typedef struct _VirtBlocksVmDescription VirtBlocksVmDescription;

VirtBlocksVmDescription *virtblocks_vm_description_new(VirtBlocksVmModel model);
void virtblocks_vm_description_free(VirtBlocksVmDescription *description);

void virtblocks_vm_description_set_emulator(VirtBlocksVmDescription *description,
                                            const char *emulator);
void virtblocks_vm_description_set_cpus(VirtBlocksVmDescription *description,
                                        unsigned int cpus);
void virtblocks_vm_description_set_memory(VirtBlocksVmDescription *description,
                                          unsigned int memory);
void virtblocks_vm_description_set_disk_slots(VirtBlocksVmDescription *description,
                                              unsigned int slots);
void virtblocks_vm_description_set_disk_for_slot(VirtBlocksVmDescription *description,
                                                 VirtBlocksVmDisk *disk,
                                                 unsigned int slot);
void virtblocks_vm_description_set_serial(VirtBlocksVmDescription *description,
                                          VirtBlocksVmSerial *serial);
VirtBlocksArray* virtblocks_vm_description_get_qemu_commandline(VirtBlocksVmDescription *description,
                                                                VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmDescription,
                               virtblocks_vm_description_free);
