// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package vm

import (
	"errors"
	"github.com/virtblocks/virtblocks/go/native/pkg/devices"
	"strconv"
)

type Description struct {
	emulator string
	memory   int
	disk     *devices.Disk
}

func NewDescription() *Description {
	return &Description{
		emulator: "/usr/bin/qemu-system-x86_64",
	}
}

func (self *Description) SetEmulator(emulator string) *Description {
	self.emulator = emulator
	return self
}

func (self *Description) SetMemory(memory int) *Description {
	self.memory = memory
	return self
}

func (self *Description) SetDisk(disk *devices.Disk) *Description {
	self.disk = disk
	return self
}

func (self *Description) QemuCommandLine() ([]string, error) {
	var ret = []string{
		self.emulator,
		"-M",
		strconv.Itoa(self.memory),
	}

	if self.memory == 0 {
		return ret, errors.New("no memory size set")
	}

	disk, err := self.disk.QemuCommandLine()
	if err != nil {
		return ret, err
	}

	ret = append(ret, disk...)

	return ret, nil
}
