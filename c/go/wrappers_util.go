// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

/*
#include "libvirtblocks_c_go.h"
#include "virtblocks.h"

int
virtblocks_util_build_file_name(char **file_name,
                                const char *base,
                                const char *ext)
{
    return util_build_file_name(file_name, (char *) base, (char *) ext);
}
*/
import "C"
