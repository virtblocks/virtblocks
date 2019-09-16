/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

struct _VirtBlocksPlaygroundToy {
    unsigned int goPtr;
};

VirtBlocksPlaygroundToy *playground_toy_wrap(unsigned int goPtr);

bool playground_toy_callback_call(VirtBlocksPlaygroundToyCallback callback,
                                  const VirtBlocksPlaygroundToy *toy,
                                  const char *ext,
                                  void *data);
void playground_toy_callback_data_free_call(VirtBlocksPlaygroundToyCallbackDataFree dataFree,
                                            void *data);
