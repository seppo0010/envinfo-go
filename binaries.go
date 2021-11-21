package envinfo

func GetBinaries() []*Item {
	return getItems([]func() (*Item, error){GetNodeVersion, GetNpmVersion, GetYarnVersion, GetWatchmanVersion})
}

func GetGoVersion() (*Item, error) {
	return NewGetItemBuilder("go", "Go").Flag("version").Get()
}

func GetNpmVersion() (*Item, error) {
	return GetItem("npm", "npm")
}

func GetYarnVersion() (*Item, error) {
	return GetItem("yarn", "Yarn")
}

func GetWatchmanVersion() (*Item, error) {
	return GetItem("watchman", "Watchman")
}
