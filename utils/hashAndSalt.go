package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}
	return string(hash)
}
func ComparePassword(HashedPwd string, PlainPassword []byte) bool {
	byteHash := []byte(HashedPwd)
	pass := bcrypt.CompareHashAndPassword(byteHash, PlainPassword)
	if pass != nil {
		return false
	}
	return true
}
