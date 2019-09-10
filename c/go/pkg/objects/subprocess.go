// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package objects

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/subprocess"
	"sync"
)

var subprocessObjectsLock sync.RWMutex
var subprocessObjects = make([]*subprocess.Subprocess, 1)

func SubprocessAdd(s *subprocess.Subprocess) int {
	subprocessObjectsLock.Lock()
	defer subprocessObjectsLock.Unlock()

	subprocessObjects = append(subprocessObjects, s)
	return len(subprocessObjects) - 1
}

func SubprocessGet(ref int) *subprocess.Subprocess {
	subprocessObjectsLock.RLock()
	defer subprocessObjectsLock.RUnlock()

	return subprocessObjects[ref]
}

func SubprocessDel(ref int) {
	subprocessObjectsLock.Lock()
	defer subprocessObjectsLock.Unlock()

	subprocessObjects[ref] = nil
}
