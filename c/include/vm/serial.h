/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"

typedef struct _VirtBlocksVmSerial VirtBlocksVmSerial;

VirtBlocksVmSerial *virtblocks_vm_serial_new(void);
void virtblocks_vm_serial_free(VirtBlocksVmSerial *serial);

void virtblocks_vm_serial_set_path(VirtBlocksVmSerial *serial,
                                   const char *path);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksVmSerial,
                               virtblocks_vm_serial_free);
