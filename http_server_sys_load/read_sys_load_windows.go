package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// example result:
// "Disk Usage: 68.4/112.7 GiB (65%)"
func GetDiskUsage() string {
	cmd := exec.Command(`cmd`, `/c`,
		`wmic LogicalDisk get FreeSpace,Size`)
	stdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetDiskUsage: %v", err)
	}
	lines := strings.Split(string(stdoutB), "\r\n")
	if len(lines) < 2 {
		return fmt.Sprintf("error GetDiskUsage: unexpected stdout: %s", stdoutB)
	}
	words := strings.Fields(lines[1])
	if len(words) != 2 {
		return fmt.Sprintf("error GetDiskUsage: unexpected stdout: %s", stdoutB)
	}
	freeDiskBytes, _ := strconv.ParseInt(words[0], 10, 64)
	totalDiskBytes, _ := strconv.ParseInt(words[1], 10, 64)
	totalDiskGiB := float64(totalDiskBytes) / 1024 / 1024 / 1024
	usedDiskGiB := float64(totalDiskBytes-freeDiskBytes) / 1024 / 1024 / 1024
	return fmt.Sprintf("Disk Usage: %.1f/%.1f GiB (%.0f%%)",
		usedDiskGiB, totalDiskGiB, 100*usedDiskGiB/totalDiskGiB)
}

// example result:
// "Memory Usage: 8248/15990 MiB (51.58%)"
func GetMemoryUsage() string {
	cmd := exec.Command(`cmd`, `/c`,
		`wmic ComputerSystem get TotalPhysicalMemory`)
	//fmt.Printf("%#v", cmd.Args)
	stdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetMemoryUsage: %v", err)
	}
	lines := strings.Split(string(stdoutB), "\r\n")
	if len(lines) < 2 {
		return fmt.Sprintf("error GetMemoryUsage: unexpected stdout: %s", stdoutB)
	}
	totalMemoryBytes, _ := strconv.ParseInt(strings.TrimSpace(lines[1]), 10, 64)
	totalMemoryMiB := totalMemoryBytes / 1024 / 1024
	cmd = exec.Command(`cmd`, `/c`,
		`wmic OS get FreePhysicalMemory`)
	stdoutB, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetMemoryUsage: %v", err)
	}
	lines = strings.Split(string(stdoutB), "\r\n")
	freeMemoryKiBs, _ := strconv.ParseInt(strings.TrimSpace(lines[1]), 10, 64)
	freeMemoryMiB := freeMemoryKiBs / 1024
	usedMemoryMiB := totalMemoryMiB - freeMemoryMiB
	return fmt.Sprintf("Memory Usage: %v/%v MiB (%.0f%%)",
		usedMemoryMiB, totalMemoryMiB, 100*float64(usedMemoryMiB)/float64(totalMemoryMiB))
}

func GetCPUAverageUsage() string {
	cmd := exec.Command(`cmd`, `/c`, `wmic cpu get LoadPercentage`)
	wmicStdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetCPUModel: %v", err)
	}
	wmicStdout := strings.TrimSpace(string(wmicStdoutB))
	cpuLines := strings.Split(wmicStdout, "\r\n")
	nCPUSockets := 0
	sumLoadPercent := 0
	for i, line := range cpuLines {
		if i == 0 {
			continue
		}
		nCPUSockets += 1
		loadPercent, _ := strconv.Atoi(line)
		sumLoadPercent += loadPercent
	}
	if nCPUSockets == 0 {
		return fmt.Sprintf("error GetCPUAverageUsage: nCPUSockets is 0")
	}
	return fmt.Sprintf("CPU Usage: %v%%", sumLoadPercent/nCPUSockets)
}

// returns example: "4 cores (Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz, 1 socket)"
func GetCPUModel() string {
	cmd := exec.Command(`cmd`, `/c`, `wmic cpu get Name,NumberOfCores`)
	wmicStdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetCPUModel: %v", err)
	}
	wmicStdout := strings.TrimSpace(string(wmicStdoutB))
	//println(wmicStdout)
	var cpuModel string
	cpuLines := strings.Split(wmicStdout, "\r\n")
	nCPUSockets := len(cpuLines) - 1 // the first line is column label
	if nCPUSockets >= 1 {
		cpu0 := strings.TrimSpace(cpuLines[1])
		splitIdx := strings.LastIndex(cpu0, " ")
		if splitIdx != -1 {
			cpuName := strings.TrimSpace(cpu0[:splitIdx])
			nCoresPerSocket, _ := strconv.Atoi(cpu0[splitIdx+1:]) // per CPU socket
			cpuModel = fmt.Sprintf("%v cores (%v, %v socket)",
				nCoresPerSocket*nCPUSockets, cpuName, nCPUSockets)
		}
		return cpuModel
	}
	return "error GetCPUModel: unexpected wmic output: " + wmicStdout
}
