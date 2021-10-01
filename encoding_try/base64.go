package main

import (
	"encoding/base64"
	"log"
)

func main() {
	s := `username:password`
	base64ed := base64.StdEncoding.EncodeToString([]byte(s))
	println("base64ed:")
	println(base64ed)
	println("-----------------------------------------------------------")

	if false { // decode method 0
		decodedB, err := base64.StdEncoding.DecodeString(base64ed)
		if err != nil {
			log.Fatal(err)
		}
		decoded := string(decodedB)
		println(decoded)
	} else { // decode method 1
		decodedB := make([]byte, base64.StdEncoding.DecodedLen(len(base64ed)))
		n, err := base64.StdEncoding.Decode(decodedB, []byte(base64ed))
		if err != nil {
			log.Fatal(err)
		}
		decoded := string(decodedB[:n])
		println("decoded:")
		println(decoded)
	}
}
