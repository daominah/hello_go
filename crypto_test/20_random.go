package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	//func main20()  {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid := fmt.Sprintf("%X", b)
	uuidB64 := base64.StdEncoding.EncodeToString(b)
	fmt.Println("uuid:", uuid, uuidB64)
}
