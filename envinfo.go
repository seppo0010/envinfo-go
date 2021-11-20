package envinfo

import "sync"

type EnvInfoBuilder struct {
	languages bool
	binaries  bool
	system    bool
}

type EnvInfo struct {
	Languages []*Item `json:"Languages,omitempty"`
	Binaries  []*Item `json:"Binaries,omitempty"`
	System    *System `json:"System,omitempty"`
}

type Item struct {
	Name    string `json:"-"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

func NewEnvInfoBuilder() *EnvInfoBuilder {
	return &EnvInfoBuilder{}
}

func (b *EnvInfoBuilder) System() {
	b.system = true
}

func (b *EnvInfoBuilder) Binaries() {
	b.binaries = true
}

func (b *EnvInfoBuilder) Languages() {
	b.languages = true
}

func (b *EnvInfoBuilder) Build() *EnvInfo {
	envinfo := &EnvInfo{}
	var wg sync.WaitGroup
	if b.languages {
		wg.Add(1)
		func() {
			defer wg.Done()
			envinfo.Languages = GetLanguages()
		}()
	}
	if b.binaries {
		wg.Add(1)
		func() {
			defer wg.Done()
			envinfo.Binaries = GetBinaries()
		}()
	}
	if b.system {
		wg.Add(1)
		func() {
			defer wg.Done()
			envinfo.System = GetSystem()
		}()
	}
	wg.Wait()
	return envinfo
}
