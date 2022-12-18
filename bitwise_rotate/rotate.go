package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func main() {
	tmp, _ := strconv.ParseUint("10110101", 2, 8)
	b := uint8(tmp)
	fmt.Printf("b: %08b\n", b)
	for iLeft := 0; iLeft < 8; iLeft++ {
		rotate := bits.RotateLeft8(b, iLeft+1)
		fmt.Printf("digit index %v: rotate %08b", iLeft, rotate)
		digit := rotate & 1
		fmt.Printf(", digit index: %v\n", digit)
	}
}
