package envinfo

import (
	"os/exec"
)

func GetNodeVersion() (string, error) {
	cmd := exec.Command("node", "--version")
	stdout, _ := cmd.Output()
	return versionRegex.FindString(string(stdout)), nil
}
