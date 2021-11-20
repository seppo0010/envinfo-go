package envinfo

import "sync"

func GetBinaries() []*Item {
	funcs := []func() (*Item, error){GetNodeVersion, GetNpmVersion, GetYarnVersion, GetWatchmanVersion}
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

func GetGoVersion() (*Item, error) {
	return GetItem("go", "Go", "version")
}

func GetNpmVersion() (*Item, error) {
	return GetItem("npm", "npm", "--version")
}

func GetYarnVersion() (*Item, error) {
	return GetItem("yarn", "Yarn", "--version")
}

func GetWatchmanVersion() (*Item, error) {
	return GetItem("watchman", "Watchman", "--version")
}
