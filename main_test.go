package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	double := func(i int) int { return 2 * i }
	if got, want := double(1), 2; got != want {
		t.Errorf("error TestExample: got %v, want %v", got, want)
	}
}
