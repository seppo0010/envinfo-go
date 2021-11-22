package envinfo

func GetVirtualization() []*Item {
	return getItems([]func() *Item{
		GetDockerVersion,
		GetParallelsVersion,
		GetVirtualBoxVersion,
		GetVMWareFusionVersion,
	})
}

func GetDockerVersion() *Item {
	return NewGetItemBuilder("docker", "Docker").Flag("-v").Get()
}

func GetParallelsVersion() *Item {
	return NewGetItemBuilder("prlctl", "Parallels").Flag("-v").Get()
}

func GetVirtualBoxVersion() *Item {
	return NewGetItemBuilder("vboxmanage", "VirtualBox").Flag("-v").Get()
}
