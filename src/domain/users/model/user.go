package model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     string
	Username string
	Password []byte
}

func CreateUser(username string, password []byte) User {
	return User{
		UUID:     generateUUID(),
		Username: username,
		Password: password,
	}
}

func generateUUID() string {
	return uuid.NewString()
}
