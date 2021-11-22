package envinfo

func GetIDEs() []*Item {
	return getItems([]func() (*Item, error){
		GetAndroidStudioVersions,
		// TODO: IntelliJ
		// TODO: Nano
		// TODO: Nvim
		// TODO: PhpStorm
		// TODO: Sublime Text
		// TODO: WebStorm
		// TODO: Xcode
		GetAtomVersion,
		GetEmacsVersion,
		GetVIMVersion,
		GetVSCodeVersion,
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
