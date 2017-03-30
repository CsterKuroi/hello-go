package main

import (
        "fmt"
        "crypto/rand"
        "github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/nacl/box"
)


func byteString(p []byte) string {
        for i := 0; i < len(p); i++ {
                if p[i] == 0 {
                        return string(p[0:i])
                }
        }
        return string(p)
}


func main() {
	// Create random public and private keys
        //var publicKeyBytes,privateKeyBytes *[32]byte
	publicKeyBytes, privateKeyBytes, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
        slice1 := publicKeyBytes[:]
        slice2 := privateKeyBytes[:]
	publicKeyBase58 := base58.Encode(slice1)
	privateKeyBase58 := base58.Encode(slice2)
        fmt.Printf("%s,%s\n",publicKeyBase58,privateKeyBase58)
}

