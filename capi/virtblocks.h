#include "libvirtblocks.h"

inline int
virtblocks_util_build_file_name(char** file_name,
                                const char* base,
                                const char* ext) {
    return util_build_file_name(file_name, (char*) base, (char*) ext);
}

int virtblocks_devices_memballoon_new();
void virtblocks_devices_memballoon_free(int p0);

void virtblocks_devices_memballoon_set_model(int p0, int p1);
int virtblocks_devices_memballoon_get_model(int p0);

char* virtblocks_devices_memballoon_to_str(int p0);
