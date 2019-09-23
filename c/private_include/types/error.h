/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_TYPES_ERROR_H_
#define _PRIVATE_TYPES_ERROR_H_

struct _VirtBlocksError {
    unsigned int goPtr;
};

VirtBlocksError* error_wrap(unsigned int goPtr);

#endif
