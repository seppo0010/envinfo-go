package envinfo

import (
	"os/exec"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetLanguages() []*Item {
	return getItems([]func() *Item{
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

func GetNodeVersion() *Item {
	return GetItem("node", "Node")
}

func GetBashVersion() *Item {
	return GetItem("bash", "Bash")
}

func GetElixirVersion() *Item {
	versionRegex := regexp.MustCompile(`[Elixir]+\s[\d+.[\d+|.]+`)
	item := NewGetItemBuilder("elixir", "Elixir").Regex(versionRegex).Get()
	if item == nil {
		return nil
	}
	if len(item.Version) > 7 {
		item.Version = item.Version[7:]
	}
	return item
}

func GetErlangVersion() *Item {
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
		return nil
	}
	which := strings.TrimSpace(string(whichBytes))
	cmd = exec.Command(string(which), "-eval", "{ok, Version} = file:read_file(filename:join([code:root_dir(), 'releases', erlang:system_info(otp_release), 'OTP_VERSION'])), io:fwrite(Version), halt().", "-noshell")
	stdout, _ := cmd.Output()
	return &Item{
		Name:    name,
		Version: strings.TrimSpace(string(stdout)),
		Path:    string(which),
	}
}

func GetJavaVersion() *Item {
	regex := regexp.MustCompile(`\d+\.?[\w+|.|_|-]+`)
	return NewGetItemBuilder("javac", "Java").Flag("-version").Regex(regex).Stderr().Get()
}

func GetPerlVersion() *Item {
	return GetItem("perl", "Perl")
}

func GetPHPVersion() *Item {
	return GetItem("php", "PHP")
}

func GetProtocVersion() *Item {
	return GetItem("protoc", "Protoc")
}

func GetPythonVersion() *Item {
	return NewGetItemBuilder("python", "Python").Stderr().NoStdout().Get()
}

func GetPython3Version() *Item {
	return GetItem("python3", "Python3")
}

func GetRVersion() *Item {
	return GetItem("R", "R")
}

func GetRubyVersion() *Item {
	return GetItem("ruby", "Ruby")
}

func GetRustVersion() *Item {
	return GetItem("rustc", "Rust")
}

func GetScalaVersion() *Item {
	return GetItem("scalac", "Scala")
}
