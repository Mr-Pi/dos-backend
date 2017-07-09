package permissions

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
)

func HashPasswordNew(password string) (hashOut string, saltOut string) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, 64)
	if err != nil {
		log.Fatal(err)
	}
	saltOut = hex.EncodeToString(salt)
	hashOut = hex.EncodeToString(hash)
	return hashOut, saltOut
}
func HashPasswordOld(password string, saltIn string) string {
	salt, err := hex.DecodeString(saltIn)
	if err != nil {
		log.Fatal(err)
	}
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, 64)
	if err != nil {
		log.Fatal(err)
	}
	hashOut := hex.EncodeToString(hash)
	return hashOut
}
