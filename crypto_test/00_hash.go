package main

import (
	"crypto/md5"
	//	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func init() {
	_ = json.Marshal
}

//func main() {
func main00() {
	//	h := sha256.New()
	h := md5.New()

	h.Write([]byte("123qwe"))
	result := h.Sum(nil)
	fmt.Printf("len %v %T %v\n", len(result), result, result)
	fmt.Printf("hex %v\n", hex.EncodeToString(result))
	fmt.Printf("base64 %v\n", base64.StdEncoding.EncodeToString(result))
}
