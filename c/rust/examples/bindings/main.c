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
    VIRTBLOCKS_AUTOPTR(VirtBlocksDevicesMemballoon) memballoon = NULL;
    VIRTBLOCKS_AUTOFREE(char *) file_name = NULL;
    VIRTBLOCKS_AUTOFREE(char *) model_before = NULL;
    VIRTBLOCKS_AUTOFREE(char *) model_after = NULL;

    virtblocks_util_build_file_name(&file_name, "guest", ".xml");
    if (file_name)
        printf("%s\n", file_name);

    memballoon = virtblocks_devices_memballoon_new();

    model_before = virtblocks_devices_memballoon_to_string(memballoon);
    if (model_before)
        printf("%s\n", model_before);

    virtblocks_devices_memballoon_set_model(memballoon,
                                            VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO);

    model_after = virtblocks_devices_memballoon_to_string(memballoon);
    if (model_after)
        printf("%s\n", model_after);

    return 0;
}
