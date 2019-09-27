// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package command

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

type Command struct {
	cmd *exec.Cmd
}

func NewCommand(prog string) *Command {
	return &Command{
		cmd: exec.Command(prog),
	}
}

func (self *Command) AddArgs(args ...string) *Command {
	self.cmd.Args = append(self.cmd.Args, args...)
	return self
}

func (self *Command) TakeFd(fd uintptr) *Command {
	file := os.NewFile(fd, "")
	self.cmd.ExtraFiles = append(self.cmd.ExtraFiles, file)
	return self
}

func (self *Command) SetStdin(reader io.Reader) *Command {
	self.cmd.Stdin = reader
	return self
}

func (self *Command) SetStdout(writer io.Writer) *Command {
	self.cmd.Stdout = writer
	return self
}

func (self *Command) SetStderr(writer io.Writer) *Command {
	self.cmd.Stderr = writer
	return self
}

func (self *Command) Spawn() error {
	return self.cmd.Start()
}

func (self *Command) Id() (uint, error) {
	if self.cmd.Process == nil {
		return 0, errors.New("process not running")
	}
	return uint(self.cmd.Process.Pid), nil
}

func (self *Command) Kill() error {
	if self.cmd.Process == nil {
		return errors.New("process not running")
	}
	self.cmd.Process.Kill()
	return nil
}

func (self *Command) Wait() (int, error) {
	var err = self.cmd.Wait()

	// Successful command execution
	if err == nil {
		return 0, nil
	}

	// Failed command execution we can obtain a sensible exit code for
	if exiterr, ok := err.(*exec.ExitError); ok {
		return exiterr.ExitCode(), exiterr
	}

	// Unknown failure, only option is to return -1
	return -1, err
}
