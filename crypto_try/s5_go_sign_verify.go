package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// read RSA private key from file

	encodedPrivatePem, err := ioutil.ReadFile(`/home/tungdt/go/src/github.com/daominah/hello_go/crypto_try/myorg1.key`)
	if err != nil {
		log.Fatal(err)
	}
	block0, _ := pem.Decode([]byte(encodedPrivatePem))
	decrypted, err := x509.DecryptPEMBlock(block0, []byte("123qwe"))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(decrypted)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("privateKey: %#v\n", privateKey)

	// sign a example text

	msg0 := []byte("text to be signed")
	hashedMsg0 := sha256.Sum256(msg0)

	sign, err := privateKey.Sign(rand.Reader, hashedMsg0[:], crypto.SHA256)
	if err != nil {
		log.Fatalf("error SignPKCS1v15: %v\n", sign)
		return
	}
	base64edSignature := base64.StdEncoding.EncodeToString(sign)
	log.Printf("base64edSignature: %v\n", base64edSignature)

	// read RSA public key from certificate

	publicPem, err := ioutil.ReadFile("/home/tungdt/go/src/github.com/daominah/hello_go/crypto_try/myorg1.crt")
	if err != nil {
		log.Fatal(err)
	}
	block1, _ := pem.Decode([]byte(publicPem))
	certificate, err := x509.ParseCertificate(block1.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	pubKey, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		log.Fatal("error not a rsa public key")
	}

	// verify a signature

	fakeHashed := bytes.Repeat([]byte("a"), 32)
	verifyErr1 := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, fakeHashed[:32], sign)
	log.Printf("verifyErr fakeHashed: %v\n", verifyErr1)
	verifyErr2 := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashedMsg0[:32], sign)
	log.Printf("verifyErr: %v\n", verifyErr2)
}
