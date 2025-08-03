//go:build !windows

package sigintproc

import (
	"os"
	"syscall"
)

func signalInterrupt(process *os.Process) error {
	return process.Signal(os.Interrupt)
}

func defaultSysProcAttr() *syscall.SysProcAttr {
	return nil
}
