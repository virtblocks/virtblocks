// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package subprocess

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

type Subprocess struct {
	Cmd *exec.Cmd
}

type SubStdout struct {
	buf bytes.Buffer
}

func (b *SubStdout) Write(p []byte) (n int, err error) {
	log.Println(p)
	return b.buf.Write(p)
}

func NewSubprocess(cmd string) *Subprocess {
	c := exec.Command(cmd)

	var out SubStdout
	c.Stdout = &out
	var err bytes.Buffer
	c.Stderr = &err

	return &Subprocess{
		Cmd: c,
	}
}

func (s *Subprocess) AddArg(arg string) {
	s.Cmd.Args = append(s.Cmd.Args, arg)
}

func (s *Subprocess) TakeFd(fd uintptr) {
	file := os.NewFile(fd, "")
	s.Cmd.ExtraFiles = append(s.Cmd.ExtraFiles, file)
}

func (s *Subprocess) Spawn() error {
	return s.Cmd.Start()
}

func (s *Subprocess) Id() (int, error) {
	if s.Cmd.Process != nil {
		return s.Cmd.Process.Pid, nil
	}
	return -1, nil
}

func (s *Subprocess) Wait() error {
	err := s.Cmd.Wait()
	log.Println(s.Cmd.Stderr)
	return err
}
