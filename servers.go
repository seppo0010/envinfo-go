package envinfo

func GetServers() []*Item {
	return getItems([]func() *Item{
		GetApacheVersion,
		GetNginxVersion,
	})
}

func GetApacheVersion() *Item {
	return NewGetItemBuilder("apachectl", "Apache").Flag("-v").Get()
}

func GetNginxVersion() *Item {
	return NewGetItemBuilder("nginx", "Nginx").Flag("-v").Stderr().Get()
}
