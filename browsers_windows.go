//go:build windows
// +build windows

package envinfo

func GetBraveVersion() *Item {
	return nil
}

func GetChromeVersion() *Item {
	return nil
}

func GetChromeCanaryVersion() *Item {
	return nil
}

func GetChromiumVersion() *Item {
	return nil
}

func GetEdgeVersion() *Item {
	return nil
}

func GetFirefoxVersion() *Item {
	return nil
}

func GetFirefoxNightlyVersion() *Item {
	return nil
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
