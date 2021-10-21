package login

import (
	"log"
)

type SignUpRequest struct {
	Username string
	Password Password
	Email    string
	Mood     string
}

func (s UserService) SingUp(req SignUpRequest) (*User, error) {
	user := &User{Username: req.Username,
		Email: req.Email, Mood: req.Mood}
	userId, err := s.Create(user, req.Password.Hash())
	if err != nil {
		log.Printf("CreateUser %v", err)
		return nil, err
	}
	user.UserID = userId
	return user, nil
}
