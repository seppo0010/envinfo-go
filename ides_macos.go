//go:build darwin
// +build darwin

package envinfo

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var idesBundleIdentifiers = map[string]string{
	"IntelliJ":     "com.jetbrains.intellij",
	"PhpStorm":     "com.jetbrains.PhpStorm",
	"Sublime Text": "com.sublimetext.3",
	"WebStorm":     "com.jetbrains.WebStorm",
}

func GetSublimeVersion() *Item {
	name := "Sublime Text"
	item := NewGetItemBuilder("subl", name).Regex(regexp.MustCompile(`\d+`)).Get()
	if item == nil || item.Version == "" {
		return getApplication(name, idesBundleIdentifiers[name])
	}
	return item
}

func GetIntelliJVersion() *Item {
	name := "IntelliJ"
	return getApplication(name, idesBundleIdentifiers[name])
}

func GetPhpStormVersion() *Item {
	name := "PhpStorm"
	return getApplication(name, idesBundleIdentifiers[name])
}

func GetWebStormVersion() *Item {
	name := "WebStorm"
	return getApplication(name, idesBundleIdentifiers[name])
}

func GetAndroidStudioVersions() *Item {
	for _, path := range []string{
		"/Applications/Android Studio.app",
		"/Applications/JetBrains Toolbox/Android Studio.app",
	} {
		version, _ := exec.Command("/usr/libexec/PlistBuddy", "-c", "Print:CFBundleShortVersionString", "-c", "Print:CFBundleVersion", fmt.Sprintf("%s%s", path, "/Contents/Info.plist")).Output()
		if len(version) > 0 {
			return &Item{
				Path:    path,
				Version: strings.TrimSpace(strings.Replace(string(version), "\n", " ", -1)),
				Name:    "Android Studio",
			}
		}
	}
	return nil
}

func GetXcodeVersion() *Item {
	b := NewGetItemBuilder("xcodebuild", "Xcode")
	return b.Flag("-version").ParseVersion(func(unparsed string) string {
		splitted := strings.Split(unparsed, "Build version ")
		if len(splitted) <= 1 {
			return b.regex.FindString(unparsed)
		}
		return fmt.Sprintf("%s/%s",
			b.regex.FindString(unparsed),
			splitted[1],
		)
	}).Get()
}
