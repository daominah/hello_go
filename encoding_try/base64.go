package main

import (
	"encoding/base64"
	"log"
)

func main() {
	s := `username:password`
	base64ed := base64.StdEncoding.EncodeToString([]byte(s))
	log.Println(base64ed)

	if false {
		decodedB, err := base64.StdEncoding.DecodeString(base64ed)
		if err != nil {
			log.Fatal(err)
		}
		decoded := string(decodedB)
		log.Println(decoded)
	} else {
		decodedB := make([]byte, base64.StdEncoding.DecodedLen(len(base64ed)))
		n, err := base64.StdEncoding.Decode(decodedB, []byte(base64ed))
		if err != nil {
			log.Fatal(err)
		}
		decoded := string(decodedB[:n])
		log.Println(decoded)
	}
}