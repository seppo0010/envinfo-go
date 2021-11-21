package envinfo

func GetServers() []*Item {
	return getItems([]func() (*Item, error){
		GetApacheVersion,
		GetNginxVersion,
	})
}

func GetApacheVersion() (*Item, error) {
	return NewGetItemBuilder("apachectl", "Apache").Flag("-v").Get()
}

func GetNginxVersion() (*Item, error) {
	return NewGetItemBuilder("nginx", "Nginx").Flag("-v").Stderr().Get()
}
