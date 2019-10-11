/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_VM_DISK_H_
#define _PRIVATE_VM_DISK_H_

struct _VirtBlocksVmDisk {
    unsigned int goPtr;
};

VirtBlocksVmDisk* vm_disk_wrap(unsigned int goPtr);

#endif
