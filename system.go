package envinfo

import (
	"fmt"
	"os"
	"runtime"

	"code.cloudfoundry.org/bytefmt"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
	log "github.com/sirupsen/logrus"
)

type System struct {
	OS        string
	CPU       string
	Memory    string
	Shell     string
	Container string
}

func getCPU() string {
	cpus, err := cpu.Info()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("failed to get cpu info")
		return fmt.Sprintf("(%d)", runtime.NumCPU())
	}
	return fmt.Sprintf("(%d) %s %s", runtime.NumCPU(), runtime.GOARCH, cpus[0].ModelName)
}

func getMemory() string {
	return fmt.Sprintf("%s / %s", bytefmt.ByteSize(memory.FreeMemory()), bytefmt.ByteSize(memory.TotalMemory()))
}

func getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		log.WithFields(log.Fields{}).Warn("unknown shell")
		return "N/A"
	}
	item, err := GetItem(shell, shell, "--version")
	if err != nil {
		return "Unknown"
	}
	return fmt.Sprintf("%s - %s", item.Version, item.Path)
}

func getContainer() string {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "Yes"
	}
	if _, err := os.Stat("/proc/self/cgroup"); err == nil {
		return "Yes"
	}
	return ""
}

func GetSystem() *System {
	return &System{
		OS:        getOS(),
		CPU:       getCPU(),
		Memory:    getMemory(),
		Shell:     getShell(),
		Container: getContainer(),
	}
}
