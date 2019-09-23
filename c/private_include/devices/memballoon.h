/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_DEVICES_MEMBALLOON_H_
#define _PRIVATE_DEVICES_MEMBALLOON_H_

struct _VirtBlocksDevicesMemballoon {
    unsigned int goPtr;
};

VirtBlocksDevicesMemballoon* devices_memballoon_wrap(unsigned int goPtr);

#endif
