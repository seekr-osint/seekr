package restart

import (
	"errors"
	"os"
	"os/exec"
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

	// Start a new process with the current executable
	cmd := exec.Command(executable, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return ErrStartNewBin
	}

	// Wait for the new process to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
