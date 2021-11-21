//go:build darwin
// +build darwin

package envinfo

func GetVMWareFusionVersion() (*Item, error) {
	return getApplication("VMWare Fusion", "com.vmware.fusion")
}
