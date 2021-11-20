package envinfo

type EnvInfo struct {
	Languages []*Item
	Binaries  []*Item
	System    *System
}

func NewEnvInfo() *EnvInfo {
	envinfo := &EnvInfo{}
	envinfo.Languages = GetLanguages()
	envinfo.Binaries = GetBinaries()
	envinfo.System = GetSystem()
	return envinfo
}
