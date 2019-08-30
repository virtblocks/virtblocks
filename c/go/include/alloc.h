/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

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

#define VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type) _type##AutoPtrFree

#define VIRTBLOCKS_DEFINE_AUTOPTR_FUNC(_type, _func) \
    static inline void VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type)(_type *_ptr) \
    { \
        if (*_ptr) \
            (_func)(*_ptr); \
        *_ptr = 0; \
    }

#define VIRTBLOCKS_AUTOPTR(_type) \
    __attribute__((cleanup(VIRTBLOCKS_AUTOPTR_FUNC_NAME(_type)))) _type
