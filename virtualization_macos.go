//go:build darwin
// +build darwin

package envinfo

func GetVMWareFusionVersion() *Item {
	return getApplication("VMWare Fusion", "com.vmware.fusion")
}
