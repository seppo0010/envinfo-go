package envinfo

import (
	"bytes"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

type GetItemBuilder struct {
	executable   string
	name         string
	flag         string
	regex        *regexp.Regexp
	stdout       bool
	stderr       bool
	parseVersion func(string) string
}

func NewGetItemBuilder(executable, name string) *GetItemBuilder {
	var b *GetItemBuilder
	b = &GetItemBuilder{
		executable: executable,
		name:       name,
		flag:       "--version",
		regex:      versionRegex,
		stdout:     true,
		parseVersion: func(unparsed string) string {
			return b.regex.FindString(unparsed)
		},
	}
	return b
}

func (b *GetItemBuilder) Regex(regex *regexp.Regexp) *GetItemBuilder {
	b.regex = regex
	return b
}

func (b *GetItemBuilder) NoStdout() *GetItemBuilder {
	b.stdout = false
	return b
}

func (b *GetItemBuilder) Stderr() *GetItemBuilder {
	b.stderr = true
	return b
}

func (b *GetItemBuilder) Flag(flag string) *GetItemBuilder {
	b.flag = flag
	return b
}

func (b *GetItemBuilder) ParseVersion(parseVersion func(string) string) *GetItemBuilder {
	b.parseVersion = parseVersion
	return b
}

func (b *GetItemBuilder) Get() *Item {
	start := time.Now()
	which := b.executable
	version := ""
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.WithFields(log.Fields{
			"executable": b.executable,
			"name":       b.name,
			"flag":       b.flag,
		}).Debug("looking for executable")

		whichCmd := exec.Command("which", b.executable)
		whichBytes, err := whichCmd.Output()
		if err != nil {
			log.WithFields(log.Fields{
				"name": b.name,
			}).Warn("executable not found")
			return
		}
		which = strings.TrimSpace(string(whichBytes))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command(string(which), b.flag)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		_ = cmd.Run()
		log.WithFields(log.Fields{
			"name":   b.name,
			"cmd":    []string{string(which), b.flag},
			"stdout": string(outb.Bytes()),
			"stderr": string(errb.Bytes()),
		}).Debug("unparsed version")
		var parseBytes []byte
		if b.stdout {
			parseBytes = append(parseBytes, outb.Bytes()...)
		}
		if b.stderr {
			parseBytes = append(parseBytes, errb.Bytes()...)
		}
		version = strings.TrimSpace(b.parseVersion(string(parseBytes)))
	}()
	wg.Wait()
	log.WithFields(log.Fields{
		"name":     b.name,
		"duration": time.Now().Sub(start),
	}).Debug("parsed version")
	return &Item{
		Name:    b.name,
		Version: version,
		Path:    string(which),
	}
}

func GetItem(executable, name string) *Item {
	return NewGetItemBuilder(executable, name).Flag("--version").Get()
}

type ByName []*Item

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getItems(funcs []func() *Item) []*Item {
	results := make(chan (*Item), len(funcs))
	var wg sync.WaitGroup
	for _, f := range funcs {
		wg.Add(1)
		go func(f func() *Item) {
			defer wg.Done()
			res := f()
			results <- res
		}(f)
	}
	wg.Wait()
	close(results)

	items := make([]*Item, 0, len(funcs))
	for item := range results {
		if item != nil {
			items = append(items, item)
		}
	}

	sort.Sort(ByName(items))
	return items
}
