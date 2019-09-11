// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use std::ffi::CStr;
use std::os::raw::{c_char, c_int, c_uint};
use std::ptr;

use crate::errors;

use virtblocks::command;

#[no_mangle]
pub extern "C" fn virtblocks_command_new(c_prog: *const c_char) -> *mut command::Command {
    if c_prog.is_null() {
        return ptr::null_mut();
    }
    let rust_prog = unsafe { CStr::from_ptr(c_prog) }.to_str().unwrap();

    Box::into_raw(Box::new(command::Command::new(rust_prog)))
}

#[no_mangle]
pub extern "C" fn virtblocks_command_free(c_command: *mut command::Command) {
    if c_command.is_null() {
        return;
    }
    let _rust_command = unsafe { Box::from_raw(c_command) };
}

#[no_mangle]
pub extern "C" fn virtblocks_command_add_arg(
    c_command: *mut command::Command,
    c_arg: *const c_char,
) {
    assert!(!c_command.is_null());
    assert!(!c_arg.is_null());

    let rust_command = unsafe { &mut *c_command };
    let rust_arg = unsafe { CStr::from_ptr(c_arg) }.to_str().unwrap();

    rust_command.add_arg(rust_arg);
}

#[no_mangle]
pub extern "C" fn virtblocks_command_take_fd(
    c_command: *mut command::Command,
    c_sourcefd: c_int,
    c_targetfd: c_int,
) {
    assert!(!c_command.is_null());

    let rust_command = unsafe { &mut *c_command };

    rust_command.take_fd(c_sourcefd, c_targetfd);
}

#[no_mangle]
pub extern "C" fn virtblocks_command_spawn(
    c_command: *mut command::Command,
    c_error: *mut *mut errors::Error,
) -> bool {
    assert!(!c_command.is_null());
    assert!(!c_error.is_null());

    let rust_command = unsafe { &mut *c_command };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_command.spawn() {
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
pub extern "C" fn virtblocks_command_id(c_command: *mut command::Command) -> c_uint {
    assert!(!c_command.is_null());

    let rust_command = unsafe { &mut *c_command };
    rust_command.id()
}

#[no_mangle]
pub extern "C" fn virtblocks_command_kill(
    c_command: *mut command::Command,
    c_error: *mut *mut errors::Error,
) -> bool {
    assert!(!c_command.is_null());
    assert!(!c_error.is_null());

    let rust_command = unsafe { &mut *c_command };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_command.kill() {
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
pub extern "C" fn virtblocks_command_wait(
    c_command: *mut command::Command,
    c_error: *mut *mut errors::Error,
) -> c_int {
    assert!(!c_command.is_null());
    assert!(!c_error.is_null());

    let rust_command = unsafe { &mut *c_command };

    unsafe {
        *c_error = ptr::null_mut() as *mut errors::Error;
    }

    match rust_command.wait() {
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
