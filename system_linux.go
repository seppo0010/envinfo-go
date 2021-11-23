//go:build linux
// +build linux

package envinfo

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/acobaugh/osrelease"
)

func getOS() string {
	rel, err := osrelease.Read()
	if err != nil {
		return "Linux"
	}
	name := rel["NAME"]
	version := rel["VERSION"]
	cmd := exec.Command("uname", "-r")
	unameBytes, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("Linux %s %s", name, version)
	}
	uname := strings.Join(strings.Split(strings.TrimSpace(string(unameBytes)), ".")[:2], ".")
	return fmt.Sprintf("Linux %s %s %s", uname, name, version)
}
