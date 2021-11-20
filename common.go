package envinfo

import (
	"os/exec"
	"regexp"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)

func GetItemRegex(executable, name, flag string, regex *regexp.Regexp) (*Item, error) {
	log.WithFields(log.Fields{
		"executable": executable,
		"name":       name,
		"flag":       flag,
	}).Debug("looking for executable")

	cmd := exec.Command("which", executable)
	whichBytes, err := cmd.Output()
	if err != nil {
		log.WithFields(log.Fields{
			"stderr": string(err.(*exec.ExitError).Stderr),
		}).Warn("executable not found")
		return nil, err
	}
	which := strings.TrimSpace(string(whichBytes))
	cmd = exec.Command(string(which), flag)
	stdout, _ := cmd.Output()
	return &Item{
		Name:    name,
		Version: strings.TrimSpace(regex.FindString(string(stdout))),
		Path:    string(which),
	}, nil
}
func GetItem(executable, name, flag string) (*Item, error) {
	return GetItemRegex(executable, name, flag, versionRegex)
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
