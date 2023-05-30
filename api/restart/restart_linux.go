package restart

import (
	"errors"
	"os"
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
	}

	return nil
}
