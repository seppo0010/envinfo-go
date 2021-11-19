package envinfo

import (
	"os/exec"
	"regexp"
	"strings"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

type Item struct {
	Name    string
	Version string
	Path    string
}

func GetItem(executable, name, flag string) (*Item, error) {
	cmd := exec.Command("which", executable)
	whichBytes, err := cmd.Output()
	if err != nil {
		return nil, nil
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
