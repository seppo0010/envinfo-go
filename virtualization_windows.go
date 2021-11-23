//go:build windows
// +build windows

package envinfo

func GetVMWareFusionVersion() *Item {
	return &Item{
		Name:    "VMWare Fusion",
		Path:    "",
		Version: "",
	}
}
