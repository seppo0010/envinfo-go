package envinfo

func GetVirtualization() []*Item {
	return getItems([]func() (*Item, error){
		GetDockerVersion,
		GetParallelsVersion,
		GetVirtualBoxVersion,
		GetVMWareFusionVersion,
	})
}

func GetDockerVersion() (*Item, error) {
	return NewGetItemBuilder("docker", "Docker").Flag("-v").Get()
}

func GetParallelsVersion() (*Item, error) {
	return NewGetItemBuilder("prlctl", "Parallels").Flag("-v").Get()
}

func GetVirtualBoxVersion() (*Item, error) {
	return NewGetItemBuilder("vboxmanage", "VirtualBox").Flag("-v").Get()
}
