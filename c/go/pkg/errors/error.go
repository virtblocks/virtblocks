// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package errors

type DummyError int

const (
	DummyCode DummyError = iota
)

func (self DummyError) Error() string {
	var ret string

	switch self {
	case DummyCode:
		ret = "DummyMessage"
	}

	return ret
}

type ErrorDomain int

const (
	DummyDomain ErrorDomain = iota
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
	case DummyError:
		domain = ErrorDomain(DummyDomain)
	}

	return domain
}

func (self *Error) Code() int {
	var code int

	switch self.native.(type) {
	case DummyError:
		code = int(self.native.(DummyError))
	}

	return code
}

func (self *Error) Message() string {
	return self.native.Error()
}
