package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		defer wg.Add(-1)
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("gen about to return")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	wg.Add(1)
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel()
	wg.Wait()
	fmt.Println("end")
}
