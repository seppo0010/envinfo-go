package envinfo

import (
	"os/exec"
	"regexp"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	stdout, _ := cmd.Output()
	return versionRegex.FindString(string(stdout)), nil
}
