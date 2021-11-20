package envinfo

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

type GetItemBuilder struct {
	executable string
	name       string
	flag       string
	regex      *regexp.Regexp
	stdout     bool
	stderr     bool
}

func NewGetItemBuilder(executable, name, flag string) *GetItemBuilder {
	return &GetItemBuilder{
		executable: executable,
		name:       name,
		flag:       flag,
		regex:      versionRegex,
		stdout:     true,
	}
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

func (b *GetItemBuilder) Get() (*Item, error) {
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
		return nil, err
	}
	which := strings.TrimSpace(string(whichBytes))
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
	version := strings.TrimSpace(b.regex.FindString(string(parseBytes)))
	return &Item{
		Name:    b.name,
		Version: version,
		Path:    string(which),
	}, nil
}

func GetItemRegex(executable, name, flag string, regex *regexp.Regexp) (*Item, error) {
	return NewGetItemBuilder(executable, name, flag).Regex(regex).Get()
}
func GetItem(executable, name, flag string) (*Item, error) {
	return NewGetItemBuilder(executable, name, flag).Get()
}

func getItems(funcs []func() (*Item, error)) []*Item {
	results := make(chan (*Item), len(funcs))
	var wg sync.WaitGroup
	for _, f := range funcs {
		wg.Add(1)
		go func(f func() (*Item, error)) {
			defer wg.Done()
			res, _ := f()
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

	return items
}
