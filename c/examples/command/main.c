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
    unsigned int id;
    int rc;

    command = virtblocks_command_new("./test.sh");

    virtblocks_command_spawn(command, &err);
    printf("Command started: err=%s\n", virtblocks_error_get_message(err));
    VIRTBLOCKS_ERROR_FREE(err);

    id = virtblocks_command_get_id(command, &err);
    printf("Command ID: id=%d, err=%s\n", id, virtblocks_error_get_message(err));
    VIRTBLOCKS_ERROR_FREE(err);

    sleep(1);

    rc = virtblocks_command_wait(command, &err);
    printf("Command returned: rc=%d, err=%s\n", rc, virtblocks_error_get_message(err));
    VIRTBLOCKS_ERROR_FREE(err);

    return 0;
}
