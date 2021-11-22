package envinfo

func GetBrowsers() []*Item {
	return getItems([]func() *Item{
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
