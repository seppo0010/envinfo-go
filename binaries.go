package envinfo

func GetBinaries() []*Item {
	items := []*Item{}
	nodeversion, _ := GetNodeVersion()
	if nodeversion != nil {
		items = append(items, nodeversion)
	}
	npmversion, _ := GetNpmVersion()
	if npmversion != nil {
		items = append(items, npmversion)
	}
	return items
}

func GetGoVersion() (*Item, error) {
	return GetItem("go", "Go", "version")
}

func GetNpmVersion() (*Item, error) {
	return GetItem("npm", "npm", "--version")
}
