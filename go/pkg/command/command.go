// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package command

import (
	"errors"
	"os"
	"os/exec"
)

type Command struct {
	cmd *exec.Cmd
}

type CommandStdout struct{}

func (b *CommandStdout) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

type CommandStderr struct{}

func (b *CommandStderr) Write(p []byte) (n int, err error) {
	return os.Stderr.Write(p)
}

func NewCommand(prog string) *Command {
	c := exec.Command(prog)

	c.Stdout = &CommandStdout{}
	c.Stderr = &CommandStderr{}

	return &Command{
		cmd: c,
	}
}

func (self *Command) AddArg(arg string) {
	self.cmd.Args = append(self.cmd.Args, arg)
}

func (self *Command) TakeFd(fd uintptr) {
	file := os.NewFile(fd, "")
	self.cmd.ExtraFiles = append(self.cmd.ExtraFiles, file)
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
