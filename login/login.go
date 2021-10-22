package login

import (
	"github.com/go-playground/validator/v10"
)
import _ "crypto/sha256"

type LoginRequest struct {
	Email    string
	Password Password `validate:"required"`
}

var validate = validator.New()

func (s UserService) Login(req LoginRequest) (*User, *Error) {
	errs := validate.Struct(req)
	if errs != nil {
		return nil, WrapError(errs, "validation error", ValidationError)
	}
	user, err := s.CheckPassword(req.Email, req.Password.Hash())
	if err != nil {
		return nil, err
	}
	return user, nil
}
