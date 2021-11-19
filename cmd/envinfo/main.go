package main

import (
	"io"
	"os"

	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

func GetLanguages() map[string]string {
	goversion, _ := envinfo.GetGoVersion()
	return map[string]string{
		"Go": goversion,
	}
}

func GetBinaries() map[string]string {
	nodeversion, _ := envinfo.GetNodeVersion()
	return map[string]string{
		"Node": nodeversion,
	}
}

type Versions struct {
	Languages map[string]string
	Binaries  map[string]string
}

func NewVersions() *Versions {
	return &Versions{
		Languages: map[string]string{},
		Binaries:  map[string]string{},
	}
}

func main() {
	versions := NewVersions()
	versions.Languages = GetLanguages()
	versions.Binaries = GetBinaries()
	PrintCLI("Languages", versions.Languages, os.Stdout)
	PrintCLI("Binaries", versions.Binaries, os.Stdout)
}

func PrintCLI(title string, versions map[string]string, w io.Writer) {
	io.WriteString(w, "  ")
	io.WriteString(w, chalk.Underline.TextStyle(title))
	io.WriteString(w, "\n")
	for name, version := range versions {
		io.WriteString(w, "    ")
		io.WriteString(w, name)
		io.WriteString(w, ": ")
		io.WriteString(w, version)
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "\n")
}
