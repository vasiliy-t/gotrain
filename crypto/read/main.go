package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	privdata, _ := ioutil.ReadFile("priv.pem")
	bpriv, _ := pem.Decode(privdata)
	privkey, _ := x509.ParseECPrivateKey(bpriv.Bytes)
	fmt.Printf("%+v\n\n", privkey)

	pubdata, _ := ioutil.ReadFile("pub.pem")
	bpub, _ := pem.Decode(pubdata)
	pubkey, _ := x509.ParsePKIXPublicKey(bpub.Bytes)
	fmt.Printf("%+v", pubkey.(*ecdsa.PublicKey))
}
