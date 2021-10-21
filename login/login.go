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

func (s UserService) Login(req LoginRequest) (user *User, err error) {
	errs := validate.Struct(req)
	if errs != nil {
		return nil, errs
	}
	return s.CheckPassword(req.Email, req.Password.Hash())
}
