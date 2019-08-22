package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/rust/pkg/devices"
	"github.com/virtblocks/virtblocks/go/rust/pkg/util"
)

func main() {
	var file_name = util.BuildFileName("guest", ".xml")
	fmt.Println(file_name)

	var memballoon = devices.NewMemballoon()
	defer memballoon.Free()

	fmt.Println(memballoon)

	memballoon.SetModel(devices.MemballoonModelVirtio)
	fmt.Println(memballoon)
}
