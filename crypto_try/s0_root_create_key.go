package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalf("error MarshalPKCS1PrivateKey: %v", err)
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	outCAKeyFile := `crypto_try/generatedCAWithGo.key`
	outCAKeyFileWriter, err := os.OpenFile(
		outCAKeyFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error OpenFile: %v", err)
	}
	err = pem.Encode(outCAKeyFileWriter,
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	if err != nil {
		log.Fatalf("error pem_Encode: %v", err)
	}
	err = outCAKeyFileWriter.Close()
	if err != nil {
		log.Fatalf("error pem_Encode: %v", err)
	}
	log.Println("done")

	outPublicKeyFile := `crypto_try/generatedPublicKeyWithGo.pub`
	outPublicKeyFileWriter, err := os.OpenFile(
		outPublicKeyFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error OpenFile: %v", err)
	}
	err = pem.Encode(outPublicKeyFileWriter,
		&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})
	if err != nil {
		log.Fatalf("error pem_Encode: %v", err)
	}
	err = outPublicKeyFileWriter.Close()
	if err != nil {
		log.Fatalf("error pem_Encode: %v", err)
	}
	log.Println("done")
}
