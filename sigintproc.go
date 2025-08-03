package sigintproc

import (
	"os"
	"syscall"
)

// Interrupt will call "process.Signal(os.Interrupt)" for all operating systems except Windows
// for Windows it will call "GenerateConsoleCtrlEvent(CTRL_BREAK_EVENT, process_id)"
func Interrupt(process *os.Process) error {
	return signalInterrupt(process)
}

// DefaultSysProcAttr prevents the terminal the application is running in from being "interrupted" too.
// Without this, the Interrupt signal can for example, kill the terminal in VSCode.
//
// - Windows: "&syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}"
// - Other: nil
func DefaultSysProcAttr() *syscall.SysProcAttr {
	return defaultSysProcAttr()
}
