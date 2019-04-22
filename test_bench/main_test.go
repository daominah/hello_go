package main

import "testing"

func TestFib(t *testing.T) {
	if Fib(10) != 55 {
		t.Error()
	}
	if Fib(25) != 75025 {
		t.Error()
	}
}

func TestFiber(t *testing.T) {
	f := NewFiber()
	if f.Fib(10) != 55 {
		t.Error()
	}
	if f.Fib(25) != 75025 {
		t.Error()
	}
}

// from fib_test.go
func BenchmarkFib(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(25)
	}
}

func BenchmarkFiber(b *testing.B) {
	f := NewFiber()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		f.Fib(25)
	}
}

