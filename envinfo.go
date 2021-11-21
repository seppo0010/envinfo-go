package envinfo

import "sync"

type EnvInfoBuilder struct {
	languages      bool
	binaries       bool
	managers       bool
	browsers       bool
	utilities      bool
	virtualization bool
	servers        bool
	system         bool
}

type EnvInfo struct {
	Languages      []*Item `json:"Languages,omitempty"`
	Binaries       []*Item `json:"Binaries,omitempty"`
	Managers       []*Item `json:"Managers,omitempty"`
	Browsers       []*Item `json:"Browsers,omitempty"`
	Utilities      []*Item `json:"Utilities,omitempty"`
	Virtualization []*Item `json:"Virtualization,omitempty"`
	Servers        []*Item `json:"Servers,omitempty"`
	System         *System `json:"System,omitempty"`
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

func (b *EnvInfoBuilder) Managers() {
	b.managers = true
}

func (b *EnvInfoBuilder) Browsers() {
	b.browsers = true
}

func (b *EnvInfoBuilder) Utilities() {
	b.utilities = true
}

func (b *EnvInfoBuilder) Virtualization() {
	b.virtualization = true
}

func (b *EnvInfoBuilder) Servers() {
	b.servers = true
}

func (b *EnvInfoBuilder) Languages() {
	b.languages = true
}

func (b *EnvInfoBuilder) Build() *EnvInfo {
	envinfo := &EnvInfo{}
	var wg sync.WaitGroup
	if b.languages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Languages = GetLanguages()
		}()
	}
	if b.browsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Browsers = GetBrowsers()
		}()
	}
	if b.utilities {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Utilities = GetUtilities()
		}()
	}
	if b.virtualization {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Virtualization = GetVirtualization()
		}()
	}
	if b.servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Servers = GetServers()
		}()
	}
	if b.managers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Managers = GetManagers()
		}()
	}
	if b.binaries {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.Binaries = GetBinaries()
		}()
	}
	if b.system {
		wg.Add(1)
		go func() {
			defer wg.Done()
			envinfo.System = GetSystem()
		}()
	}
	wg.Wait()
	return envinfo
}
