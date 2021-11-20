package envinfo

import "sync"

func GetBinaries() []*Item {
	funcs := []func() (*Item, error){GetNodeVersion, GetNpmVersion, GetYarnVersion}
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

	items := make([]*Item, len(funcs))
	for i, _ := range items {
		items[i] = <-results
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
