//go:build windows
// +build windows

package envinfo

import (
	"fmt"

	"github.com/acobaugh/osrelease"
)

func getOS() string {
	rel, err := osrelease.Read()
	if err != nil {
		return "Windows"
	}
	name := rel["NAME"]
	version := rel["VERSION"]
	return fmt.Sprintf("Windows %s %s", name, version)
}
