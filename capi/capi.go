package main

import "C"

import (
	"github.com/virtblocks/virtblocks/pkg/devices"
	"github.com/virtblocks/virtblocks/pkg/util"
)

//export util_build_file_name
func util_build_file_name(cFileName **C.char, cBase *C.char, cExt *C.char) C.int {
	if cBase == nil || cExt == nil || cFileName == nil {
		return -1
	}
	var goBase = C.GoString(cBase)
	var goExt = C.GoString(cExt)

	var goRes = util.BuildFileName(goBase, goExt)

	*cFileName = C.CString(goRes)
	return 0
}

// Go objects are not allow to cross the language boundary, but we still
// need the C code to manipulate them. The way we achieve this is by
// having an array of objects on the Go side and passing "references" to
// them back and forth from the C side: each "reference" is simply the
// index of the object in the array, with 0 being a special "reference"
// that always points to nil
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
	return C.CString(goMemballoon.String())
}

func main() {}
