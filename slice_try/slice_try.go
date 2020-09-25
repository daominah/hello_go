package main

import (
	"github.com/mywrap/log"
)

type A struct {
	slice []int
}

func main() {
	a := A{slice: []int{0, 1, 2, 3, 4, 5}}
	b := a

	a.slice[1] = -1
	log.Printf("a: %#v", a)
	log.Printf("b: %#v, b[1] == -1: %v", b, b.slice[1] == -1)

	a.slice = append(a.slice, 6)
	a.slice[2] = -2
	log.Printf("a: %#v", a)
	log.Printf("b: %#v, b[2] == -2, %v", b, b.slice[2] == -2)

	log.Printf("after append, the underlying array changed")
}
