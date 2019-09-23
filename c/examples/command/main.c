/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#include <stdio.h>
#include <unistd.h>
#include <virtblocks.h>

int
main(int argc,
     char **argv)
{
    VIRTBLOCKS_AUTOPTR(VirtBlocksCommand) command = NULL;

    command = virtblocks_command_new("./test.sh");
    virtblocks_command_spawn(command);
    sleep(1);
    virtblocks_command_wait(command);

    return 0;
}
