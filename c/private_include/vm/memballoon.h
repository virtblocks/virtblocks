/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_VM_MEMBALLOON_H_
#define _PRIVATE_VM_MEMBALLOON_H_

struct _VirtBlocksVmMemballoon {
    unsigned int goPtr;
};

VirtBlocksVmMemballoon* vm_memballoon_wrap(unsigned int goPtr);

#endif
