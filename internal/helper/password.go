package helper

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed", err)
	}
	return string(hashedPassword), nil
}
func CheckPassword(hashedPass, plainPass string)error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass),[]byte(plainPass))
	if err != nil {
		return errors.New("invalid password" + err.Error())
	}
	return nil
}