package login

import log "github.com/sirupsen/logrus"

type SignUpRequest struct {
	Username string   `validate:"required"`
	Password Password `validate:"required"`
	Email    string   `validate:"required"`
	Mood     string   `validate:"required"`
	Type     string   `validate:"required"`
}

func userFromSignupRequest(req SignUpRequest) *User {
	return &User{
		Username: req.Username,
		Email:    req.Email,
		Mood:     req.Mood,
		Type:     UserType(req.Type),
	}
}

func (s UserService) SignUp(req SignUpRequest) (*User, *Error) {
	err := validate.Struct(req)
	if err != nil {
		log.WithField("error", err).Info("Login user validation error")
		return nil, WrapError(err, "validation error", ValidationError)
	}
	user := userFromSignupRequest(req)
	log.WithFields(log.Fields{"username": user.Username}).Info("Create new user")
	userId, signUpError := s.Create(user, req.Password.Hash())
	if signUpError != nil {
		log.WithFields(log.Fields{"username": user.Username, "error": signUpError}).
			Info("Create new user error")
		return nil, signUpError
	}
	user.UserID = userId
	log.WithFields(log.Fields{"username": user.Username, "id": userId}).
		Info("New user created")
	return user, nil
}
