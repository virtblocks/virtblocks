package main

import (
    "fmt"
    "gitlab.com/virtblocks/virtblocks/util"
)

func main() {
    var file_name = util.BuildFileName("guest", ".xml")
    fmt.Println(file_name)
}
