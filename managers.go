package envinfo

func GetManagers() []*Item {
	return getItems([]func() *Item{
		GetAptVersion,
		GetCargoVersion,
		GetCocoaPodsVersion,
		GetComposerVersion,
		GetGradleVersion,
		GetHomebrewVersion,
		GetMavenVersion,
		GetPip2Version,
		GetPip3Version,
		GetRubyGemsVersion,
		GetYumVersion,
	})
}

func GetAptVersion() *Item {
	return GetItem("apt", "Apt")
}

func GetCargoVersion() *Item {
	return GetItem("cargo", "Cargo")
}

func GetCocoaPodsVersion() *Item {
	return GetItem("pod", "CocoaPods")
}

func GetComposerVersion() *Item {
	return GetItem("composer", "Composer")
}

func GetGradleVersion() *Item {
	return GetItem("gradle", "Gradle")
}

func GetHomebrewVersion() *Item {
	return GetItem("brew", "Homebrew")
}

func GetMavenVersion() *Item {
	return GetItem("mvn", "Maven")
}

func GetPip2Version() *Item {
	return GetItem("pip2", "pip2")
}

func GetPip3Version() *Item {
	return GetItem("pip3", "pip3")
}

func GetRubyGemsVersion() *Item {
	return GetItem("gem", "RubyGems")
}

func GetYumVersion() *Item {
	return GetItem("yum", "Yum")
}
