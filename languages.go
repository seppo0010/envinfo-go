package envinfo

func GetLanguages() []*Item {
	return getItems([]func() (*Item, error){GetNodeVersion, GetBashVersion})
}

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node", "--version")
}

func GetBashVersion() (*Item, error) {
	return GetItem("bash", "Bash", "--version")
}
