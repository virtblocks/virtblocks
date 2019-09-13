/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#include <stdio.h>
#include <virtblocks.h>

int
main(int argc,
     char **argv)
{
    VIRTBLOCKS_AUTOPTR(VirtBlocksDevicesDisk) disk = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksVmDescription) desc = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksArray) args = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
    unsigned int i;

    disk = virtblocks_devices_disk_new();
    virtblocks_devices_disk_set_filename(disk, "test.qcow2");

    desc = virtblocks_vm_description_new();
    virtblocks_vm_description_set_memory(desc, 512);
    virtblocks_vm_description_set_disk(desc, disk);

    args = virtblocks_vm_description_get_qemu_commandline(desc, &err);
    if (err != NULL) {
        VIRTBLOCKS_AUTOFREE(char *) msg = NULL;

        msg = virtblocks_error_get_message(err);

        printf("domain=%d, code=%d, message=%s\n",
               virtblocks_error_get_domain(err),
               virtblocks_error_get_code(err),
               msg);
        return 1;
    }

    for (i = 0; i < virtblocks_array_get_length(args); i++) {
        const char *arg = (const char *) virtblocks_array_get_item(args, i);

        printf("%s\n", arg);
    }

    return 0;
}
