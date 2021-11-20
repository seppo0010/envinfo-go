package envinfo

import (
	"fmt"
	"runtime"

	"code.cloudfoundry.org/bytefmt"
	"github.com/pbnjay/memory"
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

func getMemory() string {
	return fmt.Sprintf("%s / %s", bytefmt.ByteSize(memory.FreeMemory()), bytefmt.ByteSize(memory.TotalMemory()))
}

func GetSystem() *System {
	return &System{
		OS:     getOS(),
		CPU:    getCPU(),
		Memory: getMemory(),
	}
}
