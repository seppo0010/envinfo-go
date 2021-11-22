package envinfo

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type EnvInfoBuilder struct {
	languages      bool
	binaries       bool
	managers       bool
	browsers       bool
	utilities      bool
	virtualization bool
	servers        bool
	sdks           bool
	ides           bool
	system         bool
}

type SDKs struct {
	Android *SDKManagerPackage
	IOS     []string
}
type EnvInfo struct {
	Languages      []*Item `json:"Languages,omitempty"`
	Binaries       []*Item `json:"Binaries,omitempty"`
	Managers       []*Item `json:"Managers,omitempty"`
	Browsers       []*Item `json:"Browsers,omitempty"`
	Utilities      []*Item `json:"Utilities,omitempty"`
	IDEs           []*Item `json:"IDEs,omitempty"`
	Virtualization []*Item `json:"Virtualization,omitempty"`
	Servers        []*Item `json:"Servers,omitempty"`
	SDKs           *SDKs   `json:"SDKs,omitempty"`
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

func (b *EnvInfoBuilder) SDKs() {
	b.sdks = true
}

func (b *EnvInfoBuilder) IDEs() {
	b.ides = true
}

func (b *EnvInfoBuilder) Languages() {
	b.languages = true
}

func (b *EnvInfoBuilder) Build() *EnvInfo {
	start := time.Now()
	envinfo := &EnvInfo{SDKs: &SDKs{}}
	var wg sync.WaitGroup
	if b.languages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Languages = GetLanguages()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("languages time")
		}()
	}
	if b.browsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Browsers = GetBrowsers()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("browsers time")
		}()
	}
	if b.utilities {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Utilities = GetUtilities()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("utilities time")
		}()
	}
	if b.virtualization {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Virtualization = GetVirtualization()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("virtualization time")
		}()
	}
	if b.servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Servers = GetServers()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("servers time")
		}()
	}
	if b.sdks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			version, _ := GetIOSVersions()
			envinfo.SDKs.IOS = version
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("ios sdk time")
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			version, _ := GetAndroidVersions()
			envinfo.SDKs.Android = version
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("android sdk time")
		}()
	}
	if b.ides {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.IDEs = GetIDEs()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("ides time")
		}()
	}
	if b.managers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Managers = GetManagers()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("managers time")
		}()
	}
	if b.binaries {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.Binaries = GetBinaries()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("binaries time")
		}()
	}
	if b.system {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			envinfo.System = GetSystem()
			log.WithFields(log.Fields{
				"duration": time.Now().Sub(start),
			}).Debug("system time")
		}()
	}
	log.WithFields(log.Fields{
		"duration": time.Now().Sub(start),
	}).Debug("will wait")
	wg.Wait()
	log.WithFields(log.Fields{
		"duration": time.Now().Sub(start),
	}).Debug("did wait")
	return envinfo
}
