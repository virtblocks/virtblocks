/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_VM_DESCRIPTION_H_
#define _PRIVATE_VM_DESCRIPTION_H_

struct _VirtBlocksVmDescription {
    unsigned int goPtr;
};

VirtBlocksVmDescription* vm_description_wrap(unsigned int goPtr);

#endif
