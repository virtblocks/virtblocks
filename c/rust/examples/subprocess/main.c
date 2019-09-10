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
    VIRTBLOCKS_AUTOPTR(VirtBlocksSubprocess) sub = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
    int fds[2];

    sub = virtblocks_subprocess_new("cat");
    virtblocks_subprocess_add_arg(sub, "/does-not-exist");
    virtblocks_subprocess_add_arg(sub, "-");

    if (socketpair(PF_LOCAL, SOCK_STREAM, 0, fds) < 0) {
        perror("failed to create socketpair:");
        exit(1);
    }
    virtblocks_subprocess_take_fd(sub, fds[1], 0);
    write(fds[0], "hello!", 6);
    close(fds[0]);

    if (!virtblocks_subprocess_spawn(sub, &err)) {
        print_err(err);
    } else {
        printf("PID: %u\n", virtblocks_subprocess_id(sub));
        sleep(1);
        if (!virtblocks_subprocess_kill(sub, &err)) {
            print_err(err);
        }

        int status = virtblocks_subprocess_wait(sub, &err);
        if (err) {
            print_err(err);
        }
        assert(status != 0);
    }

    return 0;
}
