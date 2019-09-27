// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/virtblocks/virtblocks/go/pkg/command"
	"github.com/virtblocks/virtblocks/go/pkg/devices"
	"github.com/virtblocks/virtblocks/go/pkg/vm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func processFlags() (disk string, cpus uint, mem uint, err error) {
	flag.StringVar(&disk, "disk", "", "path to disk image (qcow2)")
	flag.UintVar(&cpus, "cpus", 1, "number of CPUs")
	flag.UintVar(&mem, "mem", 128, "amount of memory (MiB)")

	flag.Parse()

	if disk == "" {
		err = errors.New("need a disk")
		return
	}
	if filepath.Ext(disk) != ".qcow2" {
		err = errors.New("need disk in qcow2 format")
		return
	}
	if cpus < 1 {
		err = errors.New("need at least one CPU")
		return
	}
	if mem < 64 {
		err = errors.New("need more memory")
		return
	}

	return
}

func main() {
	diskPath, cpus, mem, err := processFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing flags: %v\n", err)
		os.Exit(1)
	}

	// Derive the path of the socket used for serial console communication
	// from the path of the disk image
	serialSockPath := fmt.Sprintf("%s%s", strings.TrimSuffix(diskPath, "qcow2"), "serial-sock")

	vm := vm.NewDescription()
	vm.SetCpus(cpus).SetMemory(mem)
	vm.SetDisk(devices.NewDisk().SetFilename(diskPath))
	vm.SetSerial(devices.NewSerial().SetPath(serialSockPath))

	qemuArgs, err := vm.QemuCommandLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating QEMU command: %v\n", err)
		os.Exit(1)
	}

	qemu := command.NewCommand(qemuArgs[0])
	qemu.AddArgs(qemuArgs[1:]...)

	socat := command.NewCommand("socat")
	socat.AddArgs("stdin,raw,echo=0,escape=0x11", fmt.Sprintf("unix-connect:%s", serialSockPath))
	socat.SetStdin(os.Stdin).SetStdout(os.Stdout).SetStderr(os.Stderr)

	fmt.Print("Starting VM, please wait... (Ctrl+Q to quit)")
	qemu.Spawn()

	// Wait a bit before starting socat, because we can't connect to the
	// socket until QEMU has created it. This is extremely racy and only
	// acceptable in a silly demo such as this one
	time.Sleep(500 * time.Millisecond)
	socat.Spawn()

	// Wait for the user to quit socat, and kill QEMU when that happens
	socat.Wait()
	qemu.Kill()

	rc, err := qemu.Wait()
	if err != nil {
		if rc == 255 {
			// 255 means killed by a signal, which is the expected outcome
			// and doesn't actually represent an error most of the time.
			// If the QEMU process has been killed by someone else, well...
			// Let's assume that won't happen O:-)
			rc = 0
		} else {
			fmt.Fprintf(os.Stderr, "error during execution: %v\n", err)
		}
	}

	os.Exit(rc)
}
