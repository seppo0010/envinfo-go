package envinfo

func GetLanguages() []*Item {
	items := []*Item{}
	goversion, _ := GetGoVersion()
	if goversion != nil {
		items = append(items, goversion)
	}
	return items
}

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node", "--version")
}
