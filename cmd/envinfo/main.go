package main

import (
	"io"
	"os"

	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

type EnvInfo struct {
	Languages []*envinfo.Item
	Binaries  []*envinfo.Item
	System    *envinfo.System
}

func NewEnvInfo() *EnvInfo {
	return &EnvInfo{
		Languages: []*envinfo.Item{},
		Binaries:  []*envinfo.Item{},
	}
}

func SystemToItems(system *envinfo.System) []*envinfo.Item {
	if system == nil {
		return []*envinfo.Item{}
	}
	return []*envinfo.Item{
		&envinfo.Item{
			Name:    "OS",
			Version: system.OS,
		},
	}
}

func main() {
	envInfo := NewEnvInfo()
	envInfo.Languages = envinfo.GetLanguages()
	envInfo.Binaries = envinfo.GetBinaries()
	envInfo.System = envinfo.GetSystem()
	PrintCLI("System", SystemToItems(envInfo.System), os.Stdout)
	PrintCLI("Languages", envInfo.Languages, os.Stdout)
	PrintCLI("Binaries", envInfo.Binaries, os.Stdout)
}

func PrintCLI(title string, item []*envinfo.Item, w io.Writer) {
	io.WriteString(w, "  ")
	io.WriteString(w, chalk.Underline.TextStyle(title))
	io.WriteString(w, "\n")
	for _, item := range item {
		io.WriteString(w, "    ")
		io.WriteString(w, item.Name)
		io.WriteString(w, ": ")
		io.WriteString(w, item.Version)
		if item.Path != "" {
			io.WriteString(w, " - ")
			io.WriteString(w, item.Path)
		}
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "\n")
}
