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
    VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
    int rc;

    command = virtblocks_command_new("./test.sh");
    virtblocks_command_spawn(command);
    sleep(1);

    rc = virtblocks_command_wait(command, &err);
    if (err != NULL) {
        VIRTBLOCKS_AUTOFREE(char *) msg = NULL;
        msg = virtblocks_error_get_message(err);
        printf("Command returned: rc=%d, err=%s\n", rc, msg);
    } else {
        printf("Command returned: rc=%d, err=<nil>\n", rc);
    }

    return 0;
}
