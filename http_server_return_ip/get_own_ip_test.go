package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	ip, _ := GetOutboundIP()
	hostName, _ := os.Hostname()
	t.Log(fmt.Sprintf("%v_%v", ip, hostName))
}
