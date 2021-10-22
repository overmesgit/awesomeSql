package login

type SignUpRequest struct {
	Username string
	Password Password
	Email    string
	Mood     string
}

func (s UserService) SingUp(req SignUpRequest) (*User, *Error) {
	user := &User{Username: req.Username,
		Email: req.Email, Mood: req.Mood}
	userId, err := s.Create(user, req.Password.Hash())
	if err != nil {
		return nil, err
	}
	user.UserID = userId
	return user, nil
}
