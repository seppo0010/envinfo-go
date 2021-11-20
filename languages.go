package envinfo

import (
	"os/exec"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetLanguages() []*Item {
	return getItems([]func() (*Item, error){GetGoVersion, GetNodeVersion, GetBashVersion, GetElixirVersion, GetErlangVersion})
}

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node", "--version")
}

func GetBashVersion() (*Item, error) {
	return GetItem("bash", "Bash", "--version")
}

func GetElixirVersion() (*Item, error) {
	versionRegex := regexp.MustCompile(`[Elixir]+\s[\d+.[\d+|.]+`)
	item, err := GetItemRegex("elixir", "Elixir", "--version", versionRegex)
	if err != nil {
		return nil, err
	}
	item.Version = item.Version[7:]
	return item, nil
}

func GetErlangVersion() (*Item, error) {
	name := "Erlang"
	executable := "erl"
	log.WithFields(log.Fields{
		"executable": executable,
		"name":       name,
	}).Debug("looking for executable")

	cmd := exec.Command("which", "erl")
	whichBytes, err := cmd.Output()
	if err != nil {
		log.WithFields(log.Fields{
			"stderr": string(err.(*exec.ExitError).Stderr),
		}).Warn("executable not found")
		return nil, err
	}
	which := strings.TrimSpace(string(whichBytes))
	cmd = exec.Command(string(which), "-eval", "{ok, Version} = file:read_file(filename:join([code:root_dir(), 'releases', erlang:system_info(otp_release), 'OTP_VERSION'])), io:fwrite(Version), halt().", "-noshell")
	stdout, _ := cmd.Output()
	return &Item{
		Name:    name,
		Version: strings.TrimSpace(string(stdout)),
		Path:    string(which),
	}, nil
}
