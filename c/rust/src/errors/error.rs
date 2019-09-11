// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use libc;
use std::ffi::CString;
use std::io;
use std::os::raw::c_char;
use std::os::raw::c_uint;

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

fn convert_io_error(err: &io::Error) -> u32 {
    let ret = match err.raw_os_error() {
        Some(errno) => errno,
        None => {
            // The error was not originally created from an errno value,
            // so we can't obtain the value directly, but we can still
            // perform the opposite mapping as decode_error_kind() in
            // Rust's src/libstd/sys/unix/mod.rs
            match err.kind() {
                io::ErrorKind::ConnectionRefused => libc::ECONNREFUSED,
                io::ErrorKind::ConnectionReset => libc::ECONNRESET,
                io::ErrorKind::BrokenPipe => libc::EPIPE,
                io::ErrorKind::NotConnected => libc::ENOTCONN,
                io::ErrorKind::ConnectionAborted => libc::ECONNABORTED,
                io::ErrorKind::AddrNotAvailable => libc::EADDRNOTAVAIL,
                io::ErrorKind::AddrInUse => libc::EADDRINUSE,
                io::ErrorKind::NotFound => libc::ENOENT,
                io::ErrorKind::Interrupted => libc::EINTR,
                io::ErrorKind::InvalidInput => libc::EINVAL,
                io::ErrorKind::TimedOut => libc::ETIMEDOUT,
                io::ErrorKind::AlreadyExists => libc::EEXIST,
                // These two cases are special because two separate
                // errno values can map to the same ErrorKind: we just
                // pick one, knowing that if the error was originally
                // coming from the OS then we'd have executed the case
                // above and gotten the original errno back, and if not
                // then either one works
                io::ErrorKind::PermissionDenied => libc::EPERM,
                io::ErrorKind::WouldBlock => libc::EWOULDBLOCK,
                // If the error was entirely custom, then the only
                // sensible value we can return is zero
                io::ErrorKind::Other => 0,
                // If a new ErrorKind was introduced and we're not
                // handling it correctly, our best bet is to panic
                _ => panic!("Unrecognized io::ErrorKind"),
            }
        }
    };
    ret as u32
}

impl Error {
    fn domain(&self) -> ErrorDomain {
        match self {
            Error::PlaygroundToyError(_) => ErrorDomain::PlaygroundToyError,
            Error::IOError(_) => ErrorDomain::IOError,
        }
    }

    fn code(&self) -> u32 {
        match self {
            Error::PlaygroundToyError(e) => *e as u32,
            Error::IOError(e) => convert_io_error(e),
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
pub extern "C" fn virtblocks_error_get_code(c_error: *const Error) -> c_uint {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    rust_error.code() as c_uint
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_message(c_error: *const Error) -> *mut c_char {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    let rust_ret = rust_error.message();

    let tmp = CString::new(rust_ret).unwrap();
    unsafe { libc::strdup(tmp.as_ptr()) }
}
