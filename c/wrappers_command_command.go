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

void
virtblocks_command_spawn(VirtBlocksCommand *command,
                         VirtBlocksError **error)
{
    unsigned int goPtr;

    assert(command != NULL);
    assert(error != NULL);

    command_spawn(command->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }
}

unsigned int
virtblocks_command_get_id(VirtBlocksCommand *command,
                          VirtBlocksError **error)
{
    unsigned int goPtr;
    unsigned int ret;

    assert(command != NULL);
    assert(error != NULL);

    ret = command_get_id(command->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return ret;
}

void
virtblocks_command_kill(VirtBlocksCommand *command,
                        VirtBlocksError **error)
{
    unsigned int goPtr;

    assert(command != NULL);
    assert(error != NULL);

    command_kill(command->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }
}

int
virtblocks_command_wait(VirtBlocksCommand *command,
                        VirtBlocksError **error)
{
    unsigned int goPtr;
    int ret;

    assert(command != NULL);
    assert(error != NULL);

    ret = command_wait(command->goPtr, &goPtr);

    if (goPtr) {
        *error = error_wrap(goPtr);
    } else {
        *error = NULL;
    }

    return ret;
}
*/
import "C"
