package main

import "testing"

func TestFib(t *testing.T) {
	if g, w := Fib(10), 55; g != w {
		t.Errorf("error Fib got %v, want %v", g, w)
	}
	if g, w := Fib(25), 75025; g != w {
		t.Errorf("error Fib got %v, want %v", g, w)
	}
}

func TestFiber(t *testing.T) {
	f := NewFiber()
	if g, w := f.Fib(10), 55; g != w {
		t.Errorf("error Fib got %v, want %v", g, w)
	}
	if g, w := f.Fib(25), 75025; g != w {
		t.Errorf("error Fib got %v, want %v", g, w)
	}
}

// 380000 ns/op (CPU i7-1260P)
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(25)
	}
}

// 1795 ns/op (CPU i7-1260P)
func BenchmarkFiber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := NewFiber()
		f.Fib(25)
	}
}
