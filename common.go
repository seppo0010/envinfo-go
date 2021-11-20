package envinfo

import (
	"os/exec"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

func GetItem(executable, name, flag string) (*Item, error) {
	log.WithFields(log.Fields{
		"executable": executable,
		"name":       name,
		"flag":       flag,
	}).Debug("looking for executable")

	cmd := exec.Command("which", executable)
	whichBytes, err := cmd.Output()
	if err != nil {
		log.WithFields(log.Fields{
			"stderr": string(err.(*exec.ExitError).Stderr),
		}).Warn("executable not found")
		return nil, err
	}
	which := strings.TrimSpace(string(whichBytes))
	cmd = exec.Command(string(which), flag)
	stdout, _ := cmd.Output()
	return &Item{
		Name:    name,
		Version: strings.TrimSpace(versionRegex.FindString(string(stdout))),
		Path:    string(which),
	}, nil
}
