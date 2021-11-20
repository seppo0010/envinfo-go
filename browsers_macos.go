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

func GetBraveVersion() (*Item, error) {
	name := "Brave Browser"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromeVersion() (*Item, error) {
	name := "Chrome"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromeCanaryVersion() (*Item, error) {
	name := "Chrome Canary"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetChromiumVersion() (*Item, error) {
	name := "Chromium"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetEdgeVersion() (*Item, error) {
	name := "Microsoft Edge"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxVersion() (*Item, error) {
	name := "Firefox"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxDeveloperEditionVersion() (*Item, error) {
	name := "Firefox Developer Edition"
	return getApplication(name, browserBundleIdentifiers[name])
}

func GetFirefoxNightlyVersion() (*Item, error) {
	name := "Firefox Nightly"
	return getApplication(name, browserBundleIdentifiers[name])
}
