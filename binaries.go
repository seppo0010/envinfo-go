package envinfo

func GetBinaries() []*Item {
	return getItems([]func() (*Item, error){GetNodeVersion, GetNpmVersion, GetYarnVersion, GetWatchmanVersion})
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
