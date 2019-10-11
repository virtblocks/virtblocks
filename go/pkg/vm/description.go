// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package vm

import (
	"errors"
	"strconv"
)

type Description struct {
	model    Model
	emulator string
	cpus     uint
	memory   uint
	disk     *Disk
	serial   *Serial
}

func NewDescription(model Model) *Description {
	return &Description{
		model:    model,
		emulator: "/usr/bin/qemu-system-x86_64",
	}
}

func (self *Description) Model() Model {
	return self.model
}

func (self *Description) SetEmulator(emulator string) *Description {
	self.emulator = emulator
	return self
}

func (self *Description) SetCpus(cpus uint) *Description {
	self.cpus = cpus
	return self
}

func (self *Description) SetMemory(memory uint) *Description {
	self.memory = memory
	return self
}

func (self *Description) SetDisk(disk *Disk) *Description {
	self.disk = disk
	return self
}

func (self *Description) SetSerial(serial *Serial) *Description {
	self.serial = serial
	return self
}

func (self *Description) validate() error {
	if self.cpus == 0 {
		return errors.New("number of CPUs not set")
	}

	if self.memory == 0 {
		return errors.New("memory size not set")
	}

	return nil
}

func (self *Description) QemuCommandLine() ([]string, error) {
	var ret = []string{
		self.emulator,
		"-nodefaults",
		"-no-user-config",
		"-display",
		"none",
		"-machine",
		"accel=kvm:tcg",
		"-smp",
		strconv.FormatUint(uint64(self.cpus), 10),
		"-m",
		strconv.FormatUint(uint64(self.memory), 10),
	}

	err := self.validate()
	if err != nil {
		return ret, err
	}

	switch self.model {
	case ModelLegacyV1:
		ret = append(ret, "-machine", "pc")
	case ModelModernV1:
		ret = append(ret, "-machine", "q35")
	}

	diskArgs, err := self.disk.qemuCommandLine(self.model)
	if err != nil {
		return ret, err
	}

	ret = append(ret, diskArgs...)

	serialArgs, err := self.serial.qemuCommandLine(self.model)
	if err != nil {
		return ret, err
	}

	ret = append(ret, serialArgs...)

	return ret, nil
}
