/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"
#include "../types/array.h"
#include "../types/error.h"

typedef struct _VirtBlocksVmSerial VirtBlocksVmSerial;

VirtBlocksVmSerial *virtblocks_vm_serial_new(void);
void virtblocks_vm_serial_free(VirtBlocksVmSerial *serial);

void virtblocks_vm_serial_set_path(VirtBlocksVmSerial *serial,
                                   const char *path);
VirtBlocksArray* virtblocks_vm_serial_get_qemu_commandline(VirtBlocksVmSerial *serial,
                                                           VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmSerial,
                               virtblocks_vm_serial_free);
