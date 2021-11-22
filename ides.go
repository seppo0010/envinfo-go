package envinfo

func GetIDEs() []*Item {
	return getItems([]func() *Item{
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

func GetAtomVersion() *Item {
	return GetItem("atom", "Atom")
}

func GetEmacsVersion() *Item {
	return GetItem("emacs", "Emacs")
}

func GetVIMVersion() *Item {
	return GetItem("vim", "Vim")
}

func GetVSCodeVersion() *Item {
	return GetItem("code", "VSCode")
}

func GetNanoVersion() *Item {
	return GetItem("nano", "Nano")
}

func GetNvimVersion() *Item {
	return GetItem("nvim", "Nvim")
}
