// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#pragma once

#include <virtblocks.h>

bool playground_toy_callback_trampoline_c(const VirtBlocksPlaygroundToy *toy,
                                          const char *ext,
                                          void *data);
int playground_toy_callback_trampoline_go(VirtBlocksPlaygroundToy *toy,
                                          char *ext,
                                          int callback);
