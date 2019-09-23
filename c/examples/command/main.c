/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

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
