// This package implement map data structure, capacity is 10 000 000 key,
// (average time to insert , search , delete is O(1))
// using division method as hash function and resolve collision by chaining
package main

import (
	"container/list"
	"fmt"
)

const (
	// number of slots for hashing
	NSlot = 10000019
)

// A good hash func satisfies each key is equally likely to hash to any of the slots
// This func take a string as a represent of a radix-128 integer,
// return remainder when the string is divided by number of slots
func DivisionHash(s string, nSlot int) int {
	base := 128
	result := 0
	for _, r := range s {
		ri := int(r) // char to ascii
		result = (result*base + ri) % nSlot
	}
	return result
}

type Map struct {
	// array of linked lists
	Table []*list.List
	Len   int
}
type KeyAndValue struct {
	Key   string
	Value interface{}
}

func (m *Map) Init() {
	m.Table = make([]*list.List, NSlot)
}

// worst case time of hashing with chaining is O(n),
// when all n keys hash to the same list, creating a list of length n
// average case time O(1+n/NSlot).
// Reference "Introduction to Algorithms", chap 11 Hash Tables
func (m *Map) Insert(key string, value interface{}) {
	if m.Table[DivisionHash(key, NSlot)] == nil {
		m.Table[DivisionHash(key, NSlot)] = list.New()
	}
	l := m.Table[DivisionHash(key, NSlot)]
	for e := l.Front(); e != nil; e = e.Next() {
		e2 := e.Value.(*KeyAndValue)
		if e2.Key == key {
			e2.Value = value
			return
		}
	}
	m.Table[DivisionHash(key, NSlot)].PushFront(&KeyAndValue{Key: key, Value: value})
	m.Len += 1
}
func (m *Map) Delete(key string) {
	if m.Table[DivisionHash(key, NSlot)] == nil {
		m.Table[DivisionHash(key, NSlot)] = list.New()
	}
	l := m.Table[DivisionHash(key, NSlot)]
	for e := l.Front(); e != nil; e = e.Next() {
		e2 := e.Value.(*KeyAndValue)
		if e2.Key == key {
			l.Remove(e)
			m.Len -= 1
			return
		}
	}
}
func (m *Map) Search(key string) interface{} {
	if m.Table[DivisionHash(key, NSlot)] == nil {
		m.Table[DivisionHash(key, NSlot)] = list.New()
	}
	l := m.Table[DivisionHash(key, NSlot)]
	for e := l.Front(); e != nil; e = e.Next() {
		e2 := e.Value.(*KeyAndValue)
		if e2.Key == key {
			return e2.Value
		}
	}
	return nil
}

func main() {
	fmt.Println("b")
}
