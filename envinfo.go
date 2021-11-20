package envinfo

import "sync"

type EnvInfo struct {
	Languages []*Item
	Binaries  []*Item
	System    *System
}

func NewEnvInfo() *EnvInfo {
	envinfo := &EnvInfo{}
	var wg sync.WaitGroup
	wg.Add(1)
	func() {
		defer wg.Done()
		envinfo.Languages = GetLanguages()
	}()
	wg.Add(1)
	func() {
		defer wg.Done()
		envinfo.Binaries = GetBinaries()
	}()
	wg.Add(1)
	func() {
		defer wg.Done()
		envinfo.System = GetSystem()
	}()
	wg.Wait()
	return envinfo
}
