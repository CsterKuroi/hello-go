package main
import "io"
import "os"
import "fmt"
import "hash"
import "crypto/md5"
import "crypto/sha1"
import "encoding/hex"
import "crypto/sha256"
import "crypto/sha512"
import "golang.org/x/crypto/sha3"

func GenHash(val, hash_type string) string{
	var hash hash.Hash
	var x string = "bad_hash_type"

	switch hash_type {
		case "md5": hash = md5.New()
		case "sha1": hash = sha1.New()
		case "sha256": hash = sha256.New()
		case "sha512": hash = sha512.New()
		case "sha3-256": hash = sha3.New256()
		case "sha3-512": hash = sha3.New512()
	}

	if hash != nil {
		file, err := os.Open(val)
		if err == nil {
			io.Copy(hash, file)
			defer file.Close()
			x = "file: "
		}else{
			hash.Write([]byte(val))
			x = hash_type+": "
		}
		x += hex.EncodeToString(hash.Sum(nil))
	}

	return x
}

func main(){
	if len(os.Args) < 3 {
		fmt.Println("genhash string|file md5|sha1|sha256|sha512|sha3-256|sha3-512")
		os.Exit(0)
	}
	fmt.Println(GenHash(os.Args[1], os.Args[2]))
}
