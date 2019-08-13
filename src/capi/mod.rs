use std::ffi::CStr;
use std::ffi::CString;
use std::mem;
use std::os::raw::c_char;

use super::util;
use super::devices;

#[no_mangle]
pub extern "C"
fn virtblocks_util_build_file_name(file_name: *mut *mut c_char,
                                   base: *const c_char,
                                   ext: *const c_char) -> i32
{
    if base.is_null() || ext.is_null() || file_name.is_null() {
        return -1;
    }
    let base = unsafe { CStr::from_ptr(base) }.to_str().unwrap();
    let ext = unsafe { CStr::from_ptr(ext) }.to_str().unwrap();

    let res = util::build_file_name(base, ext);

    unsafe { *file_name = CString::new(res).unwrap().into_raw() }
    return 0;
}

#[no_mangle]
pub extern "C"
fn virtblocks_devices_memballoon_new() -> *mut devices::Memballoon {
    Box::into_raw(Box::new(devices::Memballoon::new()))
}

#[no_mangle]
pub extern "C"
fn virtblocks_devices_memballoon_free(c_memballoon: *mut devices::Memballoon) {
    let _rust_memballoon = unsafe {
        assert!(!c_memballoon.is_null());
        Box::from_raw(c_memballoon)
    };
}

#[no_mangle]
pub extern "C"
fn virtblocks_devices_memballoon_set_model(c_memballoon: *mut devices::Memballoon,
                                           model: devices::MemballoonModel)
{
    let rust_memballoon = unsafe {
        assert!(!c_memballoon.is_null());
        &mut *c_memballoon
    };

    rust_memballoon.set_model(model)
}

#[no_mangle]
pub extern "C"
fn virtblocks_devices_memballoon_get_model(c_memballoon: *const devices::Memballoon) -> devices::MemballoonModel
{
    let rust_memballoon = unsafe {
        assert!(!c_memballoon.is_null());
        &*c_memballoon
    };

    rust_memballoon.get_model()
}

#[no_mangle]
pub extern "C"
fn virtblocks_devices_memballoon_to_str(c_memballoon: *const devices::Memballoon) -> *const c_char {
    let rust_memballoon = unsafe {
        assert!(!c_memballoon.is_null());
        &*c_memballoon
    };

    let s = CString::new(rust_memballoon.to_str()).unwrap();
    let ret = s.as_ptr();
    mem::forget(s);
    ret
}
