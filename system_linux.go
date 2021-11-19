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

/*

   version = utils.run('cat /etc/os-release').then(v => {
     const distro = (v || '').match(/NAME="(.+)"/) || '';
     const versionInfo = (v || '').match(/VERSION="(.+)"/) || ['', ''];
     const versionStr = versionInfo !== null ? versionInfo[1] : '';
     return `${distro[1]} ${versionStr}`.trim() || '';
   });
*/
