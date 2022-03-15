package main

import (
	"testing"
)

func TestGetMemoryUsage(t *testing.T) {
	println(GetDiskUsage())
	println(GetMemoryUsage())
	println(GetCPUAverageUsage())
	println(GetCPUModel())
}
