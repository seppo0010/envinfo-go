package envinfo

func GetBinaries() []*Item {
	return getItems([]func() *Item{GetNodeVersion, GetNpmVersion, GetYarnVersion, GetWatchmanVersion})
}

func GetGoVersion() *Item {
	return NewGetItemBuilder("go", "Go").Flag("version").Get()
}

func GetNpmVersion() *Item {
	return GetItem("npm", "npm")
}

func GetYarnVersion() *Item {
	return GetItem("yarn", "Yarn")
}

func GetWatchmanVersion() *Item {
	return GetItem("watchman", "Watchman")
}
