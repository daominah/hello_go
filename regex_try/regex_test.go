package main

import (
	"log"
	"regexp"
	"testing"
)

func TestRegexpMatch(t *testing.T) {
	pattern0, err := regexp.Compile(`^[a-zA-Z0-9._, -]*$`)
	if err != nil {
		log.Fatalf("error regexp Compile: %v", err)
	}

	for _, c := range []struct {
		s string
		m bool
	}{
		{`key0_val0_1.23, Reg-Exp is not for human`, true},
		{`tung@gmail.com`, false},
		{`Đào Lán`, false},
	} {
		if r := pattern0.MatchString(c.s); r != c.m {
			t.Errorf("error MatchString: real: %v, expected: %v", r, c.m)
		}
	}
}
