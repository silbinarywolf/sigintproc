package sigintproc

import (
	"os"
	"syscall"
)

// Logic borrowed from: https://go.dev/src/os/signal/signal_windows_test.go
//
// GenerateConsoleCtrlEvent caveats:
//   - To make it actually work, the process that calls the function must share the console with the one we want to send it, otherwise it won't work.
//     That also means the sender will receive that signal, too
//   - By default on Windows it simply exits the process, if you have a GUI application, what might look like, it doesn't work and crashed.
//
// Source: https://blog.codetitans.pl/post/sending-ctrl-c-signal-to-another-application-on-windows/
func signalInterrupt(process *os.Process) error {
	d, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	p, err := d.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return err
	}
	r, _, err := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(process.Pid))
	if r == 0 {
		return err
	}
	return nil
}

func defaultSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		// Force the new process into a new process group so that it won't terminate the parent
		// process Terminal.
		//
		// For example: If you start a Bash Terminal in VSCode and trigger "GenerateConsoleCtrlEvent", it will kill that terminal in VSCode as well.
		// if you don't use this.
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}
