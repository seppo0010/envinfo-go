//go:build linux
// +build linux

package envinfo

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func findStudioVersion(name, executable string) (*Item, error) {
	studioPathBytes, _ := exec.Command("which", fmt.Sprintf("%s.sh", executable)).Output()
	if len(studioPathBytes) == 0 {
		studioPathBytes = []byte(fmt.Sprintf("/snap/%s/current/bin/%s.sh", executable, executable))
		if _, err := os.Stat(string(studioPathBytes)); err != nil {
			return &Item{
				Name:    name,
				Version: "",
				Path:    "",
			}, nil
		}
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
	return findStudioVersion(name, "studio")
}

func GetIntelliJVersion() (*Item, error) {
	name := "IntelliJ"
	return findStudioVersion(name, "idea")
}

func GetPhpStormVersion() (*Item, error) {
	name := "PhpStorm"
	return findStudioVersion(name, "phpstorm")
}

func GetWebStormVersion() (*Item, error) {
	name := "WebStorm"
	return findStudioVersion(name, "webstorm.sh")
}
