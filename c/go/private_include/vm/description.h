/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

struct _VirtBlocksVmDescription {
    int goPtr;
};

VirtBlocksVmDescription* vm_description_wrap(int goPtr);
