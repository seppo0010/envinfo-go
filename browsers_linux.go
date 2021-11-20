//go:build linux
// +build linux

package envinfo

func GetBraveVersion() (*Item, error) {
	item, _ := GetItem("brave", "Brave Browser", "--version")
	if item != nil {
		return item, nil
	}
	return GetItem("brave-browser", "Brave Browser", "--version")
}
