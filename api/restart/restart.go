package restart

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

var (
	ErrGetCurrentBinary = errors.New("failed to get current executable path")
	ErrStartNewBin      = errors.New("failed to start new binary instance")
)

func RestartBinary() error {
	// Get the path of the current executable
	executable, err := os.Executable()
	if err != nil {
		return ErrGetCurrentBinary
	}

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if err := syscall.Exec(executable, os.Args, os.Environ()); err != nil {
			return ErrStartNewBin
		}
	} else if runtime.GOOS == "windows" {
		// Start a new instance of the current binary
		cmd := exec.Command(executable, os.Args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}
		if err := cmd.Start(); err != nil {
			return ErrStartNewBin
		}

		// Terminate the current process
		currentProcess, err := syscall.GetCurrentProcess()
		if err != nil {
			return err
		}
		err = syscall.TerminateProcess(currentProcess, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
