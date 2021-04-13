package main

import (
	"fmt"

	"golang.org/x/text/encoding/unicode"
)

func main() {
	s := "exampleString0"
	encoder16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder()
	encoded, err := encoder16be.Bytes([]byte(s))
	fmt.Printf("encoder16be Bytes: %v, %#v\n", err, encoded)

	encoder8 := unicode.UTF8.NewEncoder()
	encoded, err = encoder8.Bytes([]byte(s))
	fmt.Printf("encoder8 Bytes: %v, %#v\n", err, encoded)
	fmt.Printf("slice bytes: %#v\n", []byte(s))
}
