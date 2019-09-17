// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package devices

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

func (self *Disk) QemuCommandLine() ([]string, error) {
	var ret = make([]string, 0)

	if self.filename == "" {
		return ret, errors.New("filename not set")
	}

	ret = append(ret,
		"-drive",
		fmt.Sprintf("file=%s,format=qcow2,id=disk0", self.filename),
		"-device",
		"virtio-blk-pci,drive=disk0",
	)

	return ret, nil
}
