// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package types

import (
	"unsafe"
)

type Array struct {
	items []unsafe.Pointer
}

func NewArray(items []unsafe.Pointer) *Array {
	return &Array{items: items}
}

func (self *Array) Length() uint {
	return uint(len(self.items))
}

func (self *Array) Item(index uint) unsafe.Pointer {
	return self.items[index]
}
