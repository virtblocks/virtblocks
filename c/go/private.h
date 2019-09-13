/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

#include "virtblocks.h"

struct _VirtBlocksError {
    int goPtr;
};

VirtBlocksError* error_wrap(int goPtr);

struct _VirtBlocksArray {
    int goPtr;
};

VirtBlocksArray* array_wrap(int goPtr);

struct _VirtBlocksPlaygroundToy {
    int goPtr;
};

VirtBlocksPlaygroundToy *playground_toy_wrap(int goPtr);

bool playground_toy_callback_call(VirtBlocksPlaygroundToyCallback callback,
                                  const VirtBlocksPlaygroundToy *toy,
                                  const char *ext,
                                  void *data);
void playground_toy_callback_data_free_call(VirtBlocksPlaygroundToyCallbackDataFree dataFree,
                                            void *data);

struct _VirtBlocksDevicesMemballoon {
    int goPtr;
};

VirtBlocksDevicesMemballoon* devices_memballoon_wrap(int goPtr);

struct _VirtBlocksDevicesDisk {
    int goPtr;
};

VirtBlocksDevicesDisk* devices_disk_wrap(int goPtr);

struct _VirtBlocksCommand {
    int goPtr;
};

VirtBlocksCommand* command_wrap(int goPtr);

struct _VirtBlocksVmDescription {
    int goPtr;
};

VirtBlocksVmDescription* vm_description_wrap(int goPtr);
