/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#ifndef _PRIVATE_COMMAND_COMMAND_H_
#define _PRIVATE_COMMAND_COMMAND_H_

struct _VirtBlocksCommand {
    unsigned int goPtr;
};

VirtBlocksCommand* command_wrap(unsigned int goPtr);

#endif
