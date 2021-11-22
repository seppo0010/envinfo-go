//go:build windows
// +build windows

package envinfo

func GetAndroidStudioVersions() *Item {
	return nil
}

func GetIntelliJVersion() *Item {
	return nil
}

func GetPhpStormVersion() *Item {
	return nil
}

func GetWebStormVersion() *Item {
	return nil
}

func GetXcodeVersion() *Item {
	return &Item{
		Name:    "Xcode",
		Path:    "",
		Version: "",
	}
}

func GetSublimeVersion() *Item {
	return nil
}
