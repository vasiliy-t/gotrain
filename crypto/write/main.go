package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	priv, _ := x509.MarshalECPrivateKey(key)
	privpem := pem.EncodeToMemory(&pem.Block{Type: "PRIV KEY", Bytes: priv})

	fmt.Printf("%+v", key.D)
	fmt.Printf("%+v", key.PublicKey)

	pub, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)
	pubpem := pem.EncodeToMemory(&pem.Block{Type: "PUB KEY", Bytes: pub})

	f, _ := os.OpenFile("key.pem", os.O_RDWR|os.O_CREATE, 0755)

	f.Write(privpem)
	f.Write(pubpem)
	f.Close()
}
