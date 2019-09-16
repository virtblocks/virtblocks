/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

struct _VirtBlocksDevicesDisk {
    unsigned int goPtr;
};

VirtBlocksDevicesDisk* devices_disk_wrap(unsigned int goPtr);
