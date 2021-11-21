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
	return GetItem("apt", "Apt", "--version")
}

func GetCargoVersion() (*Item, error) {
	return GetItem("cargo", "Cargo", "--version")
}

func GetCocoaPodsVersion() (*Item, error) {
	return GetItem("pod", "CocoaPods", "--version")
}

func GetComposerVersion() (*Item, error) {
	return GetItem("composer", "Composer", "--version")
}

func GetGradleVersion() (*Item, error) {
	return GetItem("gradle", "Gradle", "--version")
}

func GetHomebrewVersion() (*Item, error) {
	return GetItem("brew", "Homebrew", "--version")
}

func GetMavenVersion() (*Item, error) {
	return GetItem("mvn", "Maven", "--version")
}

func GetPip2Version() (*Item, error) {
	return GetItem("pip2", "pip2", "--version")
}

func GetPip3Version() (*Item, error) {
	return GetItem("pip3", "pip3", "--version")
}

func GetRubyGemsVersion() (*Item, error) {
	return GetItem("gem", "RubyGems", "--version")
}

func GetYumVersion() (*Item, error) {
	return GetItem("yum", "Yum", "--version")
}
