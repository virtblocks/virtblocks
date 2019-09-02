// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package errors

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/playground"
)

type ErrorDomain int

const (
	PlaygroundToyError ErrorDomain = iota
)

type Error struct {
	native error
}

func New(err error) *Error {
	return &Error{native: err}
}

func (self *Error) Domain() ErrorDomain {
	var domain ErrorDomain

	switch self.native.(type) {
	case playground.ToyError:
		domain = ErrorDomain(PlaygroundToyError)
	}

	return domain
}

func (self *Error) Code() int {
	var code int

	switch self.native.(type) {
	case playground.ToyError:
		code = int(self.native.(playground.ToyError))
	}

	return code
}

func (self *Error) Message() string {
	return self.native.Error()
}
