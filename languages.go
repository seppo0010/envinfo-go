package envinfo

func GetNodeVersion() (*Item, error) {
	return GetItem("node", "Node", "--version")
}
