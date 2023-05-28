package version

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func (ver SchematicVersion) GetLatest() SchematicVersion {
	latestVersion, err := GetLatestSeekrVersion()
	if err != nil {
		log.Printf("error getting latest seekr version: %s\n", ver)
	}
	return latestVersion
}

func (ver SchematicVersion) IsLatest() bool { // false if error
	latestVersion, err := GetLatestSeekrVersion()
	if err != nil {
		log.Printf("error getting latest seekr version: %s\n", ver)
		return false
	}
	return ver.Latest(latestVersion)

}
func (ver SchematicVersion) DownloadURL() string {
	return fmt.Sprintf("https://github.com/seekr-osint/seekr/releases/download/%s/%s", ver, GetBinaryName(ver))
}
func promptYesNo(question string) bool {
	prompt := promptui.Select{
		Label: question,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	return result == "Yes"
}

func updateBinary(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "temp-binary-*")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return err
	}

	binaryPath, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command(binaryPath)
	cmd.ExtraFiles = []*os.File{tempFile}
	cmd.Env = append(os.Environ(), fmt.Sprintf("_SEEKR_UPDATE_BINARY=%s", tempFile.Name()))
	err = cmd.Start()
	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}

func downloadAndReplaceBinary(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	binaryPath, err := os.Executable()
	if err != nil {
		return err
	}

	tempFile, err := os.CreateTemp("", "newbinary")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return err
	}

	err = os.Chmod(tempFile.Name(), 0755)
	if err != nil {
		return err
	}

	binaryDir := filepath.Dir(binaryPath)

	newBinaryPath := filepath.Join(binaryDir, "newbinary")

	err = os.Rename(tempFile.Name(), newBinaryPath)
	if err != nil {
		return err
	}

	err = os.Rename(newBinaryPath, binaryPath)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrInvalidVersionFormat         = errors.New("invalid schematic version format")
	ErrInvalidMajorVersionComponent = errors.New("invalid major version component")
	ErrInvalidMinorVersionComponent = errors.New("invalid minor version component")
	ErrInvalidPatchVersionComponent = errors.New("invalid patch version component")
)

const (
	platformWin = "windows"
	platformLin = "linux"
	archAmd64   = "amd64"
	archArm64   = "arm64"
)

func GetStats() (string, string) {
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

func GetBinaryName(version SchematicVersion) string {
	platform, arch := GetStats()
	fileName := fmt.Sprintf("seekr_%s_%s_%s", version, platform, arch)
	if runtime.GOOS == "windows" {
		fileName += ".exe"
	}
	return fileName
}

func (sv SchematicVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.Major, sv.Minor, sv.Patch)
}

func (sv1 SchematicVersion) Latest(sv2 SchematicVersion) bool {
	if sv1.Major > sv2.Major {
		return true
	} else if sv1.Major < sv2.Major {
		return false
	}

	if sv1.Minor > sv2.Minor {
		return true
	} else if sv1.Minor < sv2.Minor {
		return false
	}

	if sv1.Patch > sv2.Patch {
		return true
	} else {
		return false
	}
}

func GetLatestSeekrVersion() (SchematicVersion, error) {
	resp, err := http.Get("https://github.com/seekr-osint/seekr/releases/latest")
	if err != nil {
		return SchematicVersion{}, err
	}
	finalUrl := resp.Request.URL.String()

	return ParseSchematicVersion(path.Base(finalUrl))
}

func ParseSchematicVersion(versionStr string) (SchematicVersion, error) {
	var version SchematicVersion

	parts := strings.Split(versionStr, ".")
	if len(parts) != 3 {
		return version, ErrInvalidVersionFormat
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return version, ErrInvalidMajorVersionComponent
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return version, ErrInvalidMinorVersionComponent
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return version, ErrInvalidPatchVersionComponent
	}

	version.Major = major
	version.Minor = minor
	version.Patch = patch

	return version, nil
}

func checkVer() {
	destPath := os.Getenv("_SEEKR_UPDATE_BINARY")
	if destPath != "" {
		exePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exeFile, err := os.Open(exePath)
		if err != nil {
			panic(err)
		}
		defer exeFile.Close()

		destFile, err := os.Create(filepath.Join(destPath, filepath.Base(exePath)))
		if err != nil {
			panic(err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, exeFile)
		if err != nil {
			panic(err)
		}
	}
}
