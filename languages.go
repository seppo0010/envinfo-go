package envinfo

import (
	"os/exec"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetLanguages() []*Item {
	return getItems([]func() (*Item, error){
		GetGoVersion,
		GetNodeVersion,
		GetBashVersion,
		GetElixirVersion,
		GetErlangVersion,
		GetJavaVersion,
		GetPerlVersion,
		GetPHPVersion,
		GetProtocVersion,
		GetPythonVersion,
		GetPython3Version,
		GetRVersion,
	})
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
		log.WithFields(log.Fields{}).Warn("executable not found")
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

func GetJavaVersion() (*Item, error) {
	versionRegex := regexp.MustCompile(`\d+\.?[\w+|.|_|-]+`)
	executable, name, flag, regex := "javac", "Java", "-version", versionRegex
	return NewGetItemBuilder(executable, name, flag).Regex(regex).Stderr().Get()
}

func GetPerlVersion() (*Item, error) {
	return GetItem("perl", "Perl", "--version")
}

func GetPHPVersion() (*Item, error) {
	return GetItem("php", "PHP", "--version")
}

func GetProtocVersion() (*Item, error) {
	return GetItem("protoc", "Protoc", "--version")
}

func GetPythonVersion() (*Item, error) {
	executable, name, flag := "python", "Python", "--version"
	return NewGetItemBuilder(executable, name, flag).Stderr().NoStdout().Get()
}

func GetPython3Version() (*Item, error) {
	return GetItem("python3", "Python3", "--version")
}

func GetRVersion() (*Item, error) {
	return GetItem("R", "R", "--version")
}
