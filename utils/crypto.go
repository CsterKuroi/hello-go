package main
import "os"
import "fmt"
import "hash"
import "encoding/hex"
import "golang.org/x/crypto/sha3"

func HashData(val string) string{
	var hash hash.Hash
	var x string = ""
        hash = sha3.New256()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Please input string to hash(sha3-256)")
		os.Exit(0)
	}
	fmt.Println(HashData(os.Args[1]))
}
