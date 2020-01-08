package KeiPassUtil

import (
	"golang.org/x/crypto/bcrypt"
	"kadvisor/server/repository/structs"
	"log"
)

func HashAndSalt(user *structs.User) {
	byte := []byte(user.Login.Password)

	hash, err := bcrypt.GenerateFromPassword(byte, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	user.Login.Password = string(hash)
}

func IsValidPassword(hashedPwd string, plainPwd string) bool {
	bytePwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
