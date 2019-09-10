// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use std::ffi::CString;
use std::io;
use std::os::raw::c_char;
use std::os::raw::c_int;

use virtblocks::playground;

#[repr(u32)]
#[derive(Copy, Clone, Debug)]
pub enum ErrorDomain {
    PlaygroundToyError,
    IOError,
}

#[derive(Debug)]
pub enum Error {
    PlaygroundToyError(playground::ToyError),
    IOError(io::Error),
}

impl From<playground::ToyError> for Error {
    fn from(err: playground::ToyError) -> Self {
        Error::PlaygroundToyError(err)
    }
}

impl From<io::Error> for Error {
    fn from(err: io::Error) -> Self {
        Error::IOError(err)
    }
}

impl Error {
    fn domain(&self) -> ErrorDomain {
        match self {
            Error::PlaygroundToyError(_) => ErrorDomain::PlaygroundToyError,
            Error::IOError(_) => ErrorDomain::IOError,
        }
    }

    fn code(&self) -> i32 {
        match self {
            Error::PlaygroundToyError(e) => *e as i32,
            Error::IOError(e) => e.raw_os_error().unwrap_or(-1),
        }
    }

    fn message(&self) -> String {
        match self {
            Error::PlaygroundToyError(e) => e.to_string(),
            Error::IOError(e) => e.to_string(),
        }
    }
}

#[no_mangle]
pub extern "C" fn virtblocks_error_free(c_error: *mut Error) {
    if c_error.is_null() {
        return;
    }
    let _rust_error = unsafe { Box::from_raw(c_error) };
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_domain(c_error: *const Error) -> ErrorDomain {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    rust_error.domain()
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_code(c_error: *const Error) -> c_int {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    rust_error.code() as c_int
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_message(c_error: *const Error) -> *mut c_char {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    let rust_ret = rust_error.message();

    let tmp = CString::new(rust_ret).unwrap();
    unsafe { libc::strdup(tmp.as_ptr()) }
}
