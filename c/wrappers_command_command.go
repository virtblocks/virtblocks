// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

VirtBlocksCommand*
virtblocks_command_new(const char *prog)
{
    return command_wrap(command_new((char *) prog));
}

void
virtblocks_command_free(VirtBlocksCommand *command)
{
    if (command == NULL) return;
    command_free(command->goPtr);
    free(command);
}

int
virtblocks_command_spawn(VirtBlocksCommand *command)
{
    assert(command != NULL);
    return command_spawn(command->goPtr);
}

int
virtblocks_command_wait(VirtBlocksCommand *command)
{
    assert(command != NULL);
    return command_wait(command->goPtr);
}
*/
import "C"
