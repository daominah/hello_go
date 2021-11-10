package main

import (
	"fmt"
)

func main() {
	a := 0
	done := make(chan bool)
	go func() {
		for i := 0; i < 100000; i++ {
			a += 1
		}
		done <- true
	}()
	for i := 0; i < 100000; i++ {
		a += 1
	}
	<-done
	fmt.Println(a)
}
