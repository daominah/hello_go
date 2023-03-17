package main

import (
	"testing"
)

func TestCompress(t *testing.T) {
	origin := []byte("hohohahaolala4567890")
	compressed, err := compress(origin)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}
	uncompressed, err := uncompress(compressed)
	if err != nil {
		t.Fatalf("error uncompress: %v", err)
	}
	if string(uncompressed) != string(origin) {
		t.Errorf("error compress uncompress origin: %v, uncompressed: %v", origin, uncompressed)
	}
}
