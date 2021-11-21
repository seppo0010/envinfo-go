package envinfo

func GetServers() []*Item {
	return getItems([]func() (*Item, error){
		GetApacheVersion,
		GetNginxVersion,
	})
}

func GetApacheVersion() (*Item, error) {
	return GetItem("apachectl", "Apache", "-v")
}

func GetNginxVersion() (*Item, error) {
	executable, name, flag := "nginx", "Nginx", "-v"
	return NewGetItemBuilder(executable, name, flag).Stderr().Get()
}
