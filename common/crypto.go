package main

import(
	"os"
	"fmt"
	"hash"
	"bytes"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
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

func GetPubByPriv(priv string) string {
	privByte := base58.Decode(priv)
	publicKeyBytes, _, err := ed25519.GenerateKey(bytes.NewReader(privByte))
	if err != nil {
                panic(err)
        }
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	return publicKeyBase58
}

func Sign(priv string,msg string) string {
	pub := GetPubByPriv(priv)
	privByte := base58.Decode(priv)
	pubByte := base58.Decode(pub)
	privateKey := make([]byte, 64)
	copy(privateKey[:32], privByte)
	copy(privateKey[32:], pubByte)
	sigByte := ed25519.Sign(privateKey,[]byte(msg))
        return base58.Encode(sigByte)
}

func Verify(pub string,msg string,sig string) bool {
	pubByte := base58.Decode(pub)
	publicKey := make([]byte, 32)
	copy(publicKey, pubByte)
	return ed25519.Verify(publicKey,[]byte(msg), base58.Decode(sig))
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

	pub2 := GetPubByPriv(privateKeyBase58)
	fmt.Println("pub2:",pub2)

	sig := Sign(privateKeyBase58,os.Args[1])
	fmt.Println("sig:",sig)

	result :=Verify(publicKeyBase58,os.Args[1],sig)
	fmt.Println("verify:",result)
}
