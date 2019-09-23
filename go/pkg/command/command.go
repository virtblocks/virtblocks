// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package command

import (
	"os"
	"os/exec"
)

type Command struct {
	Cmd *exec.Cmd
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
		Cmd: c,
	}
}

func (self *Command) AddArg(arg string) {
	self.Cmd.Args = append(self.Cmd.Args, arg)
}

func (self *Command) TakeFd(fd uintptr) {
	file := os.NewFile(fd, "")
	self.Cmd.ExtraFiles = append(self.Cmd.ExtraFiles, file)
}

func (self *Command) Spawn() error {
	return self.Cmd.Start()
}

func (self *Command) Id() (int, error) {
	if self.Cmd.Process != nil {
		return self.Cmd.Process.Pid, nil
	}
	return -1, nil
}

func (self *Command) Wait() error {
	return self.Cmd.Wait()
}
