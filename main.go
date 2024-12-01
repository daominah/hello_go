// this file can be ignored so not push changes to git with command:
// git update-index --skip-worktree main.go
// git update-index --no-skip-worktree main.go
package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	a := 1
	b := "hehe"
	log.Printf("Hello Go")
	log.Printf("a = %v, b = %v", a, b)
	log.Printf("got there")
	a = 123
	log.Printf("a = %v, b = %v", a, b)
	log.Printf("main returned")
}
