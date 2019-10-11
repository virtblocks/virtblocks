/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_VM_SERIAL_H_
#define _PRIVATE_VM_SERIAL_H_

struct _VirtBlocksVmSerial {
    unsigned int goPtr;
};

VirtBlocksVmSerial* vm_serial_wrap(unsigned int goPtr);

#endif
