/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#include "../alloc.h"
#include "../types/error.h"

typedef enum {
    VIRTBLOCKS_PLAYGROUND_TOY_ERROR_CALLBACK_FAILED = 0,
} VirtBlocksPlaygroundToyError;

typedef struct _VirtBlocksPlaygroundToy VirtBlocksPlaygroundToy;

typedef bool (*VirtBlocksPlaygroundToyCallback)(const VirtBlocksPlaygroundToy *toy,
                                                const char *ext,
                                                void *data);
typedef void (*VirtBlocksPlaygroundToyCallbackDataFree)(void *data);

VirtBlocksPlaygroundToy *virtblocks_playground_toy_new(const char *base,
                                                       VirtBlocksPlaygroundToyCallback filter,
                                                       void *data,
                                                       VirtBlocksPlaygroundToyCallbackDataFree dataFree);
void virtblocks_playground_toy_free(VirtBlocksPlaygroundToy *toy);

char *virtblocks_playground_toy_get_base(const VirtBlocksPlaygroundToy *toy);

char *virtblocks_playground_toy_run(const VirtBlocksPlaygroundToy *toy,
                                    const char *ext,
                                    VirtBlocksError **error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksPlaygroundToy,
                               virtblocks_playground_toy_free);
