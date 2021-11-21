package envinfo

func GetManagers() []*Item {
	return getItems([]func() (*Item, error){
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

func GetAptVersion() (*Item, error) {
	return GetItem("apt", "Apt")
}

func GetCargoVersion() (*Item, error) {
	return GetItem("cargo", "Cargo")
}

func GetCocoaPodsVersion() (*Item, error) {
	return GetItem("pod", "CocoaPods")
}

func GetComposerVersion() (*Item, error) {
	return GetItem("composer", "Composer")
}

func GetGradleVersion() (*Item, error) {
	return GetItem("gradle", "Gradle")
}

func GetHomebrewVersion() (*Item, error) {
	return GetItem("brew", "Homebrew")
}

func GetMavenVersion() (*Item, error) {
	return GetItem("mvn", "Maven")
}

func GetPip2Version() (*Item, error) {
	return GetItem("pip2", "pip2")
}

func GetPip3Version() (*Item, error) {
	return GetItem("pip3", "pip3")
}

func GetRubyGemsVersion() (*Item, error) {
	return GetItem("gem", "RubyGems")
}

func GetYumVersion() (*Item, error) {
	return GetItem("yum", "Yum")
}
