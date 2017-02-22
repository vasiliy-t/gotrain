// This demo show how to generate ECDSA key pair
//
// Equivalent openssl utility commands are:
// $ openssl ecparam -genkey -name prime256v1 -noout -out priv.pem
// $ openssl ec -in priv.pem -pubout -out pub.pem
//
// OpenSSL docs on subject
// https://wiki.openssl.org/index.php/Command_Line_Elliptic_Curve_Operations
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	priv, _ := x509.MarshalECPrivateKey(key)
	privpem := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: priv})
	privfile, _ := os.OpenFile("priv.pem", os.O_RDWR|os.O_CREATE, 0755)
	privfile.Write(privpem)
	privfile.Close()

	pub, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)
	pubpem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pub})
	pubfile, _ := os.OpenFile("pub.pem", os.O_RDWR|os.O_CREATE, 0755)
	pubfile.Write(pubpem)
	pubfile.Close()
}
