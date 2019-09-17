/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#ifndef _PRIVATE_VM_DESCRIPTION_H_
#define _PRIVATE_VM_DESCRIPTION_H_

struct _VirtBlocksVmDescription {
    unsigned int goPtr;
};

VirtBlocksVmDescription* vm_description_wrap(unsigned int goPtr);

#endif
