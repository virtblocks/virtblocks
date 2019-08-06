#include "libvirtblocks.h"

int virtblocks_util_build_file_name(char** p0, char* p1, char* p2);

int virtblocks_devices_memballoon_new();
void virtblocks_devices_memballoon_free(int p0);

void virtblocks_devices_memballoon_set_model(int p0, int p1);
int virtblocks_devices_memballoon_get_model(int p0);

char* virtblocks_devices_memballoon_to_str(int p0);
