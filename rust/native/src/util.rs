// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

/// Build a filename (as a simple FFI test)
///
/// # Examples
///
/// ```
/// use virtblocks_rust_native::util;
///
/// let fname = util::build_file_name("base", ".ext");
///
/// assert_eq!("base.ext", fname);
/// ```
pub fn build_file_name(base: &str, ext: &str) -> String {
    let mut ret = String::from(base);
    ret.push_str(ext);
    ret
}
