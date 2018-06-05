package main

import(
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
)

func main() {
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes[0:32])
	fmt.Println("pub:",publicKeyBase58)
	fmt.Println("priv:",privateKeyBase58)
}
