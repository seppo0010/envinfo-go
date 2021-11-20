package envinfo

func GetBrowsers() []*Item {
	return getItems([]func() (*Item, error){
		GetBraveVersion,
		GetChromeVersion,
		GetChromeCanaryVersion,
		GetChromiumVersion,
		GetEdgeVersion,
		GetFirefoxVersion,
		GetFirefoxDeveloperEditionVersion,
		GetFirefoxNightlyVersion,
		GetSafariVersion,
	})
}
