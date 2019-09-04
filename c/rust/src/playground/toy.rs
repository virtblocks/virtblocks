// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

#![cfg_attr(feature = "cargo-clippy", allow(clippy::not_unsafe_ptr_arg_deref))]

use std::ffi::CStr;
use std::ffi::CString;
use std::os::raw::c_char;
use std::os::raw::c_void;
use std::ptr::null_mut;

use virtblocks::playground;

use crate::errors;

type ToyCallback = unsafe extern "C" fn(*const playground::Toy, *const c_char, *mut c_void) -> bool;
type ToyCallbackDataFree = unsafe extern "C" fn(*mut c_void);

// Wrapper for ToyCallback
//
// Since C doesn't have native support for closures, it's costumary to
// allow user to attach an opaque data pointer to the object which is then
// passed to all subsequent invocations of the callback and can be used to
// maintain state; along with that comes the responsibility of eventually
// freeing this data, which requires storing a data_free callback as well.
//
// We store everything in this wrapper: when the Toy is dropped, the
// corresponding Rust callback will be dropped along with it, and since
// that callback owns the wrapper instance, the latter's drop() method
// will be invoked and the C data will be freed.
struct ToyCallbackWrapper {
    callback: ToyCallback,
    data: *mut c_void,
    data_free: Option<ToyCallbackDataFree>,
}

impl Drop for ToyCallbackWrapper {
    fn drop(&mut self) {
        // If we were not provided data or a data_free callback, then
        // we don't have anything to do here
        if self.data.is_null() || self.data_free.is_none() {
            return;
        }
        unsafe {
            self.data_free.unwrap()(self.data);
        }
        // For extra safety, replace the data pointer with NULL
        self.data = null_mut();
    }
}

#[no_mangle]
pub extern "C" fn virtblocks_playground_toy_new(
    c_base: *const c_char,
    c_filter: Option<ToyCallback>,
    c_data: *mut c_void,
    c_data_free: Option<ToyCallbackDataFree>,
) -> *mut playground::Toy {
    assert!(!c_base.is_null());
    assert!(!c_filter.is_none());

    let c_filter_wrapper = ToyCallbackWrapper {
        callback: c_filter.unwrap(),
        data: c_data,
        data_free: c_data_free,
    };

    let rust_base = unsafe { CStr::from_ptr(c_base) }.to_str().unwrap();
    let rust_filter = move |rust_toy: &playground::Toy, rust_ext: &str| unsafe {
        let c_toy = rust_toy as *const playground::Toy;
        let c_ext = CString::new(rust_ext).unwrap();

        // At this point all Rust structures have been converted to the
        // corresponding C structures, and so we can finally invoke the
        // user-provided callback
        (c_filter_wrapper.callback)(c_toy, c_ext.as_ptr(), c_filter_wrapper.data)
    };

    Box::into_raw(Box::new(playground::Toy::new(rust_base, rust_filter)))
}

#[no_mangle]
pub extern "C" fn virtblocks_playground_toy_free(c_toy: *mut playground::Toy) {
    if c_toy.is_null() {
        return;
    }
    let _rust_toy = unsafe { Box::from_raw(c_toy) };
}

#[no_mangle]
pub extern "C" fn virtblocks_playground_toy_get_base(c_toy: *const playground::Toy) -> *mut c_char {
    assert!(!c_toy.is_null());

    let rust_toy = unsafe { &*c_toy };

    let rust_ret = rust_toy.base();

    let c_ret = CString::new(rust_ret).unwrap();

    // Unfortunately we have to make a copy here because the CString we
    // have just created will be dropped once the function exit, meaning
    // if we returned a borrow to the C code it would be pointing to
    // invalid memory right away
    unsafe { libc::strdup(c_ret.as_ptr()) }
}

#[no_mangle]
pub extern "C" fn virtblocks_playground_toy_run(
    c_toy: *const playground::Toy,
    c_ext: *const c_char,
    c_error: *mut *mut errors::VirtBlocksError,
) -> *mut c_char {
    assert!(!c_toy.is_null());
    assert!(!c_ext.is_null());
    assert!(!c_error.is_null());

    let rust_toy = unsafe { &*c_toy };
    let rust_ext = unsafe { CStr::from_ptr(c_ext) }.to_str().unwrap();

    let rust_ret = rust_toy.run(rust_ext);

    if rust_ret.is_ok() {
        // Successful execution: the error returned to C will be NULL
        unsafe {
            *c_error = null_mut() as *mut errors::VirtBlocksError;
        }

        // We need to convert the Rust string to a C string
        let tmp = CString::new(rust_ret.unwrap()).unwrap();
        unsafe { libc::strdup(tmp.as_ptr()) }
    } else {
        // Failed execution: we need to convert the Rust error to a
        // C error
        let tmp = errors::VirtBlocksError::from(rust_ret.unwrap_err());
        unsafe {
            *c_error = Box::into_raw(Box::new(tmp));
        }

        // We just return NULL instead of a string in this case
        null_mut() as *mut c_char
    }
}
