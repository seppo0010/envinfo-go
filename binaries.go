package envinfo

import (
	"os/exec"
)

func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	stdout, _ := cmd.Output()
	return versionRegex.FindString(string(stdout)), nil
}
