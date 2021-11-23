package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func parseColonSeparated(s string) map[string]string {
	ret := make(map[string]string)
	for _, line := range strings.Split(s, "\n") {
		beginIdx := strings.Index(line, ":")
		if beginIdx == -1 || beginIdx+1 > len(line) {
			continue
		}
		key := strings.TrimSpace(line[:beginIdx])
		val := strings.TrimSpace(line[beginIdx+1:])
		//println("__________", key, "__________", val)
		ret[key] = val
	}
	return ret
}

func main() {
	// CPU average load, require `apt install -y sysstat`
	var preCalcCPUUsage float64
	go func() {
		for {
			cmd := exec.Command("mpstat", "1", "1")
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatalf("error: %v, plz run `apt install -y sysstat`", err)
			}
			preCalcCPUUsage = parseMpstatToCPUUsage(string(stdout))
			time.Sleep(1 * time.Second)
		}
	}()

	// CPU model
	cmd := exec.Command("/bin/bash", "-c", `lscpu`)
	lscpuStdout, err0 := cmd.Output()
	if err0 != nil {
		log.Fatalf("error get cpu model: %v", err0)
	}
	cpuInfos := parseColonSeparated(string(lscpuStdout))
	cpuModel := cpuInfos["Model name"]
	nSockets, _ := strconv.Atoi(cpuInfos["Socket(s)"])
	nCoresPerSocket, _ := strconv.Atoi(cpuInfos["Core(s) per socket"])
	nPhysicalCores := nSockets * nCoresPerSocket
	cpuInfosStr := fmt.Sprintf("%v physical cores (%v)", nPhysicalCores, cpuModel)

	// listen HTTP, default port 21864
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		cmd := exec.Command("/bin/bash", "-c",
			`free -m | awk 'NR==2{printf "Memory Usage: %s/%sMB (%.2f%%)\n", $3,$2,$3*100/$2 }'
df -h | awk '$NF=="/"{printf "Disk Usage: %d/%dGB (%s)\n", $3,$2,$5}'`)
		stdout, err := cmd.Output()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error exec: %v", err)))
			return
		}
		resp := fmt.Sprintf("%sCPU Usage: %.2f%% of %v\n",
			stdout, preCalcCPUUsage, cpuInfosStr)
		w.Write([]byte(resp))
	})

	listen := os.Getenv("LISTENING_PORT")
	if listen == "" {
		listen = ":21864"
	}
	if !strings.Contains(listen, ":") {
		listen = ":" + listen
	}

	server := &http.Server{Addr: listen, Handler: handler}
	log.Printf("listening on http://127.0.0.1%v/\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// 05:48:31 PM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
// Average:     all    2,79    0,00    1,49    0,02    0,00    0,02    0,00    0,00    0,00   95,67
func parseMpstatToCPUUsage(s string) float64 {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return 0
	}
	lastLine := strings.TrimSpace(lines[len(lines)-1])
	fields := strings.Split(lastLine, " ")
	if len(fields) < 1 {
		return 0
	}
	idleField := strings.ReplaceAll(fields[len(fields)-1], ",", ".")
	idle, err := strconv.ParseFloat(idleField, 64)
	if err != nil {
		return 0
	}
	return 100 - idle
}
