package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := `username:password`
	base64ed := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("-----base64ed:")
	fmt.Println(base64ed)
	decodedB, err := base64.StdEncoding.DecodeString(base64ed)
	if err != nil {
		fmt.Printf("error DecodeString: %v", err)
		return
	}
	decoded := string(decodedB)
	fmt.Println("-----decoded:")
	fmt.Println(decoded)
}
