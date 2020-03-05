package main

import (
	"fmt"
	"time"
)

// Sum returns b + (b+1) + (b+2) + .. + (e-1)
func Sum(b int, e int) int {
	ret := 0
	for i := b; i < e; i++ {
		ret += 1
	}
	return ret
}

func main() {
	n := 30000000000 // 30e9
	sum := 0
	bt := time.Now()

	// Expectations of (nWorkers,dur): (1, 8s), (2, 4s), (4, 2s), ..
	// but why (8, 2s) too ???
	nWorkers := 4
	retChan := make(chan int)
	segs := make([]int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		segs[i] = (i + 1) * n / nWorkers
	}
	for i, e := range segs {
		b := 0
		if i > 0 {
			b = segs[i-1]
		}
		e := e
		go func() { retChan <- Sum(b, e) }()
	}
	for i := 0; i < nWorkers; i++ {
		sum += <-retChan
	}

	//
	dur := time.Now().Sub(bt)
	fmt.Println("dur:", dur)
	fmt.Println("sum:", sum)
}
