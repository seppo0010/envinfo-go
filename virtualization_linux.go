//go:build linux
// +build linux

package envinfo

func GetVMWareFusionVersion() (*Item, error) {
	return &Item{
		Name:    "VMWare Fusion",
		Path:    "",
		Version: "",
	}, nil
}
