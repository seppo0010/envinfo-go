package envinfo

func GetGoVersion() (*Item, error) {
	return GetItem("go", "Go", "version")
}

func GetNpmVersion() (*Item, error) {
	return GetItem("npm", "npm", "--version")
}
