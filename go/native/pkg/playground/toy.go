// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package playground

const (
	CallbackFailed = iota
)

type ToyError int

func (self ToyError) Error() string {
	var ret string

	switch self {
	case CallbackFailed:
		ret = "Callback failed"
	}

	return ret
}

type ToyCallback func(*Toy, string) bool

type Toy struct {
	base   string
	filter ToyCallback
}

func NewToy(base string, filter ToyCallback) *Toy {
	return &Toy{base: base, filter: filter}
}

func (self Toy) Base() string {
	return self.base
}

func (self *Toy) Run(ext string) (string, error) {
	if self.filter(self, ext) {
		// If the user-provided filter succeeds, then build a string
		return self.base + "." + ext, nil
	} else {
		return "", ToyError(CallbackFailed)
	}
}
