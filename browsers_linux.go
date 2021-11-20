//go:build linux
// +build linux

package envinfo

func GetBraveVersion() (*Item, error) {
	item, _ := GetItem("brave", "Brave Browser", "--version")
	if item != nil {
		return item, nil
	}
	return GetItem("brave-browser", "Brave Browser", "--version")
}

func GetChromeVersion() (*Item, error) {
	return GetItem("google-chrome", "Chrome", "--version")
}

func GetChromeCanaryVersion() (*Item, error) {
	return &Item{
		Name:    "Google Chrome Canary",
		Path:    "",
		Version: "",
	}, nil
}

func GetChromiumVersion() (*Item, error) {
	return GetItem("chromium", "Chromium", "--version")
}

func GetEdgeVersion() (*Item, error) {
	return &Item{
		Name:    "Microsoft Edge",
		Path:    "",
		Version: "",
	}, nil
}

func GetFirefoxVersion() (*Item, error) {
	return GetItem("firefox", "Firefox", "--version")
}

func GetFirefoxNightlyVersion() (*Item, error) {
	return GetItem("firefox-trunk", "Firefox Nightly", "--version")
}

func GetFirefoxDeveloperEditionVersion() (*Item, error) {
	return &Item{
		Name:    "Firefox Developer Edition",
		Path:    "",
		Version: "",
	}, nil
}
