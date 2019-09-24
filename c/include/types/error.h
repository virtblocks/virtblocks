/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

typedef enum {
    VIRTBLOCKS_ERROR_DOMAIN_GENERIC_ERROR = 0,
    VIRTBLOCKS_ERROR_DOMAIN_PLAYGROUND_TOY_ERROR,
} VirtBlocksErrorDomain;

typedef struct _VirtBlocksError VirtBlocksError;

void virtblocks_error_free(VirtBlocksError *error);

VirtBlocksErrorDomain virtblocks_error_get_domain(const VirtBlocksError *error);
unsigned int virtblocks_error_get_code(const VirtBlocksError *error);
const char *virtblocks_error_get_message(const VirtBlocksError *error);

#define VIRTBLOCKS_ERROR_FREE(_err) \
    do { \
        virtblocks_error_free(_err); \
        _err = NULL; \
    } while(false)

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksError,
                               virtblocks_error_free);
