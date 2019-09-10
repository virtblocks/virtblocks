// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use std::ffi::CStr;
use std::os::raw::{c_char, c_int, c_uint};
use std::ptr;

use super::errors;

use virtblocks::subprocess;

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_new(prog: *const c_char) -> *mut subprocess::Subprocess {
    if prog.is_null() {
        return ptr::null_mut();
    }
    let prog = unsafe { CStr::from_ptr(prog) }.to_str().unwrap();

    Box::into_raw(Box::new(subprocess::Subprocess::new(prog)))
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_free(c_subprocess: *mut subprocess::Subprocess) {
    if c_subprocess.is_null() {
        return;
    }
    let _rust_subprocess = unsafe { Box::from_raw(c_subprocess) };
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_add_arg(
    c_subprocess: *mut subprocess::Subprocess,
    c_arg: *const c_char,
) {
    assert!(!c_subprocess.is_null());
    assert!(!c_arg.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };
    let rust_arg = unsafe { CStr::from_ptr(c_arg) }.to_str().unwrap();

    rust_subprocess.add_arg(rust_arg);
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_take_fd(
    c_subprocess: *mut subprocess::Subprocess,
    c_sourcefd: c_int,
    c_targetfd: c_int,
) {
    assert!(!c_subprocess.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };

    rust_subprocess.take_fd(c_sourcefd, c_targetfd);
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_spawn(
    c_subprocess: *mut subprocess::Subprocess,
    c_error: *mut *mut errors::Error,
) -> bool {
    assert!(!c_subprocess.is_null());
    assert!(!c_error.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_subprocess.spawn() {
        Err(e) => {
            let tmp = errors::Error::from(e);
            unsafe {
                *c_error = Box::into_raw(Box::new(tmp));
            };
            false
        }
        _ => true,
    }
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_id(c_subprocess: *mut subprocess::Subprocess) -> c_uint {
    assert!(!c_subprocess.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };
    rust_subprocess.id()
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_kill(
    c_subprocess: *mut subprocess::Subprocess,
    c_error: *mut *mut errors::Error,
) -> bool {
    assert!(!c_subprocess.is_null());
    assert!(!c_error.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_subprocess.kill() {
        Err(e) => {
            let tmp = errors::Error::from(e);
            unsafe {
                *c_error = Box::into_raw(Box::new(tmp));
            };
            false
        }
        _ => true,
    }
}

#[no_mangle]
pub extern "C" fn virtblocks_subprocess_wait(
    c_subprocess: *mut subprocess::Subprocess,
    c_error: *mut *mut errors::Error,
) -> c_int {
    assert!(!c_subprocess.is_null());
    assert!(!c_error.is_null());

    let rust_subprocess = unsafe { &mut *c_subprocess };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_subprocess.wait() {
        Err(e) => {
            let tmp = errors::Error::from(e);
            unsafe {
                *c_error = Box::into_raw(Box::new(tmp));
            };
            0
        }
        Ok(c) => c,
    }
}
