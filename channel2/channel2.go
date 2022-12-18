package main

import (
	"fmt"
	"time"
)

// Sum returns n by calculating 1+1+1+..
func Sum(n int64) int64 {
	ret := int64(0)
	for i := int64(0); i < n; i++ {
		ret += 1
	}
	return ret
}

func main() {
	n := int64(30000000000) // 30e9
	sum := int64(0)
	beginTime := time.Now()

	nWorkers := 16
	sumChan := make(chan int64, nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() { sumChan <- Sum(n / int64(nWorkers)) }()
	}
	for i := 0; i < nWorkers; i++ {
		sum += <-sumChan
	}

	fmt.Println("dur:", time.Since(beginTime))
	fmt.Println("sum:", sum)

	// Results on Intel Core i5-8265U (4 physical cores)
	// (nWorkers,dur): (1, 8s), (2, 4s), (4, 2s), (8, 2s), (16, 2s)
}
