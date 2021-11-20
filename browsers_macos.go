//go:build darwin
// +build darwin

package envinfo

var browserBundleIdentifiers = map[string]string{
	"Brave Browser":             "com.brave.Browser",
	"Chrome":                    "com.google.Chrome",
	"Chrome Canary":             "com.google.Chrome.canary",
	"Firefox":                   "org.mozilla.firefox",
	"Firefox Developer Edition": "org.mozilla.firefoxdeveloperedition",
	"Firefox Nightly":           "org.mozilla.nightly",
	"Microsoft Edge":            "com.microsoft.edgemac",
	"Safari":                    "com.apple.Safari",
	"Safari Technology Preview": "com.apple.SafariTechnologyPreview",
}

func GetBraveVersion() (*Item, error) {
	name := "Brave Browser"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromeVersion() (*Item, error) {
	name := "Chrome"
	return getApplication(name, browserBundleIdentifiers[name])
}
