// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

//! # Virt Blocks
//!
//! `virtblocks` is a collection of utilities which can be either
//! consumed on their own, or to replace certain parts of libvirt's
//! QEMU driver.

pub mod devices;
pub mod playground;
pub mod subprocess;

mod cvt;
