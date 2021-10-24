package login

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)
import _ "crypto/sha256"

type LoginRequest struct {
	Email    string
	Password Password `validate:"required"`
}

var validate = validator.New()

func (s UserService) Login(req LoginRequest) (*User, *Error) {
	log.Printf("Login user with email %v:", req.Email)
	err := validate.Struct(req)
	if err != nil {
		log.WithField("error", err).Info("Login user validation error")
		return nil, WrapError(err, "validation error", ValidationError)
	}
	user, loginError := s.CheckPassword(req.Email, req.Password.Hash())
	if loginError != nil {
		log.WithField("error", loginError).Info("User not found")
		return nil, loginError
	}
	log.WithField("email", req.Email).Info("User has logged in")
	return user, nil
}
