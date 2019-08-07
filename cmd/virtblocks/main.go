package virtblocks

import (
    "fmt"
    "github.com/virtblocks/virtblocks/pkg/devices"
    "github.com/virtblocks/virtblocks/pkg/util"
)

func main() {
    var file_name = util.BuildFileName("guest", ".xml")
    fmt.Println(file_name)

    var memballoon = devices.NewMemballoon()
    fmt.Println(memballoon.ToStr())

    memballoon.SetModel(devices.MEMBALLOON_MODEL_VIRTIO)
    fmt.Println(memballoon.ToStr())
}
