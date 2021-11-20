package envinfo

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
)

type System struct {
	OS     string
	CPU    string
	Memory string
	Shell  string
}

func getCPU() string {
	cpus, err := cpu.Info()
	if err != nil {
		return fmt.Sprintf("(%d)", runtime.NumCPU())
	}
	return fmt.Sprintf("(%d) %s %s", runtime.NumCPU(), runtime.GOARCH, cpus[0].ModelName)
}
func GetSystem() *System {
	return &System{
		OS:  getOS(),
		CPU: getCPU(),
	}
}
