// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package vm

import (
	"errors"
	"fmt"
)

type Disk struct {
	filename string
}

func NewDisk() *Disk {
	return &Disk{}
}

func (self *Disk) SetFilename(filename string) *Disk {
	self.filename = filename
	return self
}

func (self *Disk) validate() error {
	if self.filename == "" {
		return errors.New("filename not set")
	}

	return nil
}

func (self *Disk) qemuCommandLine(model Model, slot uint) ([]string, error) {
	var ret = make([]string, 0)

	err := self.validate()
	if err != nil {
		return ret, err
	}

	var addr string
	switch model {
	case ModelLegacyV1:
		addr = fmt.Sprintf("bus=pci.0,addr=%d", slot+3)
	case ModelModernV1:
		addr = fmt.Sprintf("bus=diskslot%d,addr=0", slot)
	}

	ret = append(ret,
		"-drive",
		fmt.Sprintf("file=%s,format=qcow2,if=none,id=drive%d", self.filename, slot),
		"-device",
		fmt.Sprintf("virtio-blk-pci,drive=drive%d,%s,id=disk%d", slot, addr, slot),
	)

	return ret, nil
}
