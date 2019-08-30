// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use std::ffi::CString;
use std::os::raw::c_char;
use std::os::raw::c_uint;

use virtblocks::playground;

#[repr(u32)]
#[derive(Copy, Clone, Debug)]
pub enum VirtBlocksErrorDomain {
    PlaygroundToyError,
}

#[derive(Copy, Clone, Debug)]
pub enum VirtBlocksError {
    PlaygroundToyError(playground::ToyError),
}

impl From<playground::ToyError> for VirtBlocksError {
    fn from(err: playground::ToyError) -> Self {
        Self::PlaygroundToyError(err)
    }
}

impl VirtBlocksError {
    fn domain(self) -> VirtBlocksErrorDomain {
        match self {
            Self::PlaygroundToyError(_) => VirtBlocksErrorDomain::PlaygroundToyError,
        }
    }

    fn code(self) -> u32 {
        match self {
            Self::PlaygroundToyError(e) => e as u32,
        }
    }

    fn message(self) -> String {
        match self {
            Self::PlaygroundToyError(e) => e.to_string(),
        }
    }
}

#[no_mangle]
pub extern "C" fn virtblocks_error_free(c_error: *mut VirtBlocksError) {
    if c_error.is_null() {
        return;
    }
    let _rust_error = unsafe { Box::from_raw(c_error) };
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_domain(
    c_error: *const VirtBlocksError,
) -> VirtBlocksErrorDomain {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    rust_error.domain()
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_code(c_error: *const VirtBlocksError) -> c_uint {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    rust_error.code() as c_uint
}

#[no_mangle]
pub extern "C" fn virtblocks_error_get_message(c_error: *const VirtBlocksError) -> *mut c_char {
    assert!(!c_error.is_null());

    let rust_error = unsafe { &*c_error };

    let rust_ret = rust_error.message();

    let tmp = CString::new(rust_ret).unwrap();
    unsafe { libc::strdup(tmp.as_ptr()) }
}
