#include "libvirtblocks.h"
#include "virtblocks.h"

int
virtblocks_util_build_file_name(char **file_name,
                                const char *base,
                                const char *ext)
{
    return util_build_file_name(file_name, (char *) base, (char *) ext);
}

VirtBlocksDevicesMemballoon
virtblocks_devices_memballoon_new()
{
    return devices_memballoon_new();
}
void
virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon memballoon)
{
    devices_memballoon_free(memballoon);
}

void
virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon memballoon,
                                        VirtBlocksDevicesMemballoonModel model)
{
    devices_memballoon_set_model(memballoon, model);
}

VirtBlocksDevicesMemballoonModel
virtblocks_devices_memballoon_get_model(VirtBlocksDevicesMemballoon memballoon)
{
    return devices_memballoon_get_model(memballoon);
}

char*
virtblocks_devices_memballoon_to_string(VirtBlocksDevicesMemballoon memballoon)
{
    return devices_memballoon_to_string(memballoon);
}
