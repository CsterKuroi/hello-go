package main

import(
	"os"
	"fmt"
	"hash"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
)

func HashData(val string) string {
	var hash hash.Hash
	var x string = ""
        hash = sha3.New256()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func GenerateKeyPair() (string,string) {
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes[0:32])
	return publicKeyBase58,privateKeyBase58
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input string to hash(sha3-256)")
		os.Exit(0)
	}
	fmt.Println(HashData(os.Args[1]))
	publicKeyBase58,privateKeyBase58 :=GenerateKeyPair()
	fmt.Println("pub:",publicKeyBase58)
        fmt.Println("priv:",privateKeyBase58)
}
