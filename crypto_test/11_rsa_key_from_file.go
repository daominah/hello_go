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
func main11() {
	publicPem := `-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDg8kjD9gbGoTjwoRVspQTNNYou
PykzVNBasJDe1z0V5jJri34bOG87AvF0qURXZzLqOmXkYHBVgK9k4utmIADxMtyW
Gwh3dmcOSy0dm7FnVR1BcHpZ6u2QaiNmMDNR19RLWjujNYGRsn/PKyuoIfSvW4or
HkXZ+h4xqT9/ejmbzQIDAQAB
-----END RSA PUBLIC KEY-----`
	privatePem := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDg8kjD9gbGoTjwoRVspQTNNYouPykzVNBasJDe1z0V5jJri34b
OG87AvF0qURXZzLqOmXkYHBVgK9k4utmIADxMtyWGwh3dmcOSy0dm7FnVR1BcHpZ
6u2QaiNmMDNR19RLWjujNYGRsn/PKyuoIfSvW4orHkXZ+h4xqT9/ejmbzQIDAQAB
AoGAMN3RWuiubiYF9Zg4zEJI+b9gxk0oSSNqo9jpj89YUNKSL3S9L3KiD0LDa2F+
HDKqB+Ip0mP0404yTAtTsfrP2S2gEHoyTBW7U7zDO+cgQ68WdDX2J7BEvYmA9aNn
VdbAYVJP3j39da/xiRbSLqNzRenD02CanHjBIm+SLhsSp20CQQD7c2i1/FmmfBzD
MYUSf0YvW9zLCc041SSnF3GFBvwWA25uX6E/GeF7Daqq/JtsX/7gn40Ltn2iVsM4
+A0CtuiPAkEA5QQe19gXaQoK7+pSYmOeE0lBLfux/9InSN9p2KqYbiUWDbJjkVHG
YVs3P7KqJG/t1OrrrBETmhVOGHKADmDL4wJBAPBshhdT5Uh5bWr5g1qPVUVdGX0N
rysDKZuWn9VpO0m1GDbyuxPBpEXraF87T0TNeL+/7rXfVLsPKHTlQFNzHmMCQQDF
ZfTT5V3gWxisTRQv3F+/je/Rm9aEg/b6mB/a8siqf+rvaWjrNEpDVmVb0TtYZuXg
FZGH4bw8nsqOxfrc6dAzAkARH4j3SmVSD6l27cUVTAzwOtfXEXCo4TqipJHbFTIg
HnJQqCqN5LLgRITyXD3GwONQGgmXheRoIMYMGALmWKZN
-----END RSA PRIVATE KEY-----`
	_, _ = publicPem, privatePem

	block, _ := pem.Decode([]byte(publicPem))
	keyI, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("err", err)
	}
	key, isOk := keyI.(*rsa.PublicKey)
	if isOk != true {
		fmt.Println("isOk", isOk)
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
