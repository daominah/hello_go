package main

import (
	"context"
	"fmt"
)

func main() {
	a := 0
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		for i := 0; i < 1000; i++ {
			a += 1
		}
	}()
	for i := 0; i < 1000; i++ {
		a += 1
	}
	<-ctx.Done()
	fmt.Println(a)
}
