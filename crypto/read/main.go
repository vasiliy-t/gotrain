package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("key.pem")
	fmt.Println(data)
	bpriv, rest := pem.Decode(data)
	fmt.Printf("%+v\n\n", bpriv)

	bpub, rest := pem.Decode(rest)
	fmt.Printf("%+v\n\n", bpub)

	privkey, _ := x509.ParseECPrivateKey(bpriv.Bytes)
	fmt.Printf("%+v", privkey.PublicKey)
}
