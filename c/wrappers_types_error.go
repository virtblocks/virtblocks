// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

/*
#include "libvirtblocks.h"
#include "virtblocks.h"
#include "private.h"

void
virtblocks_error_free(VirtBlocksError *error)
{
    if (error == NULL)
        return;

    error_free(error->goPtr);

    free(error);
}

VirtBlocksErrorDomain
virtblocks_error_get_domain(const VirtBlocksError *error)
{
    if (error == NULL)
        return VIRTBLOCKS_ERROR_DOMAIN_GENERIC_ERROR;

    return error_get_domain(error->goPtr);
}

unsigned int
virtblocks_error_get_code(const VirtBlocksError *error)
{
    if (error == NULL)
        return 0;

    return error_get_code(error->goPtr);
}

const char *
virtblocks_error_get_message(const VirtBlocksError *error)
{
    if (error == NULL)
        return "<nil>";

    return error_get_message(error->goPtr);
}
*/
import "C"
