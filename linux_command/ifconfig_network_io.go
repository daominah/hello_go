package main

import (
	"os/exec"
	"strconv"
	"strings"
)

// return number of received bytes and number of transmited bytes.
// tested in linux mint 17
func ReadNetworkInOut() (nReceivedBytes int64, nTransmitedBytes int64, err error) {
	cmd := exec.Command("ifconfig")
	stdoutB, err := cmd.Output()
	if err != nil {
		return
	}
	stdout := string(stdoutB)
	netIfs := strings.Split(stdout, "\n\n")
	var eth0 string
	for _, netIf := range netIfs {
		if !strings.Contains(netIf, "docker") &&
			!strings.Contains(netIf, "lo") &&
			!strings.Contains(netIf, "wlan") {
			eth0 = netIf
			break
		}
	}
	for _, substr := range []string{"RX bytes:", "TX bytes:"} {
		iBegin := strings.Index(eth0, substr) + len(substr)
		iEnd := iBegin + strings.Index(eth0[iBegin:], " ")
		nBytesStr := eth0[iBegin:iEnd]
		var nBytes int64
		nBytes, err = strconv.ParseInt(nBytesStr, 10, 64)
		if err != nil {
			return
		}
		switch substr {
		case "RX bytes:":
			nReceivedBytes = nBytes
		case "TX bytes:":
			nTransmitedBytes = nBytes
		}
	}
	return
}
