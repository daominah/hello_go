package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	//time.Sleep(60 * time.Millisecond)
	select {
	case <-time.After(20 * time.Millisecond):
		fmt.Println("overslept")
		fmt.Println("ctx.Err() 1:", ctx.Err()) // nil
	case <-ctx.Done():
		fmt.Println("ctx.Err() 2:", ctx.Err()) // context deadline exceeded
	}
}
