package papergirl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	initializeTest()
	username := Username("username")
	password := Password("password")
	role := RoleAdmin

	err := CreateUser(username, password, role)
	assert.Nil(t, err)

	foundUsers, _ := users.FindAll()
	assert.Len(t, foundUsers, 1)

	expectedPassword := Password("drowssap")
	createdUser := foundUsers[0]
	assert.Equal(t, username, createdUser.Username)
	assert.Equal(t, role, createdUser.Role)
	assert.Equal(t, expectedPassword, createdUser.Password)
}

func TestCountUsers(t *testing.T) {
	initializeTest()
	users.Save(User{})
	users.Save(User{})

	count, err := CountUsers()
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}
