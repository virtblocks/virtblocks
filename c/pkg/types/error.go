// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package types

import "C"

import (
	"github.com/virtblocks/virtblocks/go/pkg/playground"
)

type ErrorDomain uint

const (
	GenericError ErrorDomain = iota
	PlaygroundToyError
)

type Error struct {
	domain  C.uint
	code    C.uint
	message *C.char
}

func NewError(err error) *Error {
	return &Error{
		domain:  getDomain(err),
		code:    getCode(err),
		message: getMessage(err),
	}
}

func (self *Error) Domain() C.uint {
	return self.domain
}

func (self *Error) Code() C.uint {
	return self.code
}

func (self *Error) Message() *C.char {
	return self.message
}

func getDomain(err error) C.uint {
	var domain ErrorDomain

	switch err.(type) {
	case playground.ToyError:
		domain = ErrorDomain(PlaygroundToyError)
	default:
		domain = ErrorDomain(GenericError)
	}

	return C.uint(domain)
}

func getCode(err error) C.uint {
	var code uint

	switch err.(type) {
	case playground.ToyError:
		code = uint(err.(playground.ToyError))
	default:
		code = 0
	}

	return C.uint(code)
}

func getMessage(err error) *C.char {
	return C.CString(err.Error())
}
