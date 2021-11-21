//go:build darwin
// +build darwin

package envinfo

import (
	"os/exec"
	"regexp"
)

func contains(s string, arr []string) bool {
	for _, el := range arr {
		if el == s {
			return true
		}
	}
	return false
}

func GetIOSVersions() ([]string, error) {
	sdks, _ := exec.Command("xcodebuild", "-showsdks").Output()
	regex := regexp.MustCompile(`[\w]+\s[\d|.]+`)
	matches := regex.FindAllString(string(sdks), -1)
	platforms := make([]string, 0, len(matches))
	for _, platform := range matches {
		if !contains(platform, platforms) {
			platforms = append(platforms, platform)
		}
	}
	return platforms, nil
}
