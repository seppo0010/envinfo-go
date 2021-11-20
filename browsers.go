package envinfo

func GetBrowsers() []*Item {
	return getItems([]func() (*Item, error){
		GetBraveVersion,
	})
}
