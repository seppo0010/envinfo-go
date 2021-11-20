//go:build darwin
// +build darwin

package envinfo

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var spaceRegex = regexp.MustCompile(`\s+`)

func getApplication(name, identifier string) (*Item, error) {
	log.WithFields(log.Fields{
		"identifier": identifier,
	}).Debug("fetching identifier version")

	appPathCmd := exec.Command("mdfind", fmt.Sprintf("kMDItemCFBundleIdentifier=='%s'", identifier))
	appPathBytes, err := appPathCmd.Output()
	if err != nil {
		log.WithFields(log.Fields{
			"identifier": identifier,
		}).Warn("identifier not found")
		return nil, err
	}

	path := spaceRegex.ReplaceAllString(strings.TrimSpace(string(appPathBytes)), " ")
	if path == "" {
		log.WithFields(log.Fields{
			"identifier": identifier,
		}).Warn("path not found")
		return nil, err
	}

	versionCmd := exec.Command("/usr/libexec/PlistBuddy", "-c", "Print CFBundleShortVersionString", fmt.Sprintf("%s/Contents/Info.plist", path))
	versionBytes, err := versionCmd.Output()
	if err != nil {
		log.WithFields(log.Fields{
			"identifier": identifier,
			"stderr":     err.(*exec.ExitError).Stderr,
		}).Warn("version not found")
		return nil, err
	}

	version := strings.TrimSpace(string(versionBytes))
	if version == "" {
		log.WithFields(log.Fields{
			"identifier": identifier,
		}).Warn("version not found")
		return nil, err
	}

	return &Item{
		Name:    name,
		Version: version,
		Path:    path,
	}, nil
}
