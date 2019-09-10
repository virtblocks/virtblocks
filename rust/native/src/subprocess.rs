// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.
use std::collections::HashMap;
use std::ffi::OsStr;
use std::io;
use std::io::BufRead;
use std::io::BufReader;
use std::io::Write;
use std::mem;
use std::process::{Child, Command, Stdio};
use std::sync::mpsc;
use std::thread;

use std::os::unix::io::RawFd;
use std::os::unix::process::CommandExt;

use crate::closefds::close_fds_on_exec;
use crate::cvt::cvt_r;

#[derive(Debug)]
struct FileDesc(RawFd);

impl Drop for FileDesc {
    fn drop(&mut self) {
        unsafe { libc::close(self.0) };
    }
}

// targetfd -> sourcefd
#[derive(Debug)]
struct FdMap(HashMap<RawFd, FileDesc>);

// This is a rust anti-pattern, but should make bindings easier
// and can show how to wrap process/es
#[derive(Debug)]
#[allow(clippy::large_enum_variant)]
enum State {
    None,
    Child(Child, mpsc::Receiver<Vec<u8>>),
    Command(Command, FdMap),
}

/// A subprocess test, not meant to be a building block!
/// Good Rust design should follow Command/Child split.
///
/// # Examples
///
/// ```
/// use virtblocks_rust_native::subprocess;
///
/// fn main() -> Result<(), Box<dyn std::error::Error + 'static>> {
///     let mut sub = subprocess::Subprocess::new("ls");
///     sub.add_arg("/");
///     sub.spawn();
///     let status = sub.wait()?;
///     assert_eq!(status, 0);
///     Ok(())
/// }
/// ```
#[derive(Debug)]
pub struct Subprocess {
    state: State,
}

impl Subprocess {
    pub fn new<S: AsRef<OsStr>>(program: S) -> Self {
        let mut cmd = Command::new(program);
        let fdmap = HashMap::new();

        cmd.stdout(Stdio::piped());
        cmd.stderr(Stdio::piped());

        Self {
            state: State::Command(cmd, FdMap(fdmap)),
        }
    }

    pub fn spawn(&mut self) -> io::Result<()> {
        let state = mem::replace(&mut self.state, State::None);
        match state {
            State::Command(mut cmd, fdmap) => {
                let mut keepfds: Vec<RawFd> = fdmap.0.keys().cloned().collect();

                for fd in &[0, 1, 2] {
                    if !fdmap.0.contains_key(fd) {
                        keepfds.push(*fd);
                    }
                }

                unsafe {
                    // sadly, closing/passing fds isn't provided by rust
                    cmd.pre_exec(move || {
                        for (targetfd, sourcefd) in &fdmap.0 {
                            cvt_r(|| libc::dup2(sourcefd.0, *targetfd))?;
                        }
                        Ok(())
                    });
                    cmd.pre_exec(close_fds_on_exec(keepfds)?);
                }
                let mut child = cmd.spawn()?;
                let stdout = child.stdout.take().unwrap();
                let stderr = child.stderr.take().unwrap();

                let (tx, rx) = mpsc::channel();
                // process stdout/stderr
                // waiting for process would need another thread..
                // polling on IO+wait would be quite complicated
                thread::spawn(move || {
                    let reader = BufReader::new(stdout);
                    let mut vec = Vec::new();

                    // process stuff in a thread and send it back to main thread
                    reader
                        .lines()
                        .filter_map(|line| line.ok())
                        .for_each(|line| {
                            writeln!(&mut vec, "out: {}", line).unwrap();
                        });
                    tx.send(vec).unwrap();
                });
                thread::spawn(move || {
                    let reader = BufReader::new(stderr);
                    reader
                        .lines()
                        .filter_map(|line| line.ok())
                        .for_each(|line| println!("err: {}", line));
                });

                self.state = State::Child(child, rx);
                Ok(())
            }
            _ => Err(io::Error::new(io::ErrorKind::Other, "invalid state!")),
        }
    }

    pub fn add_arg<S: AsRef<OsStr>>(&mut self, arg: S) -> &mut Self {
        match &mut self.state {
            State::Command(cmd, _) => {
                cmd.arg(arg);
            }
            _ => panic!("invalid state"),
        }
        self
    }

    pub fn take_fd(&mut self, sourcefd: RawFd, targetfd: RawFd) -> &mut Self {
        match &mut self.state {
            State::Command(_, fdmap) => {
                fdmap.0.insert(targetfd, FileDesc(sourcefd));
            }
            _ => panic!("invalid state"),
        }
        self
    }

    pub fn id(&self) -> u32 {
        match &self.state {
            State::Child(child, _) => child.id(),
            _ => panic!("invalid state"),
        }
    }

    pub fn kill(&mut self) -> io::Result<()> {
        match &mut self.state {
            State::Child(child, _) => {
                child.kill()?;
                child.wait()?;
                Ok(())
            }
            _ => Err(io::Error::new(io::ErrorKind::Other, "invalid state!")),
        }
    }

    pub fn wait(&mut self) -> io::Result<i32> {
        match &mut self.state {
            State::Child(child, rx) => {
                let code = child
                    .wait()?
                    .code()
                    .ok_or_else(|| io::Error::new(io::ErrorKind::Other, "bad status!"));
                let received = rx.recv().unwrap();
                println!("{}", String::from_utf8(received).unwrap());
                code
            }
            _ => Err(io::Error::new(io::ErrorKind::Other, "invalid state!")),
        }
    }
}
