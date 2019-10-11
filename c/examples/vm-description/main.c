/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#include <stdio.h>
#include <virtblocks.h>

int
main(int argc,
     char **argv)
{
    VIRTBLOCKS_AUTOPTR(VirtBlocksVmDisk) disk = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksVmSerial) serial = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksVmDescription) desc = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksArray) args = NULL;
    VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
    unsigned int i;

    disk = virtblocks_vm_disk_new();
    virtblocks_vm_disk_set_filename(disk, "test.qcow2");

    serial = virtblocks_vm_serial_new();
    virtblocks_vm_serial_set_path(serial, "test.socket");

    desc = virtblocks_vm_description_new(VIRTBLOCKS_VM_MODEL_MODERN_V1);
    virtblocks_vm_description_set_cpus(desc, 1);
    virtblocks_vm_description_set_memory(desc, 512);
    virtblocks_vm_description_set_disk(desc, disk);
    virtblocks_vm_description_set_serial(desc, serial);

    args = virtblocks_vm_description_get_qemu_commandline(desc, &err);
    if (err != NULL) {
        printf("Error: %s\n", virtblocks_error_get_message(err));
        return 1;
    }

    for (i = 0; i < virtblocks_array_get_length(args); i++) {
        const char *arg = (const char *) virtblocks_array_get_item(args, i);

        printf("%s\n", arg);
    }

    return 0;
}
