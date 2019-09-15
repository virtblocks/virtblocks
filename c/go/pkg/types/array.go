// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

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
