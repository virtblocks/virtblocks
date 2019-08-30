/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#include <stdio.h>
#include <stdlib.h>
#include <virtblocks.h>

int
main(int argc,
     char **argv)
{
    VirtBlocksDevicesMemballoon *memballoon = NULL;
    char *str = NULL;

    virtblocks_util_build_file_name(&str, "guest", ".xml");
    if (str) {
        printf("%s\n", str);
        free(str);
    }

    memballoon = virtblocks_devices_memballoon_new();

    str = virtblocks_devices_memballoon_to_string(memballoon);
    if (str) {
        printf("%s\n", str);
        free(str);
    }

    virtblocks_devices_memballoon_set_model(memballoon,
                                            VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO);

    str = virtblocks_devices_memballoon_to_string(memballoon);
    if (str) {
        printf("%s\n", str);
        free(str);
    }

    virtblocks_devices_memballoon_free(memballoon);

    return 0;
}
