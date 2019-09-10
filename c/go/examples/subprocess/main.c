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
    VIRTBLOCKS_AUTOPTR(VirtBlocksSubprocess) sub = NULL;

    sub = virtblocks_subprocess_new("./test.sh");
    virtblocks_subprocess_spawn(sub);
    sleep(5);
    virtblocks_subprocess_wait(sub);

    return 0;
}
