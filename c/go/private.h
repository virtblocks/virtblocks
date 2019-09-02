/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

#include "virtblocks.h"

struct _VirtBlocksError {
    int goPtr;
};

VirtBlocksError* error_wrap(int goPtr);

struct _VirtBlocksDevicesMemballoon {
    int goPtr;
};

VirtBlocksDevicesMemballoon* devices_memballoon_wrap(int goPtr);