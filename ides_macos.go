//go:build darwin
// +build darwin

package envinfo

import (
	"fmt"
	"os/exec"
	"strings"
)

var idesBundleIdentifiers = map[string]string{
	"IntelliJ":     "com.jetbrains.intellij",
	"PhpStorm":     "com.jetbrains.PhpStorm",
	"Sublime Text": "com.sublimetext.3",
	"WebStorm":     "com.jetbrains.WebStorm",
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
