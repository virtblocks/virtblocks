/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#ifndef _PRIVATE_TYPES_ERROR_H_
#define _PRIVATE_TYPES_ERROR_H_

struct _VirtBlocksError {
    unsigned int goPtr;
};

VirtBlocksError* error_wrap(unsigned int goPtr);

#endif
