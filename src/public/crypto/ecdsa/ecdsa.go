package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func Example(msg string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}
