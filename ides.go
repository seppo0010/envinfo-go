package envinfo

func GetIDEs() []*Item {
	return getItems([]func() (*Item, error){
		GetAndroidStudioVersions,
		// TODO: Sublime Text
		// TODO: Xcode
		GetAtomVersion,
		GetEmacsVersion,
		GetIntelliJVersion,
		GetNanoVersion,
		GetNvimVersion,
		GetPhpStormVersion,
		GetVIMVersion,
		GetVSCodeVersion,
		GetWebStormVersion,
	})
}

func GetAtomVersion() (*Item, error) {
	return GetItem("atom", "Atom")
}

func GetEmacsVersion() (*Item, error) {
	return GetItem("emacs", "Emacs")
}

func GetVIMVersion() (*Item, error) {
	return GetItem("vim", "Vim")
}

func GetVSCodeVersion() (*Item, error) {
	return GetItem("code", "VSCode")
}

func GetNanoVersion() (*Item, error) {
	return GetItem("nano", "Nano")
}

func GetNvimVersion() (*Item, error) {
	return GetItem("nvim", "Nvim")
}
