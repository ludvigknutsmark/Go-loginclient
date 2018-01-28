package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

var (
	Privkey *ecdsa.PrivateKey
	Pubkey  *ecdsa.PublicKey
)

func generateKeyPair() *ecdsa.PrivateKey {
	Privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return Privkey
}
