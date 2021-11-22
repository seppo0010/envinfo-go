//go:build linux
// +build linux

package envinfo

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

func findStudioVersion(name, executable string) *Item {
	studioPathBytes, _ := exec.Command("which", fmt.Sprintf("%s.sh", executable)).Output()
	if len(studioPathBytes) == 0 {
		studioPathBytes = []byte(fmt.Sprintf("/snap/%s/current/bin/%s.sh", executable, executable))
		if _, err := os.Stat(string(studioPathBytes)); err != nil {
			return &Item{
				Name:    name,
				Version: "",
				Path:    "",
			}
		}
	}

	studioPath := strings.TrimSpace(string(studioPathBytes))
	versionPath := path.Join(studioPath, "..", "..", "build.txt")
	version, _ := os.ReadFile(versionPath)

	return &Item{
		Name:    name,
		Version: string(version),
		Path:    string(studioPath),
	}
}

func GetAndroidStudioVersions() *Item {
	name := "Android Studio"
	return findStudioVersion(name, "studio")
}

func GetIntelliJVersion() *Item {
	name := "IntelliJ"
	return findStudioVersion(name, "idea")
}

func GetPhpStormVersion() *Item {
	name := "PhpStorm"
	return findStudioVersion(name, "phpstorm")
}

func GetWebStormVersion() *Item {
	name := "WebStorm"
	return findStudioVersion(name, "webstorm.sh")
}

func GetXcodeVersion() *Item {
	return &Item{
		Name:    "Xcode",
		Path:    "",
		Version: "",
	}
}

func GetSublimeVersion() *Item {
	return NewGetItemBuilder("subl", "Sublime Text").Regex(regexp.MustCompile(`\d+`)).Get()
}
