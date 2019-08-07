package main

import (
    "fmt"
    "github.com/virtblocks/virtblocks/devices"
    "github.com/virtblocks/virtblocks/util"
)

func main() {
    var file_name = util.BuildFileName("guest", ".xml")
    fmt.Println(file_name)

    var memballoon = devices.NewMemballoon()
    fmt.Println(memballoon.ToStr())

    memballoon.SetModel(devices.MEMBALLOON_MODEL_VIRTIO)
    fmt.Println(memballoon.ToStr())
}
