package main

import (
	"encoding/base64"
	"papergirl"

	"golang.org/x/crypto/bcrypt"
)

type BCryptPasswordHasher struct {
}

func NewBCryptPasswordHasher() papergirl.PasswordHasher {
	return &BCryptPasswordHasher{}
}

func (*BCryptPasswordHasher) Hash(password papergirl.Password) (papergirl.Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return papergirl.Password(base64.StdEncoding.EncodeToString(hashedPassword)), nil
}
