/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_TYPES_ARRAY_H_
#define _PRIVATE_TYPES_ARRAY_H_

struct _VirtBlocksArray {
    unsigned int goPtr;
};

VirtBlocksArray* array_wrap(unsigned int goPtr);

#endif
