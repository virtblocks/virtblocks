#include "libvirtblocks.h"

int
virtblocks_util_build_file_name(char **file_name,
                                const char *base,
                                const char *ext)
{
    return util_build_file_name(file_name, (char *) base, (char *) ext);
}
