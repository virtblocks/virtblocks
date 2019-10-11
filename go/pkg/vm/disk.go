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

func (self *Disk) qemuCommandLine(model Model) ([]string, error) {
	var ret = make([]string, 0)

	err := self.validate()
	if err != nil {
		return ret, err
	}

	ret = append(ret,
		"-drive",
		fmt.Sprintf("file=%s,format=qcow2,if=none,id=disk0", self.filename),
		"-device",
		"virtio-blk-pci,drive=disk0",
	)

	return ret, nil
}
