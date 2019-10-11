// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
)

func main() {
	var memballoon = vm.NewMemballoon()
	fmt.Println(memballoon)

	memballoon.SetModel(vm.MemballoonModelVirtio)
	fmt.Println(memballoon)
}
