//go:build darwin
// +build darwin

package envinfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func getOS() string {
	cmd := exec.Command("sw_vers", "-productVersion")
	output, err := cmd.Output()
	if err != nil {
		return "macOS"
	}
	return fmt.Sprintf("macOS %s", strings.TrimSpace(string(output)))
}
