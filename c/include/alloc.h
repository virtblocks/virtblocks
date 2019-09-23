/* Copyright (C) 2019 Red Hat, Inc. */
/* SPDX-License-Identifier: LGPL-2.1-or-later */

#pragma once

#define VIRTBLOCKS_FREE(_ptr) \
    do { \
        if (_ptr) \
            free((void *) _ptr); \
        _ptr = NULL; \
    } while(false)

static inline void virtblocks_autofree_helper(void *ptr)
{
    if (*(void **)ptr)
        free(*(void **)ptr);
    *(void **)ptr = NULL;
}

#define VIRTBLOCKS_AUTOFREE(_type) \
    __attribute__((cleanup(virtblocks_autofree_helper))) _type

#define VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type) _type##VirtBlocksAutoPtrFunc

#define VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(_type, _func) \
    static inline void VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type)(_type **_ptr) \
    { \
        if (*_ptr) \
            (_func)(*_ptr); \
        *_ptr = NULL; \
    }

#define VIRTBLOCKS_AUTOPTR(_type) \
    __attribute__((cleanup(VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type)))) _type *
