package main

import "testing"

func Test1(t *testing.T) {
	if "1" != "1" {
		t.Error()
	}
	if "1" != "2" {
		t.Error()
	}
}