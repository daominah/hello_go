package main

import (
	"crypto/x509"
	"encoding/pem"
	"log"
)

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

	tmp, _ := pem.Decode([]byte(publicPem))
	publicKey, err := x509.ParsePKCS1PublicKey(tmp.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	tmp, _ = pem.Decode([]byte(privatePem))
	privateKey, err := x509.ParsePKCS1PrivateKey(tmp.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("privateKeyLen %v publicKey %+v\n", privateKey.N.BitLen(), privateKey)
	if privateKey.PublicKey.N.Cmp(publicKey.N) != 0 ||
		privateKey.PublicKey.E != publicKey.E {
			log.Fatal("keys is not a pair")
	}
}
