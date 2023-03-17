package main

// Fibonacci sequence 0, 1, 1, 2, 3, 5, 8, ..
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

type Fiber struct{ cache map[int]int }

func NewFiber() *Fiber {
	f := &Fiber{cache: make(map[int]int)}
	f.cache[0] = 0
	f.cache[1] = 1
	return f
}

func (f *Fiber) Fib(n int) int {
	calculated, found := f.cache[n]
	if found {
		return calculated
	}
	f.cache[n] = f.Fib(n-1) + f.Fib(n-2)
	return f.cache[n]
}
