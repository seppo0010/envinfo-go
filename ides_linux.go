//go:build linux
// +build linux

package envinfo

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

func findStudioVersion(name, executable string) (*Item, error) {
	studioPathBytes, _ := exec.Command("which", executable).Output()
	if len(studioPathBytes) == 0 {
		return &Item{
			Name:    name,
			Version: "",
			Path:    "",
		}, nil
	}

	studioPath := strings.TrimSpace(string(studioPathBytes))
	versionPath := path.Join(studioPath, "..", "..", "build.txt")
	version, _ := os.ReadFile(versionPath)

	return &Item{
		Name:    name,
		Version: string(version),
		Path:    string(studioPath),
	}, nil
}

func GetAndroidStudioVersions() (*Item, error) {
	name := "Android Studio"
	return findStudioVersion(name, "studio.sh")
}

func GetIntelliJVersion() (*Item, error) {
	name := "IntelliJ"
	return findStudioVersion(name, "idea.sh")
}

func GetPhpStormVersion() (*Item, error) {
	name := "PhpStorm"
	return findStudioVersion(name, "phpstorm.sh")
}
