package main

import (
	"encoding/base64"
	"papergirl"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	hasher := NewBCryptPasswordHasher()

	password := papergirl.Password("password")
	hashedPassword, err := hasher.Hash(password)
	assert.Nil(t, err)

	decodedHashedPassword, err := base64.StdEncoding.DecodeString(string(hashedPassword))
	assert.Nil(t, err)
	err = bcrypt.CompareHashAndPassword(decodedHashedPassword, []byte(password))
	assert.Nil(t, err)
}
