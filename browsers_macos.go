//go:build darwin
// +build darwin

package envinfo

var browserBundleIdentifiers = map[string]string{
	"Brave Browser":             "com.brave.Browser",
	"Chrome":                    "com.google.Chrome",
	"Chrome Canary":             "com.google.Chrome.canary",
	"Chromium":                  "org.chromium.Chromium",
	"Firefox":                   "org.mozilla.firefox",
	"Firefox Developer Edition": "org.mozilla.firefoxdeveloperedition",
	"Firefox Nightly":           "org.mozilla.nightly",
	"Microsoft Edge":            "com.microsoft.edgemac",
	"Safari":                    "com.apple.Safari",
	"Safari Technology Preview": "com.apple.SafariTechnologyPreview",
}

func GetBraveVersion() *Item {
	name := "Brave Browser"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromeVersion() *Item {
	name := "Chrome"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromeCanaryVersion() *Item {
	name := "Chrome Canary"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromiumVersion() *Item {
	name := "Chromium"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetEdgeVersion() *Item {
	name := "Microsoft Edge"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxVersion() *Item {
	name := "Firefox"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxDeveloperEditionVersion() *Item {
	name := "Firefox Developer Edition"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxNightlyVersion() *Item {
	name := "Firefox Nightly"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetSafariVersion() *Item {
	name := "Safari"
	return getApplication(name, browserBundleIdentifiers[name])
}
