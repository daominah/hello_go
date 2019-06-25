package main

import "testing"

func Test1(t *testing.T) {
	m := map[int][]int{}
	m[1] = []int{}
	//a := []int{}
	for i := 0; i < 10000; i++ {
		//m[1] = append(m[1], i)
		//a = append(a, i)
		//m[1] = a
	}
}
