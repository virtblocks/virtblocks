// Copyright (C) 2019 Red Hat, Inc.
// SPDX-License-Identifier: LGPL-2.1-or-later

package objects

import (
	"github.com/virtblocks/virtblocks/go/pkg/command"
	"sync"
)

var commandObjectsLock sync.RWMutex
var commandObjects = make([]*command.Command, 1)

func CommandAdd(command *command.Command) uint {
	commandObjectsLock.Lock()
	defer commandObjectsLock.Unlock()

	commandObjects = append(commandObjects, command)
	return uint(len(commandObjects) - 1)
}

func CommandGet(ref uint) *command.Command {
	commandObjectsLock.RLock()
	defer commandObjectsLock.RUnlock()

	return commandObjects[ref]
}

func CommandDel(ref uint) {
	commandObjectsLock.Lock()
	defer commandObjectsLock.Unlock()

	commandObjects[ref] = nil
}
