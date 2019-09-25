package tests

import (
	"os"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QEMU", func() {

	pids := []int{}

	It("should start a qemu process with kvm", func() {
		cmd := exec.Command("qemu-system-x86_64", "--enable-kvm", "-nographic")
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		Expect(cmd.Start()).To(Succeed())
		pids = append(pids, cmd.Process.Pid)
		done := make(chan error)
		go func() { done <- cmd.Wait() }()
		select {
		case err := <-done:
			Expect(err).ToNot(HaveOccurred(), "expected the qemu process to stay alive for at least three seconds")
		case <-time.After(3 * time.Second):
			// success
		}
	})

	AfterEach(func() {
		for _, pid := range pids {
			process, err := os.FindProcess(pid)
			Expect(err).ToNot(HaveOccurred())
			err = process.Kill()
			if err != nil && !strings.Contains(err.Error(), "process already finished") {
				Expect(err).ToNot(HaveOccurred())
			}
		}
	})
})
