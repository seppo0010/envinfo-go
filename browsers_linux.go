//go:build linux
// +build linux

package envinfo

func GetBraveVersion() *Item {
	item := GetItem("brave", "Brave Browser")
	if item.Version != "" {
		return item
	}
	return GetItem("brave-browser", "Brave Browser")
}

func GetChromeVersion() *Item {
	return GetItem("google-chrome", "Chrome")
}

func GetChromeCanaryVersion() *Item {
	return &Item{
		Name:    "Google Chrome Canary",
		Path:    "",
		Version: "",
	}
}

func GetChromiumVersion() *Item {
	return GetItem("chromium", "Chromium")
}

func GetEdgeVersion() *Item {
	return &Item{
		Name:    "Microsoft Edge",
		Path:    "",
		Version: "",
	}
}

func GetFirefoxVersion() *Item {
	return GetItem("firefox", "Firefox")
}

func GetFirefoxNightlyVersion() *Item {
	return GetItem("firefox-trunk", "Firefox Nightly")
}

func GetFirefoxDeveloperEditionVersion() *Item {
	return &Item{
		Name:    "Firefox Developer Edition",
		Path:    "",
		Version: "",
	}
}

func GetSafariVersion() *Item {
	return &Item{
		Name:    "Safari",
		Path:    "",
		Version: "",
	}
}
