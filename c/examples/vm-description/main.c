/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

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
        printf("Error: %s\n", VIRTBLOCKS_ERROR_GET_MESSAGE(err));
        return 1;
    }

    for (i = 0; i < virtblocks_array_get_length(args); i++) {
        const char *arg = (const char *) virtblocks_array_get_item(args, i);

        printf("%s\n", arg);
    }

    return 0;
}
