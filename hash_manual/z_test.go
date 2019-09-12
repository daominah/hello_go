package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDivisionHash1(t *testing.T) {
	type Case struct {
		in  string
		out int
	}
	for _, c := range []Case{
		Case{"j", 106},
		Case{"jj", 13674},
		Case{"4:[;", 13674},
		Case{"jZe", 1748325},
		Case{"abcd", 5041768},
		Case{"asdlfjaoweAAaojopa3028452ZZ", 7751450},
	} {
		in, out := c.in, c.out
		if DivisionHash(in, NSlot) != out {
			t.Error(in, out, DivisionHash(in, NSlot))
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890,./;'[]-=")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// test each key is equally likely to hash to any of the slots
func TestDivisionHash2(t *testing.T) {
	_ = fmt.Println
	nTest := 10000
	nHashSlot := 10
	counter := map[int]int{}
	for i := 0; i < nTest; i++ {
		lenS := 1 + rand.Intn(30)
		s := RandString(lenS)
		// fmt.Println(s)
		hashed := DivisionHash(s, nHashSlot)
		counter[hashed] += 1
	}
	//	fmt.Println(counter)
	for _, v := range counter {
		if v < nTest/nHashSlot*8/10 || v > nTest/nHashSlot*12/10 {
			t.Error()
			return
		}
	}
}

func TestMap(t *testing.T) {
	m := Map{}
	m.Init()
	m.Insert("hohohaha", 4)
	m.Insert("hohohaha", 7)
	m.Insert("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 9)
	m.Insert("hihi", 3)
	m.Insert("jj", 11)
	m.Insert("4:[;", 5555)
	if m.Len != 5 {
		t.Error()
	}
	if m.Table[DivisionHash("jj", NSlot)].Len() != 2 {
		t.Error()
	}
	if m.Search("jj") != 11 {
		t.Error()
	}
	if m.Search("4:[;") != 5555 {
		t.Error()
	}
	if m.Search("hohohaha") != 7 {
		t.Error()
	}
	if m.Search("oeoe") != nil {
		t.Error()
	}
	m.Delete("jj")
	if m.Len != 4 {
		t.Error()
	}
	if m.Table[DivisionHash("jj", NSlot)].Len() != 1 {
		t.Error()
	}
	if m.Search("jj") != nil {
		t.Error()
	}
}

func TestMap2(t *testing.T) {
	m := Map{}
	m.Init()
	for i := 0; i < 10000; i++ {
		m.Insert(fmt.Sprintf("%v", i), i)
	}
}
