/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

typedef struct _VirtBlocksDevicesSerial VirtBlocksDevicesSerial;

VirtBlocksDevicesSerial *virtblocks_devices_serial_new(void);
void virtblocks_devices_serial_free(VirtBlocksDevicesSerial *serial);

void virtblocks_devices_serial_set_path(VirtBlocksDevicesSerial *serial,
                                        const char *path);
VirtBlocksArray* virtblocks_devices_serial_get_qemu_commandline(VirtBlocksDevicesSerial *serial,
                                                                VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksDevicesSerial,
                               virtblocks_devices_serial_free);
