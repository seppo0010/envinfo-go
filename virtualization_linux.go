//go:build linux
// +build linux

package envinfo

func GetVMWareFusionVersion() *Item {
	return &Item{
		Name:    "VMWare Fusion",
		Path:    "",
		Version: "",
	}
}
