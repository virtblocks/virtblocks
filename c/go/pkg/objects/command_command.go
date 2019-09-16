// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/command"
	"sync"
)

var commandObjectsLock sync.RWMutex
var commandObjects = make([]*command.Command, 1)

func CommandAdd(command *command.Command) int {
	commandObjectsLock.Lock()
	defer commandObjectsLock.Unlock()

	commandObjects = append(commandObjects, command)
	return len(commandObjects) - 1
}

func CommandGet(ref int) *command.Command {
	commandObjectsLock.RLock()
	defer commandObjectsLock.RUnlock()

	return commandObjects[ref]
}

func CommandDel(ref int) {
	commandObjectsLock.Lock()
	defer commandObjectsLock.Unlock()

	commandObjects[ref] = nil
}
