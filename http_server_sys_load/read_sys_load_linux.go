package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// based on command "df", example result:
// "Disk Usage: 68.4/112.7 GiB (65%)"
func GetDiskUsage() string {
	cmd := exec.Command("/bin/bash", "-c",
		`df -BM | awk '$NF=="/"{printf "Disk Usage: %.1f/%.1f GiB (%s)\n", $3/1024,$2/1024,$5}'`)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetDiskUsage: %v", err)
	}
	return strings.TrimSpace(string(stdout))
}

// based on command "free -m", example result:
// "Memory Usage: 8248/15990 MiB (51.58%)"
func GetMemoryUsage() string {
	cmd := exec.Command("/bin/bash", "-c",
		`free -m | awk 'NR==2{printf "Memory Usage: %s/%s MiB (%.2f%%)\n", $3,$2,$3*100/$2 }'`)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetMemoryUsage: %v", err)
	}
	return strings.TrimSpace(string(stdout))
}

// based on command "mpstat" (apt package "sysstat"),
// returns CPU average usage in the last 1 second, example result:
// "CPU Usage: 12.06%"
func GetCPUAverageUsage() string {
	cmd := exec.Command("mpstat", "1", "1")
	stdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetCPUAverageUsage: %v", err)
	}
	stdout := string(stdoutB)
	//println(stdout)
	// stdout       CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
	// Average:     all    8,96    0,00    0,87    0,00    0,00    2,24    0,00    0,00    0,00   87,94

	var usageStr string
	lines := strings.Split(strings.TrimSpace(stdout), "\n")
	if len(lines) > 0 {
		lastLine := strings.TrimSpace(lines[len(lines)-1])
		fields := strings.Split(lastLine, " ")
		if len(fields) > 0 {
			idleField := strings.ReplaceAll(fields[len(fields)-1], ",", ".")
			idle, _ := strconv.ParseFloat(idleField, 64)
			usageStr = strconv.FormatFloat(100-idle, 'f', 2, 64)
		}
	}
	return `CPU Usage: ` + usageStr + `%`
}

// based on command "lscpu", example result:
// "16 cores (Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz, 2 socket)"
func GetCPUModel() string {
	cmd := exec.Command("/bin/bash", "-c", `lscpu`)
	lscpuStdoutB, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("error GetCPUModel: %v", err)
	}
	lscpuStdout := string(lscpuStdoutB)
	//println(lscpuStdout)
	// Model name:          Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
	// Socket(s):           1
	// Core(s) per socket:  4

	lscpuKeyVals := make(map[string]string)
	for _, line := range strings.Split(lscpuStdout, "\n") {
		colonIdx := strings.Index(line, ":")
		if colonIdx == -1 {
			continue
		}
		key := strings.TrimSpace(line[:colonIdx])
		val := strings.TrimSpace(line[colonIdx+1:])
		lscpuKeyVals[key] = val
	}
	nSockets, _ := strconv.Atoi(lscpuKeyVals["Socket(s)"])
	nCoresPerSocket, _ := strconv.Atoi(lscpuKeyVals["Core(s) per socket"])
	return fmt.Sprintf("%v cores (%v, %v socket)",
		nCoresPerSocket*nSockets, lscpuKeyVals["Model name"], nSockets)
}
