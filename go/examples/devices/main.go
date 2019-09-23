// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/devices"
)

func main() {
	var memballoon = devices.NewMemballoon()
	fmt.Println(memballoon)

	memballoon.SetModel(devices.MemballoonModelVirtio)
	fmt.Println(memballoon)
}
