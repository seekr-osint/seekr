package api

import (
	"errors"
	"log"
	"os"
	"syscall"
	"time"
)

func (config ApiConfig) SetState(state bool) error {
	file, err := os.Create(config.LockFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lockType syscall.Flock_t
	if state {
		lockType.Type = syscall.F_WRLCK
	} else {
		lockType.Type = syscall.F_UNLCK
	}

	err = syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &lockType)
	if err != nil {
		return errors.New("lock file already exists")
	}

	return nil
}

func (config ApiConfig) CheckState() bool {
	_, err := os.Stat(config.LockFilePath)
	return !os.IsNotExist(err)
}
func (config ApiConfig) WaitForLockFile() {
	for !config.CheckState() {
		log.Printf("Waiting for lock file: %s\n", config.LockFilePath)
		time.Sleep(time.Second * 1)
	}
}
