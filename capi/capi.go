package main

import "C"

import (
    "github.com/virtblocks/virtblocks/devices"
    "github.com/virtblocks/virtblocks/util"
)

//export util_build_file_name
func util_build_file_name(cFileName **C.char, cBase *C.char, cExt *C.char) C.int {
    if cBase == nil || cExt == nil || cFileName == nil {
        return -1;
    }
    var goBase = C.GoString(cBase)
    var goExt = C.GoString(cExt)

    var goRes = util.BuildFileName(goBase, goExt)

    *cFileName = C.CString(goRes)
    return 0
}

var memballoons = make([]*devices.Memballoon, 1)

//export devices_memballoon_new
func devices_memballoon_new() C.int {
    var goMemballoon = devices.NewMemballoon()
    memballoons = append(memballoons, goMemballoon)
    return C.int(len(memballoons) - 1)
}

//export devices_memballoon_free
func devices_memballoon_free(cMemballoon C.int) {
    memballoons[cMemballoon] = nil
}

//export devices_memballoon_set_model
func devices_memballoon_set_model(cMemballoon C.int, cModel C.int) {
    var goMemballoon = memballoons[cMemballoon]
    goMemballoon.SetModel(devices.MemballoonModel(cModel))
}

//export devices_memballoon_get_model
func devices_memballoon_get_model(cMemballoon C.int) C.int {
    var goMemballoon = memballoons[cMemballoon]
    return C.int(goMemballoon.GetModel())
}

//export devices_memballoon_to_str
func devices_memballoon_to_str(cMemballoon C.int) *C.char {
    var goMemballoon = memballoons[cMemballoon]
    return C.CString(goMemballoon.ToStr())
}

func main() {}
