/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#include <unistd.h>
#include <stdio.h>
#include <virtblocks.h>
#include <assert.h>
#include <sys/types.h>
#include <sys/socket.h>

static void print_err(VirtBlocksError *err)
{
    VIRTBLOCKS_AUTOFREE(char *) msg = NULL;
    msg = virtblocks_error_get_message(err);
    printf("Error: %s\n", msg);
}

int
main(int argc,
     char **argv)
{
    VIRTBLOCKS_AUTOPTR(VirtBlocksCommand) command = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
    int fds[2];

    command = virtblocks_command_new("cat");
    virtblocks_command_add_arg(command, "/does-not-exist");
    virtblocks_command_add_arg(command, "-");

    if (socketpair(PF_LOCAL, SOCK_STREAM, 0, fds) < 0) {
        perror("failed to create socketpair:");
        exit(1);
    }
    virtblocks_command_take_fd(command, fds[1], 0);
    write(fds[0], "hello!", 6);
    close(fds[0]);

    if (!virtblocks_command_spawn(command, &err)) {
        print_err(err);
    } else {
        printf("PID: %u\n", virtblocks_command_id(command));
        sleep(1);
        if (!virtblocks_command_kill(command, &err)) {
            print_err(err);
        }

        int status = virtblocks_command_wait(command, &err);
        if (err) {
            print_err(err);
        }
        assert(status != 0);
    }

    return 0;
}
