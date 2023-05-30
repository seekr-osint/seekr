package version

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

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

func GetLatestSeekrVersion() (SchematicVersion, error) {
	resp, err := http.Get("https://github.com/seekr-osint/seekr/releases/latest")
	if err != nil {
		return SchematicVersion{}, err
	}
	finalUrl := resp.Request.URL.String()

	return ParseSchematicVersion(path.Base(finalUrl))
}
