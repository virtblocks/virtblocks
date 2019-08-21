int virtblocks_util_build_file_name(char **fileName,
                                    const char *base,
                                    const char *ext);

typedef enum {
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_NONE,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL,
  VIRTBLOCKS_DEVICES_MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL,
} VirtBlocksDevicesMemballoonModel;

typedef void* VirtBlocksDevicesMemballoon;

VirtBlocksDevicesMemballoon *virtblocks_devices_memballoon_new(void);
void virtblocks_devices_memballoon_free(VirtBlocksDevicesMemballoon *memballoon);

void virtblocks_devices_memballoon_set_model(VirtBlocksDevicesMemballoon *memballoon,
                                             VirtBlocksDevicesMemballoonModel model);
VirtBlocksDevicesMemballoonModel virtblocks_devices_memballoon_get_model(const VirtBlocksDevicesMemballoon *memballoon);

char *virtblocks_devices_memballoon_to_string(const VirtBlocksDevicesMemballoon *memballoon);
