package main

import (
	"fmt"
	"testing"
)

// func ReadNetworkInOut depends on bash `ifconfig` output which can change
func _TestReadNetIO(t *testing.T) {
	nReceivedBytes, nTransmitedBytes, err := ReadNetworkInOut()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(nReceivedBytes, nTransmitedBytes, err)
}
