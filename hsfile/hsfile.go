package hsfile

import (
	"crypto/sha256"
	"hash"
	"io"
	"log"
	"os"
)

func OpFile(filePath string) *os.File {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("File open error: ", err)
	}

	return f
}

func HashFile(filePath string) hash.Hash {
	obj := OpFile(filePath)
	cryp := sha256.New()
	if _, err := io.Copy(cryp, obj); err != nil {
		log.Fatal("Hash Error: ", err)
	}
	defer obj.Close()
	return cryp
}
