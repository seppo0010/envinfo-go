package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

type Options struct {
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	System    bool   `long:"system" description:"Print general system info such as OS, CPU, Memory and Shell"`
	Languages bool   `long:"languages" description:"Get version numbers of installed languages such as Java, Python, PHP, etc"`
	Binaries  bool   `long:"binaries" description:"Get version numbers of node, npm, watchman, etc"`
	Browsers  bool   `long:"browsers" description:"Get version numbers of installed web browsers"`
	JSON      bool   `long:"json" description:"Print output in JSON format"`
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

	if !opts.System && !opts.Languages && !opts.Binaries && !opts.Browsers {
		opts.System = true
		opts.Languages = true
		opts.Binaries = true
		opts.Browsers = true
	}

	builder := envinfo.NewEnvInfoBuilder()
	if opts.System {
		builder.System()
	}
	if opts.Languages {
		builder.Languages()
	}
	if opts.Binaries {
		builder.Binaries()
	}
	if opts.Browsers {
		builder.Browsers()
	}
	envInfo := builder.Build()

	if opts.JSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		encoder.Encode(envInfo)
		return
	}
	if opts.System {
		PrintCLI("System", SystemToItems(envInfo.System), os.Stdout)
	}
	if opts.Languages {
		PrintCLI("Languages", envInfo.Languages, os.Stdout)
	}
	if opts.Binaries {
		PrintCLI("Binaries", envInfo.Binaries, os.Stdout)
	}
	if opts.Browsers {
		PrintCLI("Browsers", envInfo.Browsers, os.Stdout)
	}
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
