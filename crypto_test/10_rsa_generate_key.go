package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
//	func main() {
	Priv, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Printf("Priv %+v\n", Priv)
	fmt.Printf("Pub %+v\n", Priv.PublicKey)
	//	fmt.Println("N", Priv.N.BitLen())

	// MarshalPKCS1PrivateKey converts a private key to ASN.1 DER encoded form
	PrivASN1 := x509.MarshalPKCS1PrivateKey(Priv)
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: PrivASN1,
		},
	)
	ioutil.WriteFile("crypto_test/key.priv", pemdata, 0644)

	// MarshalPKIXPublicKey serialises a public key to DER-encoded PKIX format
	PubASN1 := x509.MarshalPKCS1PublicKey(&Priv.PublicKey)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	PubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})
	ioutil.WriteFile("crypto_test/key.pub", PubBytes, 0644)
	fmt.Println("xong")
}
