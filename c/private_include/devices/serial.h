/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_DEVICES_SERIAL_H_
#define _PRIVATE_DEVICES_SERIAL_H_

struct _VirtBlocksDevicesSerial {
    unsigned int goPtr;
};

VirtBlocksDevicesSerial* devices_serial_wrap(unsigned int goPtr);

#endif
