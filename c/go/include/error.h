/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#pragma once

typedef enum {
    VIRTBLOCKS_ERROR_DOMAIN_PLAYGROUND_TOY_ERROR,
} VirtBlocksErrorDomain;

typedef struct _VirtBlocksError VirtBlocksError;

void virtblocks_error_free(VirtBlocksError *error);

VirtBlocksErrorDomain virtblocks_error_get_domain(const VirtBlocksError *error);
uint32_t virtblocks_error_get_code(const VirtBlocksError *error);
char *virtblocks_error_get_message(const VirtBlocksError *error);

VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(VirtBlocksError,
                               virtblocks_error_free);
