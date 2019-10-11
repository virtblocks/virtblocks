// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package vm

import (
	"errors"
	"fmt"
)

type Serial struct {
	path string
}

func NewSerial() *Serial {
	return &Serial{}
}

func (self *Serial) SetPath(path string) *Serial {
	self.path = path
	return self
}

func (self *Serial) validate() error {
	if self.path == "" {
		return errors.New("path not set")
	}

	return nil
}

func (self *Serial) QemuCommandLine() ([]string, error) {
	var ret = make([]string, 0)

	err := self.validate()
	if err != nil {
		return ret, err
	}

	ret = append(ret,
		"-chardev",
		fmt.Sprintf("socket,path=%s,server,nowait,id=serial0", self.path),
		"-device",
		"isa-serial,chardev=serial0",
	)

	return ret, nil
}
