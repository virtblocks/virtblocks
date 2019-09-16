/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

struct _VirtBlocksDevicesDisk {
    int goPtr;
};

VirtBlocksDevicesDisk* devices_disk_wrap(int goPtr);
