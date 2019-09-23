/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#ifndef _PRIVATE_PLAYGROUND_TOY_H_
#define _PRIVATE_PLAYGROUND_TOY_H_

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

#endif
