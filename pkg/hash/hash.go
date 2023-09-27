package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func String2Hash(v string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost) //加密处理
	return string(hash)
}

func CompareHash(oldHash, new string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldHash), []byte(new))
	log.Println(err)
	return err == nil
}
