package version

import (
	"fmt"
	"runtime"
)

const (
	platformWin = "windows"
	platformLin = "linux"
	archAmd64   = "amd64"
	archArm64   = "arm64"
)

func GetOS() (string, string) {
	var platform, arch string
	switch runtime.GOOS {
	case "windows":
		platform = platformWin
		arch = archAmd64
	case "linux":
		platform = platformLin
		arch = archArm64
	default:
		fmt.Printf("WARNING: unsupported OS\n") // FIXME weird binary name in case of unsupported OS
	}
	return platform, arch
}
