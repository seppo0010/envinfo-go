package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"

	"github.com/seppo0010/envinfo-go"
)

type Options struct {
	Verbose        []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	System         bool   `long:"system" description:"Print general system info such as OS, CPU, Memory and Shell"`
	Languages      bool   `long:"languages" description:"Get version numbers of installed languages such as Java, Python, PHP, etc"`
	Binaries       bool   `long:"binaries" description:"Get version numbers of node, npm, watchman, etc"`
	Browsers       bool   `long:"browsers" description:"Get version numbers of installed web browsers"`
	Managers       bool   `long:"managers" description:"Get version numbers of installed package/dependency managers"`
	Utilities      bool   `long:"utilities" description:"Get version numbers of installed utilities"`
	Virtualization bool   `long:"virtualization" description:"Get version numbers of installed virtualization tools"`
	Servers        bool   `long:"servers" description:"Get version numbers of installed servers"`
	SDKs           bool   `long:"sdks" description:"Get version numbers of installed sdks"`
	IDEs           bool   `long:"ides" description:"Get version numbers of installed IDEs"`
	Databases      bool   `long:"databases" description:"Get version numbers of installed databases"`
	JSON           bool   `long:"json" description:"Print output in JSON format"`
	ShowNotFound   bool   `long:"showNotFound" description:"Don't filter out values marked 'Not Found'"`
	Version        bool   `long:"version" description:"Displays envinfo version"`
}

func main() {
	start := time.Now()
	var opts = Options{}
	_, err := flags.Parse(&opts)

	if flags.WroteHelp(err) {
		return
	}

	if err != nil {
		os.Exit(1)
		return
	}

	if opts.Version {
		fmt.Println(VERSION)
		return
	}

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
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

	if !opts.System && !opts.Languages && !opts.Binaries && !opts.Browsers && !opts.Managers && !opts.Utilities && !opts.Servers && !opts.Virtualization && !opts.SDKs && !opts.IDEs && !opts.Databases {
		opts.System = true
		opts.Languages = true
		opts.Binaries = true
		opts.Browsers = true
		opts.Managers = true
		opts.Utilities = true
		opts.Servers = true
		opts.Virtualization = true
		opts.SDKs = true
		opts.IDEs = true
		opts.Databases = true
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
	if opts.Managers {
		builder.Managers()
	}
	if opts.Browsers {
		builder.Browsers()
	}
	if opts.Utilities {
		builder.Utilities()
	}
	if opts.Virtualization {
		builder.Virtualization()
	}
	if opts.Servers {
		builder.Servers()
	}
	if opts.SDKs {
		builder.SDKs()
	}
	if opts.IDEs {
		builder.IDEs()
	}
	if opts.Databases {
		builder.Databases()
	}
	envInfo := builder.Build()

	if opts.JSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		encoder.Encode(envInfo)
		return
	}
	w := bufio.NewWriter(os.Stdout)
	if opts.System {
		PrintCLI("System", SystemToItems(envInfo.System), w, opts)
	}
	if opts.Languages {
		PrintCLI("Languages", envInfo.Languages, w, opts)
	}
	if opts.Binaries {
		PrintCLI("Binaries", envInfo.Binaries, w, opts)
	}
	if opts.Managers {
		PrintCLI("Managers", envInfo.Managers, w, opts)
	}
	if opts.Browsers {
		PrintCLI("Browsers", envInfo.Browsers, w, opts)
	}
	if opts.Utilities {
		PrintCLI("Utilities", envInfo.Utilities, w, opts)
	}
	if opts.Virtualization {
		PrintCLI("Virtualization", envInfo.Virtualization, w, opts)
	}
	if opts.Servers {
		PrintCLI("Servers", envInfo.Servers, w, opts)
	}
	if opts.SDKs {
		PrintSDKs("SDKs", envInfo.SDKs, w, opts)
	}
	if opts.IDEs {
		PrintCLI("IDEs", envInfo.IDEs, w, opts)
	}
	if opts.Databases {
		PrintCLI("Databases", envInfo.Databases, w, opts)
	}
	w.Flush()
	log.WithFields(log.Fields{
		"duration": time.Now().Sub(start),
	}).Debug("finishing")
}

func PrintCLI(title string, item []*envinfo.Item, w io.Writer, opts Options) {
	io.WriteString(w, "  ")
	io.WriteString(w, chalk.Underline.TextStyle(title))
	io.WriteString(w, "\n")
	for _, item := range item {
		if item.Version == "" && !opts.ShowNotFound {
			continue
		}
		io.WriteString(w, "    ")
		io.WriteString(w, item.Name)
		io.WriteString(w, ": ")
		if item.Version == "" {
			io.WriteString(w, "Not Found")
		} else {
			io.WriteString(w, item.Version)
		}
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

func PrintSDKs(title string, sdks *envinfo.SDKs, w io.Writer, opts Options) {
	io.WriteString(w, "  ")
	io.WriteString(w, chalk.Underline.TextStyle("SDKs:"))
	io.WriteString(w, "\n")
	if sdks.IOS != nil {
		io.WriteString(w, "    ")
		io.WriteString(w, chalk.Underline.TextStyle("iOS SDK:"))
		io.WriteString(w, "\n")
		io.WriteString(w, "      ")
		io.WriteString(w, "Platforms: ")
		io.WriteString(w, strings.Join(sdks.IOS, ", "))
		io.WriteString(w, "\n")
	}
	if sdks.Android != nil {
		io.WriteString(w, "    ")
		io.WriteString(w, chalk.Underline.TextStyle("Android SDK:"))
		io.WriteString(w, "\n")
		io.WriteString(w, "      ")
		io.WriteString(w, "API Levels: ")
		io.WriteString(w, strings.Join(sdks.Android.APILevels, ", "))
		io.WriteString(w, "\n")
		io.WriteString(w, "      ")
		io.WriteString(w, "Build Tools: ")
		io.WriteString(w, strings.Join(sdks.Android.BuildTools, ", "))
		io.WriteString(w, "\n")
		io.WriteString(w, "      ")
		io.WriteString(w, "System Images: ")
		io.WriteString(w, strings.Join(sdks.Android.SystemImages, ", "))
		io.WriteString(w, "\n")
	}
	io.WriteString(w, "\n")
}
