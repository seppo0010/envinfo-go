//go:build linux
// +build linux

package envinfo

func GetBraveVersion() (*Item, error) {
	item, _ := GetItem("brave", "Brave Browser")
	if item.Version != "" {
		return item, nil
	}
	return GetItem("brave-browser", "Brave Browser")
}

func GetChromeVersion() (*Item, error) {
	return GetItem("google-chrome", "Chrome")
}

func GetChromeCanaryVersion() (*Item, error) {
	return &Item{
		Name:    "Google Chrome Canary",
		Path:    "",
		Version: "",
	}, nil
}

func GetChromiumVersion() (*Item, error) {
	return GetItem("chromium", "Chromium")
}

func GetEdgeVersion() (*Item, error) {
	return &Item{
		Name:    "Microsoft Edge",
		Path:    "",
		Version: "",
	}, nil
}

func GetFirefoxVersion() (*Item, error) {
	return GetItem("firefox", "Firefox")
}

func GetFirefoxNightlyVersion() (*Item, error) {
	return GetItem("firefox-trunk", "Firefox Nightly")
}

func GetFirefoxDeveloperEditionVersion() (*Item, error) {
	return &Item{
		Name:    "Firefox Developer Edition",
		Path:    "",
		Version: "",
	}, nil
}

func GetSafariVersion() (*Item, error) {
	return &Item{
		Name:    "Safari",
		Path:    "",
		Version: "",
	}, nil
}
