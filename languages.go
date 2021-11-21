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
		GetRubyVersion,
		GetRustVersion,
		GetScalaVersion,
	})
}

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node")
}

func GetBashVersion() (*Item, error) {
	return GetItem("bash", "Bash")
}

func GetElixirVersion() (*Item, error) {
	versionRegex := regexp.MustCompile(`[Elixir]+\s[\d+.[\d+|.]+`)
	item, err := NewGetItemBuilder("elixir", "Elixir").Regex(versionRegex).Get()
	if err != nil {
		return nil, err
	}
	if len(item.Version) > 7 {
		item.Version = item.Version[7:]
	}
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
	regex := regexp.MustCompile(`\d+\.?[\w+|.|_|-]+`)
	return NewGetItemBuilder("javac", "Java").Flag("-version").Regex(regex).Stderr().Get()
}

func GetPerlVersion() (*Item, error) {
	return GetItem("perl", "Perl")
}

func GetPHPVersion() (*Item, error) {
	return GetItem("php", "PHP")
}

func GetProtocVersion() (*Item, error) {
	return GetItem("protoc", "Protoc")
}

func GetPythonVersion() (*Item, error) {
	return NewGetItemBuilder("python", "Python").Stderr().NoStdout().Get()
}

func GetPython3Version() (*Item, error) {
	return GetItem("python3", "Python3")
}

func GetRVersion() (*Item, error) {
	return GetItem("R", "R")
}

func GetRubyVersion() (*Item, error) {
	return GetItem("ruby", "Ruby")
}

func GetRustVersion() (*Item, error) {
	return GetItem("rustc", "Rust")
}

func GetScalaVersion() (*Item, error) {
	return GetItem("scalac", "Scala")
}
