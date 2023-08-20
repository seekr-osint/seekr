package version

import (
	"fmt"
	"log"
	"runtime"
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
	return ver.CompareIsLatest(latestVersion)

}

func (sv SchematicVersion) DownloadURL() string {
	return fmt.Sprintf("https://github.com/seekr-osint/seekr/releases/download/%s/%s", sv, sv.BinaryName())
}

func (sv SchematicVersion) BinaryName() string {
	platform, arch := GetOS()
	fileName := fmt.Sprintf("seekr_%s_%s_%s", sv, platform, arch)
	if runtime.GOOS == "windows" {
		fileName += ".exe"
	}
	return fileName
}

func (sv SchematicVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.Major, sv.Minor, sv.Patch)
}

func (sv1 SchematicVersion) CompareIsLatest(sv2 SchematicVersion) bool {
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
