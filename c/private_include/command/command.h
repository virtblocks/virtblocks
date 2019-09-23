/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_COMMAND_COMMAND_H_
#define _PRIVATE_COMMAND_COMMAND_H_

struct _VirtBlocksCommand {
    unsigned int goPtr;
};

VirtBlocksCommand* command_wrap(unsigned int goPtr);

#endif
