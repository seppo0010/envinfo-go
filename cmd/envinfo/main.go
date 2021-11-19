package main

import (
	"io"
	"os"

	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

func GetLanguages() []*envinfo.Item {
	items := []*envinfo.Item{}
	goversion, _ := envinfo.GetGoVersion()
	if goversion != nil {
		items = append(items, goversion)
	}
	return items
}

func GetBinaries() []*envinfo.Item {
	items := []*envinfo.Item{}
	nodeversion, _ := envinfo.GetNodeVersion()
	if nodeversion != nil {
		items = append(items, nodeversion)
	}
	npmversion, _ := envinfo.GetNpmVersion()
	if npmversion != nil {
		items = append(items, npmversion)
	}
	return items
}

type Versions struct {
	Languages []*envinfo.Item
	Binaries  []*envinfo.Item
}

func NewVersions() *Versions {
	return &Versions{
		Languages: []*envinfo.Item{},
		Binaries:  []*envinfo.Item{},
	}
}

func main() {
	versions := NewVersions()
	versions.Languages = GetLanguages()
	versions.Binaries = GetBinaries()
	PrintCLI("Languages", versions.Languages, os.Stdout)
	PrintCLI("Binaries", versions.Binaries, os.Stdout)
}

func PrintCLI(title string, versions []*envinfo.Item, w io.Writer) {
	io.WriteString(w, "  ")
	io.WriteString(w, chalk.Underline.TextStyle(title))
	io.WriteString(w, "\n")
	for _, item := range versions {
		io.WriteString(w, "    ")
		io.WriteString(w, item.Name)
		io.WriteString(w, ": ")
		io.WriteString(w, item.Version)
		io.WriteString(w, " - ")
		io.WriteString(w, item.Path)
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "\n")
}
