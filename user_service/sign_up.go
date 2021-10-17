package user_service

import (
	"log"
)

func SingUp(
	storage Storage,
	username string, password string, email string,
	mood string,
) (*User, error) {
	user := &User{Username: username,
		Email: email, Mood: mood}
	userId, err := storage.Create(user, user.Hash(password))
	if err != nil {
		log.Printf("CreateUser %v", err)
		return nil, err
	}
	user.UserID = userId
	return user, nil
}
