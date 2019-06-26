package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// output is base64 encoded,
// padding scheme is the original PKCS#1 v1.5
func RsaEncrypt(plaintext string, publicKey *rsa.PublicKey) (string, error) {
	input := []byte(plaintext)
	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, input)
	if err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(cipherBytes), nil
	}
}

func RsaDecrypt(cipherBase64 string, privateKey *rsa.PrivateKey) (string, error) {
	input, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return "", err
	}
	bs, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, input)
	if err != nil {
		return "", err
	} else {
		return string(bs), nil
	}
}

//func main() {
func main() {
	publicPem := `-----BEGIN RSA PUBLIC KEY-----
MCgCIQDAwhGew3oBBQdf8MyMQDG8JJuhUEqeozcbJYGEhbCl3QIDAQAB
-----END RSA PUBLIC KEY-----`
	privatePem := `-----BEGIN RSA PRIVATE KEY-----
MIGrAgEAAiEAwMIRnsN6AQUHX/DMjEAxvCSboVBKnqM3GyWBhIWwpd0CAwEAAQIg
KWvKo4Y3+m4dNpWlLuJAjWCbao4mgKwrehuppFTWtaECEQDJ0xUTUUXapU/r8If5
7nKFAhEA9H/113xJKAmQMtOudoIBeQIQKdkX1KKUfmqqsLx2JW+41QIRAODLUV24
sI42FLUWeJ4Os4kCEQCqPk1+fqJc7cEx46aQYgn5
-----END RSA PRIVATE KEY-----`
	_, _ = publicPem, privatePem

	block, _ := pem.Decode([]byte(publicPem))
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("pubKeyLen %v key %+v\n", key.N.BitLen(), key)
	// echo 'hohohaha' | openssl rsautl -encrypt -pubin -inkey hfc.pub -out >(base64)
	plaintext := "hohohaha"
	ciphertext, err := RsaEncrypt(plaintext, key)
	fmt.Printf("err: %v, ciphertext:\n%v\n", err, ciphertext)

	// decode
	block, _ = pem.Decode([]byte(privatePem))
	key2, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("err2", err)
	}
	fmt.Printf("privateKeyLen %v key %+v\n", key2.N.BitLen(), key2)
	ptext2, err := RsaDecrypt(ciphertext, key2)
	fmt.Printf("err: %v, ptext2:\n%v\n", err, ptext2)
}
