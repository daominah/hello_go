package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetDiskUsage() string { return "TODO" }

func GetMemoryUsage() string { return "TODO" }

func GetCPUAverageUsage() string { return "TODO" }

// returns example: "4 cores (Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz, 1 socket)"
func GetCPUModel() string {
	cmd := exec.Command(`cmd`, `/C`, `wmic cpu get Name,NumberOfCores`)
	cpuModelB, err0 := cmd.Output()
	if err0 != nil {
		log.Fatalf("error get cpu model: %v", err0)
	}
	cpuModels := strings.TrimSpace(string(cpuModelB))
	var cpuInfosStr string
	cpuLines := strings.Split(cpuModels, "\r\n")
	// sure, the first line is column label
	nCPUSockets := len(cpuLines) - 1
	if nCPUSockets >= 1 {
		cpu0 := strings.TrimSpace(cpuLines[1])
		splitIdx := strings.LastIndex(cpu0, " ")
		if splitIdx != -1 {
			cpuModel := strings.TrimSpace(cpu0[:splitIdx])
			nCoresPerSocket, _ := strconv.Atoi(cpu0[splitIdx+1:]) // per CPU socket
			cpuInfosStr = fmt.Sprintf("%v physical cores (%v, %v socket)",
				nCoresPerSocket*nCPUSockets, cpuModel, nCPUSockets)
		}
		log.Printf("cpuInfosStr: %#v\n", cpuInfosStr)
		return cpuInfosStr
	}
	return "" // unreachable
}
