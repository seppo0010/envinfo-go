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

func GetIntelliJVersion() (*Item, error) {
	name := "IntelliJ"
	return getApplication(name, idesBundleIdentifiers[name])
}

func GetAndroidStudioVersions() (*Item, error) {
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
			}, nil
		}
	}
	return nil, nil
}

/*
   .then(version => {
     if (!version) {
       return utils.run(
         utils.generatePlistBuddyCommand(
           path.join(
             '~',
             'Applications',
             'JetBrains\\ Toolbox',
             'Android\\ Studio.app',
             'Contents',
             'Info.plist'
           ),
           ['CFBundleShortVersionString', 'CFBundleVersion']
         )
       );
     }
     return version;
   })
   .then(version => version.split('\n').join(' '));
*/
