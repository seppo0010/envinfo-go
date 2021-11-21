package envinfo

func GetIDEs() []*Item {
	return getItems([]func() (*Item, error){
		// TODO: Android Studio
		// TODO: IntelliJ
		// TODO: Nano
		// TODO: PhpStorm
		// TODO: Sublime Text
		// TODO: VSCode
		// TODO: WebStorm
		// TODO: Xcode
		GetAtomVersion,
		GetEmacsVersion,
		GetVIMVersion,
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
