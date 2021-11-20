package envinfo

import "regexp"

func GetLanguages() []*Item {
	return getItems([]func() (*Item, error){GetGoVersion, GetNodeVersion, GetBashVersion, GetElixirVersion})
}

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node", "--version")
}

func GetBashVersion() (*Item, error) {
	return GetItem("bash", "Bash", "--version")
}

func GetElixirVersion() (*Item, error) {
	versionRegex := regexp.MustCompile(`[Elixir]+\s[\d+.[\d+|.]+`)
	item, err := GetItemRegex("elixir", "Elixir", "--version", versionRegex)
	if err != nil {
		return nil, err
	}
	item.Version = item.Version[7:]
	return item, nil
}
