package main

import (
	"fmt"
	"testing"
)

func TestReadNetIO(t *testing.T) {
	nReceivedBytes, nTransmitedBytes, err := ReadNetworkInOut()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(nReceivedBytes, nTransmitedBytes, err)
}
