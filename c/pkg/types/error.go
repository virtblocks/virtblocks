// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package types

import (
	"github.com/virtblocks/virtblocks/go/pkg/playground"
)

type ErrorDomain uint

const (
	GenericError ErrorDomain = iota
	PlaygroundToyError
)

type Error struct {
	native error
}

func NewError(err error) *Error {
	return &Error{native: err}
}

func (self *Error) Domain() ErrorDomain {
	var domain ErrorDomain

	switch self.native.(type) {
	case playground.ToyError:
		domain = ErrorDomain(PlaygroundToyError)
	default:
		domain = ErrorDomain(GenericError)
	}

	return domain
}

func (self *Error) Code() uint {
	var code uint

	switch self.native.(type) {
	case playground.ToyError:
		code = uint(self.native.(playground.ToyError))
	default:
		code = 0
	}

	return code
}

func (self *Error) Message() string {
	return self.native.Error()
}
