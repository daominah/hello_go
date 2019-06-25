package main

import (
	"fmt"
	"sort"
)

type Player struct {
	name string
	age  int
}

type byLexicography []Player

func (s byLexicography) Len() int {
	return len(s)
}
func (s byLexicography) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLexicography) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func main() {
	ps := []Player{
		{"tung", 20}, {"lan", 19}, {"van", 15},
		{"hoa", 21},
	}
	temp := byLexicography(ps)
	sort.Sort(temp)
	fmt.Println(ps)
}
