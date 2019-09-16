/* Virt Blocks
 *
 * Copyright (C) 2019 Red Hat, Inc.
 *
 * This software is distributed under the terms of the MIT License.
 * See the LICENSE file in the top level directory for details.
 */

#include <stdio.h>
#include <string.h>
#include <virtblocks.h>

typedef struct {
    VirtBlocksPlaygroundToy *toy;
    const char *ext;
} Input;

static void
data_free(void *data) {
    free(data);
}

static bool
ext_is_baz(const VirtBlocksPlaygroundToy *toy,
           const char *ext,
           void *data)
{
    return !strcmp(ext, "baz");
}

static VirtBlocksPlaygroundToy*
create_baz()
{
    return virtblocks_playground_toy_new("baz", ext_is_baz, NULL, NULL);
}

static bool
base_and_ext_match(const VirtBlocksPlaygroundToy *toy,
                   const char *ext,
                   void *data)
{
    VIRTBLOCKS_AUTOFREE(char *) toy_base = NULL;

    toy_base = virtblocks_playground_toy_get_base(toy);

    return !strcmp(toy_base, ext);
}

static bool
base_is_foo(const VirtBlocksPlaygroundToy *toy,
            const char *ext,
            void *data)
{
    VIRTBLOCKS_AUTOFREE(char *) toy_base = NULL;
    char *base = (char *) data;

    toy_base = virtblocks_playground_toy_get_base(toy);

    return !strcmp(toy_base, base);
}

int
main(int argc,
     char **argv)
{
    const char *base = "foo";
    VIRTBLOCKS_AUTOPTR(VirtBlocksPlaygroundToy) foo_is_foo_toy = virtblocks_playground_toy_new("foo", base_is_foo, strdup(base), data_free);
    VIRTBLOCKS_AUTOPTR(VirtBlocksPlaygroundToy) foo_match_toy = virtblocks_playground_toy_new("foo", base_and_ext_match, NULL, NULL);
    VIRTBLOCKS_AUTOPTR(VirtBlocksPlaygroundToy) bar_is_foo_toy = virtblocks_playground_toy_new("bar", base_is_foo, strdup(base), data_free);
    VIRTBLOCKS_AUTOPTR(VirtBlocksPlaygroundToy) bar_match_toy = virtblocks_playground_toy_new("bar", base_and_ext_match, NULL, NULL);
    VIRTBLOCKS_AUTOPTR(VirtBlocksPlaygroundToy) baz_toy = create_baz();
    Input inputs[] = {
        { foo_is_foo_toy, "exe" },
        { foo_match_toy, "exe" },
        { bar_is_foo_toy, "bar" },
        { bar_match_toy, "bar" },
        { baz_toy, "baz" },
        { baz_toy, "quux" },
        { NULL, NULL },
    };
    unsigned int i = 0;

    while (true) {
        VIRTBLOCKS_AUTOPTR(VirtBlocksError) err = NULL;
        VIRTBLOCKS_AUTOFREE(char *) res = NULL;
        Input *input = &inputs[i];

        if (!input->toy || !input->ext)
            break;

        res = virtblocks_playground_toy_run(input->toy, input->ext, &err);
        if (res) {
            printf("Result: %s\n", res);
        } else {
            VIRTBLOCKS_AUTOFREE(char *) msg = NULL;

            msg = virtblocks_error_get_message(err);

            printf("Error: %s\n", msg);
        }

        i++;
    }

    return 0;
}
