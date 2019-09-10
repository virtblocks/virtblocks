// Virt Blocks
//
// Copyright (C) 2019 Red Hat, Inc.
//
// This software is distributed under the terms of the MIT License.
// See the LICENSE file in the top level directory for details.

package main

import (
	"github.com/virtblocks/virtblocks/go/native/pkg/subprocess"
	"log"
)

func main() {
	var sub = subprocess.NewSubprocess("./test.sh")
	sub.Spawn()
	id, err := sub.Id()
	if err == nil {
		log.Printf("PID: %d", id)
	}
	err = sub.Wait()
	log.Printf("Command finished with error: %v", err)
}
