package main

import (
	"io"
	"os"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

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
	items := []*envinfo.Item{
		&envinfo.Item{
			Name:    "OS",
			Version: system.OS,
		},
		&envinfo.Item{
			Name:    "CPU",
			Version: system.CPU,
		},
		&envinfo.Item{
			Name:    "Memory",
			Version: system.Memory,
		},
		&envinfo.Item{
			Name:    "Shell",
			Version: system.Shell,
		},
	}
	if system.Container != "" {
		items = append(items, &envinfo.Item{
			Name:    "Container",
			Version: system.Container,
		})
	}
	return items
}

func main() {
	var opts = Options{}
	_, err := flags.Parse(&opts)

	if flags.WroteHelp(err) {
		return
	}

	if err != nil {
		os.Exit(1)
		return
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	if len(opts.Verbose) > 3 {
		log.SetLevel(log.DebugLevel)
	} else if len(opts.Verbose) > 2 {
		log.SetLevel(log.InfoLevel)
	} else if len(opts.Verbose) > 1 {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

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
