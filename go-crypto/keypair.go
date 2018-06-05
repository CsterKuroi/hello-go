package main

import(
	"fmt"
	"crypto/rand"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/nacl/box"
)

func main() {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("err")
	}
	if pub == nil || priv == nil {
		fmt.Printf("Got nil")
	}
	pub58 := base58.Encode((*pub)[:])
	priv58 := base58.Encode((*priv)[:])
	fmt.Println(pub58,priv58)
}
