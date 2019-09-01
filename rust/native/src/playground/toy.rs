// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

use std::error;
use std::fmt;

/// Possible error types for Toy
#[repr(u32)]
#[derive(Copy, Clone, Debug)]
pub enum ToyError {
    /// User-provided callback failed
    CallbackFailed,
}

impl fmt::Display for ToyError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let error_str = match self {
            Self::CallbackFailed => "Callback failed",
        };
        write!(f, "{}", error_str)
    }
}

impl error::Error for ToyError {}

/// A toy object
///
/// This object is not intended to be useful, but rather to exercise
/// some features such as:
///
///   * error reporting;
///   * callbacks;
///   * returning allocated memory;
///
/// that might be problematic when crossing the language boundary.
///
/// # Examples
///
/// ```
/// use virtblocks_rust_native::playground;
///
/// let toy = playground::Toy::new("Cargo", |toy, ext| toy.base() == "Cargo");
/// let res = toy.run("toml");
///
/// assert!(res.is_ok());
/// assert_eq!("Cargo.toml", res.unwrap());
/// ```
pub struct Toy {
    base: String,
    filter: Box<dyn Fn(&Self, &str) -> bool>,
}

impl Toy {
    pub fn new<F>(base: &str, filter: F) -> Self
    where
        F: Fn(&Self, &str) -> bool + 'static,
    {
        Self {
            base: String::from(base),
            filter: Box::new(filter),
        }
    }

    pub fn base(&self) -> &str {
        &self.base
    }

    pub fn run(&self, ext: &str) -> Result<String, ToyError> {
        if (self.filter)(self, ext) {
            // If the user-provided filter succeeds, then build a string
            let mut res = self.base.clone();
            res.push('.');
            res.push_str(ext);
            Ok(res)
        } else {
            // Otherwise error out
            Err(ToyError::CallbackFailed)
        }
    }
}
