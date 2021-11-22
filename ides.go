package envinfo

func GetIDEs() []*Item {
	return getItems([]func() *Item{
		GetAndroidStudioVersions,
		GetAtomVersion,
		GetEmacsVersion,
		GetIntelliJVersion,
		GetNanoVersion,
		GetNvimVersion,
		GetPhpStormVersion,
		GetSublimeVersion,
		GetVIMVersion,
		GetVSCodeVersion,
		GetWebStormVersion,
		GetXcodeVersion,
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
