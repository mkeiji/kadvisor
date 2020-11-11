package KeiPassUtil

import (
	"golang.org/x/crypto/bcrypt"
	"kadvisor/server/repository/structs"
)

func HashAndSalt(user *structs.User) (string, error) {
	byte := []byte(user.Login.Password)

	hash, err := bcrypt.GenerateFromPassword(byte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func IsValidPassword(hashedPwd string, plainPwd string) bool {
	bytePwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}

	return true
}
