package main

import "C"

import (
    "gitlab.com/virtblocks/virtblocks/util"
)

//export virtblocks_util_build_file_name
func virtblocks_util_build_file_name(cFileName **C.char, cBase *C.char, cExt *C.char) C.int {
    if cBase == nil || cExt == nil || cFileName == nil {
        return -1;
    }
    var goBase = C.GoString(cBase)
    var goExt = C.GoString(cExt)

    var goRes = util.BuildFileName(goBase, goExt)

    *cFileName = C.CString(goRes)
    return 0
}

func main() {}
