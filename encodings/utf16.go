package main

import (
	"log"

	"golang.org/x/text/encoding/unicode"
)

func main() {
	s := "exampleString0"
	encoder16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder()
	encoded16, err := encoder16be.Bytes([]byte(s))
	log.Printf("encoder16be Bytes: %v, %#v\n", err, encoded16)

	decoder16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
	decoded, err := decoder16be.Bytes(encoded16)
	log.Printf("decoder16be Bytes: %v\n", string(decoded))

	encoder8 := unicode.UTF8.NewEncoder()
	encoded8, err := encoder8.Bytes([]byte(s))
	log.Printf("encoder8 Bytes: %v, %#v\n", err, encoded8)
	log.Printf("slice bytes: %#v\n", []byte(s))
}
