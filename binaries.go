package envinfo

func GetGoVersion() (*Item, error) {
	return GetItem("go", "Go", "version")
}
