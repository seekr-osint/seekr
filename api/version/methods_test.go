package version

//import "testing"
import (
	"runtime"
	"testing"
)

func TestDownloadURL(t *testing.T) {
	t.Skip()
	sv := SchematicVersion{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	expectedURL := "https://github.com/seekr-osint/seekr/releases/download/1.2.3/seekr_1.2.3_platform_arch"
	if runtime.GOOS == "windows" {
		expectedURL += ".exe"
	}

	downloadURL := sv.DownloadURL()
	if downloadURL != expectedURL {
		t.Errorf("DownloadURL() returned incorrect URL, got: %s, want: %s", downloadURL, expectedURL)
	}
}

func TestBinaryName(t *testing.T) {
	t.Skip()
	sv := SchematicVersion{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	// Mock the GetOS() function to return consistent values for testing
	//	GetOS = func() (string, string) {
	//		return "platform", "arch"
	//	}

	expectedName := "seekr_1.2.3_platform_arch"
	if runtime.GOOS == "windows" {
		expectedName += ".exe"
	}

	binaryName := sv.BinaryName()
	if binaryName != expectedName {
		t.Errorf("BinaryName() returned incorrect name, got: %s, want: %s", binaryName, expectedName)
	}
}

func TestString(t *testing.T) {
	sv := SchematicVersion{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	expectedString := "1.2.3"
	stringRep := sv.String()
	if stringRep != expectedString {
		t.Errorf("String() returned incorrect string representation, got: %s, want: %s", stringRep, expectedString)
	}
}
