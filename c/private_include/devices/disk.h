/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_DEVICES_DISK_H_
#define _PRIVATE_DEVICES_DISK_H_

struct _VirtBlocksDevicesDisk {
    unsigned int goPtr;
};

VirtBlocksDevicesDisk* devices_disk_wrap(unsigned int goPtr);

#endif
