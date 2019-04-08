package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	fmt.Print("")
}

type EC struct {
	x int
	y string
}

func f1() *EC {
	a := []int{0, 1, 2}
	_ = a[5]
	return &EC{1, "2"}
}

func f11() (r float64) {
	fmt.Println("hihi 0")
	defer func() {
		if r := recover(); r != nil {
			bytes := debug.Stack()
			fmt.Println("Recovered in f11", r)
			fmt.Println("Recovered in f11", r, string(bytes))
			r = 5
		}
	}()
	panic("fuck")
	return r
}

func f2() {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			bytes := debug.Stack()
	//			fmt.Println("Recovered in f2", r, string(bytes))
	//		}
	//	}()
	// ec := f1()
	x := f11()
	fmt.Println("xong", x)
	//fmt.Println("ec", ec)
	//fmt.Println("ec.x", ec.x)
}

func main() {
	f11()
}
