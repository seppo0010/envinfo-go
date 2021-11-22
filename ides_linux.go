//go:build linux
// +build linux

package envinfo

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

func GetAndroidStudioVersions() (*Item, error) {
	name := "Android Studio"
	studioPathBytes, _ := exec.Command("which", "studio.sh").Output()
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
