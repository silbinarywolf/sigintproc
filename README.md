# Signal Interrupt Process (sigintproc)

This library allows you to send interrupt signals to a child-process in your Go application in a way that also supports Windows.

As of Go 1.24 (2025), `process.Signal(os.Interrupt)` does not work on Windows, likely because there is no straight-forward one-size-fits-all solution

## Caveats

This module uses `GenerateConsoleCtrlEvent` for Windows and has caveats such as:
- The process that calls the function must share the console with the one we want to send it, otherwise it won't work. That also means the sender will receive that signal, too.
- By default on Windows it simply exits the process, if you have a GUI application, what might look like, it doesn’t work and crashed.

Source: https://blog.codetitans.pl/post/sending-ctrl-c-signal-to-another-application-on-windows/

## Example usage

My personal use-case was that I wanted my custom build tool for my Go application to safely terminate the application it launched when I hit CTRL+C in Terminal.

```go
func run() error {
    cmd := exec.CommandContext(ctx, "my-blocking-application")
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    // Windows: DefaultSysProcAttr prevents the console the application is running in from being "interrupted" too.
    // Non-Windows: Returns nil
    cmd.SysProcAttr = sigintproc.DefaultSysProcAttr()
    if err := cmd.Start(); err != nil {
        return err
    }
    cmd.Cancel = func() error {
        // Windows:     GenerateConsoleCtrlEvent(syscall.CTRL_BREAK_EVENT, process.Pid)
        // Non-Windows: process.Signal(os.Interrupt)
        return sigintproc.Interrupt(cmd.Process)
    }
    if err := cmd.Wait(); err != nil {
        return err
    }
    return nil
}
```
