package main

import (
    "fmt"
    "gitlab.com/virtblocks/virtblocks/devices"
    "gitlab.com/virtblocks/virtblocks/util"
)

func main() {
    var file_name = util.BuildFileName("guest", ".xml")
    fmt.Println(file_name)

    var memballoon = devices.NewMemballoon()
    fmt.Println(memballoon.ToStr())

    memballoon.SetModel(devices.MEMBALLOON_MODEL_VIRTIO)
    fmt.Println(memballoon.ToStr())
}
