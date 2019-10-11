/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#include <stdio.h>
#include <virtblocks.h>

int
main(int argc,
     char **argv)
{
    VIRTBLOCKS_AUTOPTR(VirtBlocksVmMemballoon) memballoon = NULL;
    VIRTBLOCKS_AUTOFREE(char *) model_before = NULL;
    VIRTBLOCKS_AUTOFREE(char *) model_after = NULL;

    memballoon = virtblocks_vm_memballoon_new();

    model_before = virtblocks_vm_memballoon_to_string(memballoon);
    if (model_before)
        printf("%s\n", model_before);

    virtblocks_vm_memballoon_set_model(memballoon,
                                       VIRTBLOCKS_VM_MEMBALLOON_MODEL_VIRTIO);

    model_after = virtblocks_vm_memballoon_to_string(memballoon);
    if (model_after)
        printf("%s\n", model_after);

    return 0;
}
